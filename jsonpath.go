//go:generate antlr4 -Dlanguage=Go -package parser -o parser Jsonpath.g4
package jsonpath

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gobs/jsonpath/parser"
	"github.com/gobs/simplejson"
)

type ProcessOptions int

const (
	Default  ProcessOptions = 0
	Enhanced ProcessOptions = 1
	Collapse ProcessOptions = 2
)

type map_type = map[string]interface{}
type array_type = []interface{}
type set_type map[string]struct{}

func (s set_type) Set(v string) {
	s[v] = struct{}{}
}

func (s set_type) Contains(v string) bool {
	_, ok := s[v]
	return ok
}

func (s set_type) List() []string {
	var l []string

	for k := range s {
		l = append(l, k)
	}

	return l
}

func asInt(t antlr.Token, val int) int {
	if t != nil {
		val, _ = strconv.Atoi(t.GetText())
	}

	return val
}

func asFloat(t antlr.Token, val float64) float64 {
	if t != nil {
		val, _ = strconv.ParseFloat(t.GetText(), 64)
	}

	return val
}

func toInt(v interface{}) int {
	switch t := v.(type) {
	case int:
		return t

	case float64:
		return int(t)

	case string:
		i, _ := strconv.Atoi(t)
		return i
	}

	return 0
}

func toFloat(v interface{}) float64 {
	switch t := v.(type) {
	case int:
		return float64(t)

	case float64:
		return t

	case string:
		f, _ := strconv.ParseFloat(t, 64)
		return f
	}

	return 0.0
}

func toBool(v interface{}) bool {
	switch t := v.(type) {
	case bool:
		return t
	case int:
		return t > 0

	case float64:
		return int(t) > 0

	case string:
		return t != ""
	}

	return false
}

func toString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

type nodeType int

const (
	ARRAY_RANGE nodeType = iota
	ARRAY_ITEMS
	CHILD
	DESCENDANT
	ALL_ITEMS
	FILTER_EXPR
	SCRIPT_EXPR

	MAX_RANGE = 0x7ffffff

	TOKEN_ANY    = "*"
	TOKEN_LENGTH = "!" // not a valid identifiers
	TOKEN_REGEX  = "=~"
)

// Operation is a method that perform an operation (conditional or not)
type Operation func(v interface{}) interface{}

// Node defines a processing node
type Node struct {
	nodeType nodeType

	opId    string      // for DESCENDANT/FILTER_EXPR/SCRIPT_EXPR
	opValue interface{} // for FILTER_EXP/SCRIPT_EXPR
	op      Operation
	opName  string // used for debugging

	names            set_type // for CHILD
	indices          []int    // for ARRAY_ITEMS
	start, end, step int      // for ARRAY_RANGE
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
		return fmt.Sprintf("CHILD %v", n.names.List())

	case DESCENDANT:
		return fmt.Sprintf("DESCENDANT %v", n.names.List())

	case FILTER_EXPR:
		return fmt.Sprintf("FILTER %v %v %v", n.opId, n.opName, n.opValue)

	case SCRIPT_EXPR:
		return fmt.Sprintf("SCRIPT %v %v %v", n.opId, n.opName, n.opValue)
	}

	return "UNKNOWN"
}

func (n *Node) CompareExists(v interface{}) interface{} {
	return true
}

func (n *Node) CompareEqual(v interface{}) interface{} {
	switch t := v.(type) {
	case bool:
		return t == toBool(n.opValue)

	case int:
		return t == toInt(n.opValue)

	case float64:
		return t == toFloat(n.opValue)

	case string:
		return t == toString(n.opValue)
	}

	return false
}

func (n *Node) CompareNotEqual(v interface{}) interface{} {
	switch t := v.(type) {
	case bool:
		return t != toBool(n.opValue)

	case int:
		return t != toInt(n.opValue)

	case float64:
		return t != toFloat(n.opValue)

	case string:
		return t != toString(n.opValue)
	}

	return false
}

func (n *Node) CompareGreater(v interface{}) interface{} {
	switch t := v.(type) {
	case bool:
		return t == true && toBool(n.opValue) == false

	case int:
		return t > toInt(n.opValue)

	case float64:
		return t > toFloat(n.opValue)

	case string:
		return t > toString(n.opValue)
	}

	return false
}

