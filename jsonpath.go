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
	ANY

	MAX_RANGE = 0x7ffffff

	TOKEN_ANY = "*"
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

func (p *Processor) find(name string, j interface{}) (ret []interface{}) {
	var a []interface{}

	if aj, ok := j.([]interface{}); ok {
		a = aj
	} else {
		a = []interface{}{j}
	}

	for _, c := range a {
		switch t := c.(type) {
		case map[string]interface{}:
			for k, v := range t {
				if name == TOKEN_ANY || k == name {
					ret = append(ret, v)
				} else {
					res := p.find(name, v)
					ret = append(ret, res...)
				}
			}

		case []interface{}:
			res := p.find(name, t)
			ret = append(ret, res...)
		}
	}

	return ret
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

	for _, n := range p.Nodes {
		j := simplejson.AsJson(v)

		switch n.nodeType {
		case ARRAY_RANGE:
			a := j.MustArray()
			res := []interface{}{}
			start, end, step := n.start, n.end, n.step

			if end > len(a) {
				end = len(a)
			}
			if start < 0 {
				start = len(a) - start
			}

			for i := start; i < end; i += step {
				res = append(res, a[i])
			}

			v = res

		case ARRAY_ITEMS:
			a := j.MustArray()
			res := []interface{}{}
			for _, i := range n.indices {
				if i >= 0 && i < len(a) {
					res = append(res, a[i])
				}
			}

			v = res

		case CHILD:
			if a := j.GetIndex(0); !a.Nil() {
				j = a
			}
			m := j.MustMap()
			if n.name == TOKEN_ANY {
				res := []interface{}{}
				for _, mv := range m {
					res = append(res, mv)
				}

				v = res
			} else {
				v = m[n.name]
			}

		case DESCENDANT:
			v = p.find(n.name, v)
		}
	}

	return v
}
