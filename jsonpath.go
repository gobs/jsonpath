//go:generate antlr4 -Dlanguage=Go -package parser -o parser Jsonpath.g4
package jsonpath

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gobs/jsonpath/parser"
	"github.com/gobs/simplejson"
)

func asInt(t antlr.Token, val int) int {
	if t != nil {
		val, _ = strconv.Atoi(t.GetText())
	}

	return val
}

type nodeType int

const (
	ARRAY_RANGE nodeType = iota
	ARRAY_ITEMS
	CHILD
	DESCENDANT
	ALL_ITEMS

	MAX_RANGE = 0x7ffffff

	TOKEN_ANY    = "*"
	TOKEN_LENGTH = "!" // not a valid identifiers
)

// Node defines a processing node
type Node struct {
	nodeType         nodeType
	name             string
	indices          []int
	start, end, step int
}

func (n Node) String() string {
	switch n.nodeType {
	case ARRAY_RANGE:
		return fmt.Sprintf("ARRAY RANGE start:%v end:%v step:%v", n.start, n.end, n.step)

	case ARRAY_ITEMS:
		return fmt.Sprintf("ARRAY ITEMS %v", n.indices)

	case ALL_ITEMS:
		return "ALL_ITEMS"

	case CHILD:
		return fmt.Sprintf("CHILD %v", n.name)

	case DESCENDANT:
		return fmt.Sprintf("DESCENDANT %v", n.name)
	}

	return "UNKNOWN"
}

// Processor defines the JsonPath grammar processor
type Processor struct {
	*parser.BaseJsonpathListener

	Nodes []Node
	err   error
}

// NewProcessor creates a new grammar Processor
func NewProcessor() *Processor {
	return &Processor{}
}

func (p *Processor) addNode(n Node) {
	p.Nodes = append(p.Nodes, n)
}

//
// VisitErrorNode is called when there is an error
//
func (p *Processor) VisitErrorNode(node antlr.ErrorNode) {
	p.err = fmt.Errorf("%v", node)
}

//
// ExitDotExpr is called when production dotExpr is exited.
//
func (p *Processor) ExitDotExpr(ctx *parser.DotExprContext) {
	name := TOKEN_ANY // any value

	if ctx.Identifier() != nil {
		name = ctx.Identifier().GetText()
	}

	if ctx.Length() != nil {
		name = TOKEN_LENGTH
	}

	if ctx.DOTS().GetText() == ".." {
		p.addNode(Node{nodeType: DESCENDANT, name: name})
	} else {
		p.addNode(Node{nodeType: CHILD, name: name})
	}
}

//
// ExitItemsExpr is called when production itemsExpr is exited.
//
func (p *Processor) ExitItemsExpr(ctx *parser.ItemsExprContext) {
	n := Node{nodeType: ARRAY_ITEMS}

	for _, i := range ctx.AllINT() {
		n.indices = append(n.indices, asInt(i.GetSymbol(), -1))
	}

	p.addNode(n)
}

//
// ExitRangeExpr is called when production rangeExpr is exited.
//
func (p *Processor) ExitRangeExpr(ctx *parser.RangeExprContext) {
	p.addNode(Node{
		nodeType: ARRAY_RANGE,
		start:    asInt(ctx.GetStartIndex(), 0),
		end:      asInt(ctx.GetEndIndex(), MAX_RANGE),
		step:     asInt(ctx.GetStep(), 1)})
}

//
// ExitNameExpr is called when production nameExpr is exited.
//
func (p *Processor) ExitNameExpr(ctx *parser.NameExprContext) {
	p.addNode(Node{nodeType: CHILD, name: strings.Trim(ctx.QUOTED().GetText(), "'")})
}

//
// ExitStarExpr is called when production starExpr is exited.
//
func (p *Processor) ExitStarExpr(ctx *parser.StarExprContext) {
	p.addNode(Node{nodeType: ALL_ITEMS})
}

//
// Parse parses a JsonPath expression
//
func (p *Processor) Parse(expr string) error {
	input := antlr.NewInputStream(expr)
	lexer := parser.NewJsonpathLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewJsonpathParser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	antlr.ParseTreeWalkerDefault.Walk(p, parser.Jsonpath())
	return p.err
}

type map_type = map[string]interface{}
type array_type = []interface{}

func (p *Processor) find(name string, j interface{}) (ret []interface{}) {
	var a []interface{}

	if aj, ok := j.(array_type); ok {
		a = aj
	} else {
		a = array_type{j}
	}

	for _, c := range a {
		switch t := c.(type) {
		case map_type:
			for k, v := range t {
				if k == name {
					ret = append(ret, v)
				} else {
					if name == TOKEN_ANY {
						ret = append(ret, v)
					}

					res := p.find(name, v)
					ret = append(ret, res...)
				}
			}

		case array_type:
			res := p.find(name, t)
			ret = append(ret, res...)
		}
	}

	return ret
}

func atIndex(i int, a array_type) (interface{}, bool) {
	l := len(a)

	if i < 0 {
		i += l // start from the end
	}

	if i >= 0 && i < l {
		return a[i], true
	}

	return nil, false
}

func getLength(v interface{}) interface{} {
	switch t := v.(type) {
	case map_type:
		return len(t)
	case array_type:
		return len(t)
	default:
		return nil
	}
}

//
// Process input object according to parsed JsonPath
//
func (p *Processor) Process(v interface{}) interface{} {
	if p.err != nil {
		return nil
	}

	if j, ok := v.(*simplejson.Json); ok {
		v = j.Data()
	}

	last := len(p.Nodes) - 1

	for s, n := range p.Nodes {
		j := simplejson.AsJson(v)

		switch n.nodeType {
		case ARRAY_RANGE:
			a := j.MustArray()
			res := array_type{}
			start, end, step := n.start, n.end, n.step

			if end > len(a) {
				end = len(a)
			}
			if start < 0 {
				start = len(a) + start
			}

			for i := start; i < end; i += step {
				res = append(res, a[i])
			}

			v = res

		case ARRAY_ITEMS:
			a := j.MustArray()
			res := array_type{}
			for _, i := range n.indices {
				if v, ok := atIndex(i, a); ok {
					res = append(res, v)
				}
			}

			if len(res) == 1 && len(n.indices) == 1 {
				v = res[0]
			} else {
				v = res
			}

		case CHILD:
			res := array_type{}
			a, err := j.Array()
			if err != nil {
				a = array_type{j.Data()}
			} else if n.name == TOKEN_LENGTH {
				res = append(res, getLength(a))
				v = res
				continue
			}

			for _, c := range a {
				if n.name == TOKEN_LENGTH {
					res = append(res, getLength(a))
					continue
				}

				if m, ok := c.(map_type); ok {
					if n.name == TOKEN_ANY {
						for _, mv := range m {
							res = append(res, mv)
						}
					} else if v, ok := m[n.name]; ok {
						res = append(res, v)
					}
				} else if n.name == TOKEN_ANY {
					res = append(res, c)
				}
			}

			if len(res) == 1 {
				v = res[0]
			} else {
				v = res
			}

		case DESCENDANT:
			v = p.find(n.name, v)
			if len(v.(array_type)) == 1 && s < last {
				v = v.(array_type)[0]
			}

		case ALL_ITEMS:
			res := array_type{}

			if a, err := j.Array(); err == nil {
				for _, v := range a {
					res = append(res, v)
				}
			} else if m, err := j.Map(); err == nil {
				for _, v := range m {
					res = append(res, v)
				}
			} else {
				fmt.Printf("what is %T", j.Data())
			}

			v = res
		}
	}

	return v
}