func (n *Node) CompareLess(v interface{}) interface{} {
	switch t := v.(type) {
	case bool:
		return t == false && toBool(n.opValue) == true

	case int:
		return t < toInt(n.opValue)

	case float64:
		return t < toFloat(n.opValue)

	case string:
		return t < toString(n.opValue)
	}

	return false
}

func (n *Node) CompareGreaterEqual(v interface{}) interface{} {
	switch t := v.(type) {
	case bool:
		return t == true || t == toBool(n.opValue)

	case int:
		return t >= toInt(n.opValue)

	case float64:
		return t >= toFloat(n.opValue)

	case string:
		return t >= toString(n.opValue)
	}

	return false
}

func (n *Node) CompareLessEqual(v interface{}) interface{} {
	switch t := v.(type) {
	case bool:
		return t == false || t == toBool(n.opValue)

	case int:
		return t <= toInt(n.opValue)

	case float64:
		return t <= toFloat(n.opValue)

	case string:
		return t <= toString(n.opValue)
	}

	return false
}

func (n *Node) CompareMatch(v interface{}) interface{} {
	return n.opValue.(*regexp.Regexp).MatchString(toString(v))
}

func (n *Node) ValueAdd(v interface{}) interface{} {
	switch t := v.(type) {
	case int:
		return t + toInt(n.opValue)

	case float64:
		return int(t + toFloat(n.opValue))
	}

	return false
}

func (n *Node) ValueSub(v interface{}) interface{} {
	switch t := v.(type) {
	case int:
		return t - toInt(n.opValue)

	case float64:
		return int(t - toFloat(n.opValue))
	}

	return false
}

func (n *Node) ValueMul(v interface{}) interface{} {
	switch t := v.(type) {
	case int:
		return t * toInt(n.opValue)

	case float64:
		return int(t * toFloat(n.opValue))
	}

	return false
}

func (n *Node) ValueDiv(v interface{}) interface{} {
	switch t := v.(type) {
	case int:
		return t / toInt(n.opValue)

	case float64:
		return int(t / toFloat(n.opValue))
	}

	return false
}

// Processor defines the JsonPath grammar processor
type Processor struct {
	*parser.BaseJsonpathListener

	Nodes  []*Node
	errors bool
}

// NewProcessor creates a new grammar Processor
func NewProcessor() *Processor {
	return &Processor{}
}

// Reset resets the processor state, making it ready to process another input
func (p *Processor) Reset() {
	p.errors = false
}

func (p *Processor) addNode(n *Node) {
	p.Nodes = append(p.Nodes, n)
}

//
// VisitErrorNode is called when there is an error
//
func (p *Processor) VisitErrorNode(node antlr.ErrorNode) {
	p.errors = true
}

// EnterPath is called when production path is entered.
// This may be useful for debugging
// func (p *Processor) EnterPath(ctx *parser.PathContext) {
//        fmt.Println("PATH", ctx.GetText())
//}

//
// ExitDotExpr is called when production dotExpr is exited.
//
func (p *Processor) ExitDotExpr(ctx *parser.DotExprContext) {
	if p.errors {
		return
	}

	names := set_type{}

	if ctx.STAR() != nil {
		names.Set(TOKEN_ANY) // any value
	}

	if ctx.Identifier() != nil {
		names.Set(ctx.Identifier().GetText())
	}

	if ctx.Length() != nil {
		names.Set(TOKEN_LENGTH)
	}

	if ctx.DOTS().GetText() == ".." {
		p.addNode(&Node{nodeType: DESCENDANT, names: names})
	} else {
		p.addNode(&Node{nodeType: CHILD, names: names})
	}
}

//
// ExitItemsExpr is called when production itemsExpr is exited.
//
func (p *Processor) ExitItemsExpr(ctx *parser.ItemsExprContext) {
	if p.errors {
		return
	}

	n := Node{nodeType: ARRAY_ITEMS}

	for _, i := range ctx.AllINT() {
		n.indices = append(n.indices, asInt(i.GetSymbol(), -1))
	}

	p.addNode(&n)
}

//
// ExitRangeExpr is called when production rangeExpr is exited.
//
func (p *Processor) ExitRangeExpr(ctx *parser.RangeExprContext) {
	if p.errors {
		return
	}

	p.addNode(&Node{
		nodeType: ARRAY_RANGE,
		start:    asInt(ctx.GetStartIndex(), 0),
		end:      asInt(ctx.GetEndIndex(), MAX_RANGE),
		step:     asInt(ctx.GetStep(), 1)})
}

//
// ExitNamesExpr is called when production namesExpr is exited.
//
func (p *Processor) ExitNamesExpr(ctx *parser.NamesExprContext) {
	if p.errors {
		return
	}

	n := &Node{nodeType: CHILD, names: set_type{}}

	last := len(p.Nodes) - 1
	if last >= 0 && len(p.Nodes[last].names) == 0 {
		n, p.Nodes = p.Nodes[last], p.Nodes[:last]
	}

	for _, i := range ctx.AllQUOTED() {
		n.names.Set(strings.Trim(i.GetSymbol().GetText(), "'"))
	}

	p.addNode(n)
}

//
// ExitStarExpr is called when production starExpr is exited.
//
func (p *Processor) ExitStarExpr(ctx *parser.StarExprContext) {
	if p.errors {
		return
	}

	last := len(p.Nodes) - 1
	if last >= 0 && len(p.Nodes[last].names) == 0 {
		p.Nodes[last].names.Set(TOKEN_ANY)
	} else {
		p.addNode(&Node{nodeType: ALL_ITEMS})
	}
}

//
// ExitFilterExpr is called when production filterExpr is exited.
//
func (p *Processor) ExitFilterExpr(ctx *parser.FilterExprContext) {
	if p.errors {
		return
	}

	p.addQueryNode(FILTER_EXPR, ctx.QueryExpr())
}

//
// ExitScriptExpr is called when production scriptExpr is exited.
//
func (p *Processor) ExitScriptExpr(ctx *parser.ScriptExprContext) {
	if p.errors {
		return
	}

	p.addExprNode(SCRIPT_EXPR, ctx.ValueExpr())
}

func (p *Processor) addQueryNode(t nodeType, q parser.IQueryExprContext) {
	n := Node{nodeType: t}

	if q.GetExists() != nil {
		n.opId = q.GetExists().GetText()
		n.op = n.CompareExists
		n.opName = "exists"
	} else {
		n.opId = q.GetName().GetText()
		n.opName = q.GetOp().GetText()

		switch {
		case q.(*parser.QueryExprContext).QUOTED() != nil:
			n.opValue = strings.Trim(q.GetValue().GetText(), "'")

		case n.opName == TOKEN_REGEX:

		case q.GetValue().GetText() == "true":
			n.opValue = true

		case q.GetValue().GetText() == "false":
			n.opValue = false

		case q.GetValue().GetText() == "null":
			n.opValue = nil

		default:
			n.opValue = asFloat(q.GetValue(), 0.0)
		}

		switch n.opName {
		case "==":
			n.op = n.CompareEqual
		case "!=":
			n.op = n.CompareNotEqual
		case ">":
			n.op = n.CompareGreater
		case "<":
			n.op = n.CompareLess
		case ">=":
			n.op = n.CompareGreaterEqual
		case "<=":
			n.op = n.CompareLessEqual
		case TOKEN_REGEX:
			n.op = n.CompareMatch

			re := q.GetValue().GetText()[1:]   // skip first '/'
			last := strings.LastIndex(re, "/") // find last '/'
			options := re[last+1:]
			re = re[:last]

			if options != "" {
				re = fmt.Sprintf("(?%v:%v)", options, re)
			} else {
				re = fmt.Sprintf("(%v)", re)
			}

			rc, err := regexp.Compile(re)
			if err != nil {
				fmt.Printf("REGEX /%v/%v: %v\n", re, options, err)
				p.errors = true
				return
			}

			n.opValue = rc
		}
	}

	p.addNode(&n)
}

func (p *Processor) addExprNode(t nodeType, q parser.IValueExprContext) {
	n := Node{nodeType: t}

	n.opId = q.GetName().GetText()
	n.opName = q.GetOp().GetText()
	n.opValue = asFloat(q.GetValue(), 0.0)

	if q.(*parser.ValueExprContext).Length() != nil {
		n.opId = TOKEN_LENGTH
	}

	switch n.opName {
	case "+":
		n.op = n.ValueAdd
	case "-":
		n.op = n.ValueSub
	case "*":
		n.op = n.ValueMul
	case "/":
		n.op = n.ValueDiv
	}

	p.addNode(&n)
}

//
// Parse parses a JsonPath expression
//
func (p *Processor) Parse(expr string) bool {
	input := antlr.NewInputStream(expr)
	lexer := parser.NewJsonpathLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewJsonpathParser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	antlr.ParseTreeWalkerDefault.Walk(p, parser.Jsonpath())
	return !p.errors
}

func (p *Processor) find(names set_type, j interface{}, enhanced bool) (ret []interface{}) {
	//fmt.Println("find", names.List(), j)

	var a []interface{}

	addele := true
	any := names.Contains(TOKEN_ANY) || len(names) == 0

	if aj, ok := j.(array_type); ok {
		a = aj
	} else {
		a = array_type{j}
		addele = false
	}

	for _, c := range a {
		if addele && any {
			ret = append(ret, c)
		}

		switch t := c.(type) {
		case map_type:
			l := len(names)
			mret := map_type{}

			addresult := func(k string, v interface{}) {
				if enhanced && l > 1 {
					mret[k] = v
				} else {
					ret = append(ret, v)
				}
			}

			for k, v := range t {
				if names.Contains(k) {
					addresult(k, v)
				} else {
					if any {
						addresult(k, v)
					}
				}

				switch v.(type) {
				case map_type, array_type:
					res := p.find(names, v, enhanced)
					ret = append(ret, res...)
				}
			}

			if len(mret) > 0 {
				ret = append(ret, mret)
			}

		case array_type:
			res := p.find(names, t, enhanced)
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
func (p *Processor) Process(v interface{}, options ProcessOptions) interface{} {
	if p.errors {
		return nil
	}

	enhanced := (options & Enhanced) == Enhanced
	collapse := (options & Collapse) == Collapse

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
			} else if n.names.Contains(TOKEN_LENGTH) {
				res = append(res, getLength(a))
				v = res
				continue
			}

			for _, c := range a {
				if n.names.Contains(TOKEN_LENGTH) {
					res = append(res, getLength(a))
					continue
				}

				if m, ok := c.(map_type); ok {
					l := len(n.names)
					mres := map_type{}

					addresult := func(k string, v interface{}) {
						if enhanced && l > 1 {
							mres[k] = v
						} else {
							res = append(res, v)
						}
					}

					if n.names.Contains(TOKEN_ANY) {
						for mk, mv := range m {
							addresult(mk, mv)
						}
					} else {
						for nn := range n.names {
							if v, ok := m[nn]; ok {
								addresult(nn, v)
							}
						}
					}

					if len(mres) > 0 {
						res = append(res, mres)
					}

				} else if n.names.Contains(TOKEN_ANY) {
					res = append(res, c)
				}
			}

			if len(res) == 1 {
				v = res[0]
			} else {
				v = res
			}

		case DESCENDANT:
			v = p.find(n.names, v, enhanced)
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

		case FILTER_EXPR:
			res := array_type{}
			a, err := j.Array()
			if err != nil {
				a = array_type{j.Data()}
			}

			for _, c := range a {
				if m, ok := c.(map_type); ok {
					if v, ok := m[n.opId]; ok {
						if n.op(v).(bool) {
							res = append(res, m)
						}
					}
				}
			}

			if len(res) == 1 {
				v = res[0]
			} else {
				v = res
			}

		case SCRIPT_EXPR:
			a, err := j.Array()
			if err != nil {
				fmt.Printf("expected array, got %v", j.Data())
				v = nil
				continue
			}

			var res int

			if n.opId == TOKEN_LENGTH {
				res = len(a)
			}

			i := n.op(res).(int)
			v, _ = atIndex(i, a)
		}
	}

	if a, ok := v.(array_type); ok {
		if collapse && len(a) == 1 {
			v = a[0]
		}
	}

	return v
}
