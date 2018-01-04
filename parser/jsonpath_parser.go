// Code generated from Jsonpath.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Jsonpath
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 26, 117,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 6, 3, 34, 10, 3, 13,
	3, 14, 3, 35, 3, 4, 3, 4, 5, 4, 40, 10, 4, 3, 5, 3, 5, 3, 5, 5, 5, 45,
	10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 54, 10, 6, 3, 6,
	3, 6, 3, 7, 5, 7, 59, 10, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 65, 10, 7, 5,
	7, 67, 10, 7, 3, 8, 3, 8, 3, 8, 7, 8, 72, 10, 8, 12, 8, 14, 8, 75, 11,
	8, 3, 9, 3, 9, 3, 9, 7, 9, 80, 10, 9, 12, 9, 14, 9, 83, 11, 9, 3, 10, 3,
	10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 105, 10,
	13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 115,
	10, 14, 3, 14, 2, 2, 15, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
	2, 5, 4, 2, 18, 19, 25, 25, 4, 2, 12, 14, 20, 22, 3, 2, 20, 21, 2, 120,
	2, 28, 3, 2, 2, 2, 4, 33, 3, 2, 2, 2, 6, 39, 3, 2, 2, 2, 8, 44, 3, 2, 2,
	2, 10, 46, 3, 2, 2, 2, 12, 58, 3, 2, 2, 2, 14, 68, 3, 2, 2, 2, 16, 76,
	3, 2, 2, 2, 18, 84, 3, 2, 2, 2, 20, 86, 3, 2, 2, 2, 22, 90, 3, 2, 2, 2,
	24, 104, 3, 2, 2, 2, 26, 114, 3, 2, 2, 2, 28, 29, 7, 3, 2, 2, 29, 30, 5,
	4, 3, 2, 30, 31, 7, 2, 2, 3, 31, 3, 3, 2, 2, 2, 32, 34, 5, 6, 4, 2, 33,
	32, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 33, 3, 2, 2, 2, 35, 36, 3, 2, 2,
	2, 36, 5, 3, 2, 2, 2, 37, 40, 5, 8, 5, 2, 38, 40, 5, 10, 6, 2, 39, 37,
	3, 2, 2, 2, 39, 38, 3, 2, 2, 2, 40, 7, 3, 2, 2, 2, 41, 45, 7, 24, 2, 2,
	42, 43, 7, 24, 2, 2, 43, 45, 9, 2, 2, 2, 44, 41, 3, 2, 2, 2, 44, 42, 3,
	2, 2, 2, 45, 9, 3, 2, 2, 2, 46, 53, 7, 4, 2, 2, 47, 54, 5, 18, 10, 2, 48,
	54, 5, 12, 7, 2, 49, 54, 5, 14, 8, 2, 50, 54, 5, 16, 9, 2, 51, 54, 5, 20,
	11, 2, 52, 54, 5, 22, 12, 2, 53, 47, 3, 2, 2, 2, 53, 48, 3, 2, 2, 2, 53,
	49, 3, 2, 2, 2, 53, 50, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 53, 52, 3, 2, 2,
	2, 53, 54, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 56, 7, 5, 2, 2, 56, 11,
	3, 2, 2, 2, 57, 59, 7, 20, 2, 2, 58, 57, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2,
	59, 60, 3, 2, 2, 2, 60, 66, 7, 6, 2, 2, 61, 64, 7, 20, 2, 2, 62, 63, 7,
	6, 2, 2, 63, 65, 7, 20, 2, 2, 64, 62, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65,
	67, 3, 2, 2, 2, 66, 61, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 13, 3, 2, 2,
	2, 68, 73, 7, 20, 2, 2, 69, 70, 7, 7, 2, 2, 70, 72, 7, 20, 2, 2, 71, 69,
	3, 2, 2, 2, 72, 75, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2,
	74, 15, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 76, 81, 7, 22, 2, 2, 77, 78, 7,
	7, 2, 2, 78, 80, 7, 22, 2, 2, 79, 77, 3, 2, 2, 2, 80, 83, 3, 2, 2, 2, 81,
	79, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 17, 3, 2, 2, 2, 83, 81, 3, 2, 2,
	2, 84, 85, 7, 25, 2, 2, 85, 19, 3, 2, 2, 2, 86, 87, 7, 8, 2, 2, 87, 88,
	5, 24, 13, 2, 88, 89, 7, 9, 2, 2, 89, 21, 3, 2, 2, 2, 90, 91, 7, 10, 2,
	2, 91, 92, 5, 26, 14, 2, 92, 93, 7, 9, 2, 2, 93, 23, 3, 2, 2, 2, 94, 95,
	7, 11, 2, 2, 95, 105, 7, 18, 2, 2, 96, 97, 7, 11, 2, 2, 97, 98, 7, 18,
	2, 2, 98, 99, 7, 17, 2, 2, 99, 105, 9, 3, 2, 2, 100, 101, 7, 11, 2, 2,
	101, 102, 7, 18, 2, 2, 102, 103, 7, 15, 2, 2, 103, 105, 7, 23, 2, 2, 104,
	94, 3, 2, 2, 2, 104, 96, 3, 2, 2, 2, 104, 100, 3, 2, 2, 2, 105, 25, 3,
	2, 2, 2, 106, 107, 7, 11, 2, 2, 107, 108, 7, 18, 2, 2, 108, 109, 7, 16,
	2, 2, 109, 115, 9, 4, 2, 2, 110, 111, 7, 11, 2, 2, 111, 112, 7, 19, 2,
	2, 112, 113, 7, 16, 2, 2, 113, 115, 9, 4, 2, 2, 114, 106, 3, 2, 2, 2, 114,
	110, 3, 2, 2, 2, 115, 27, 3, 2, 2, 2, 13, 35, 39, 44, 53, 58, 64, 66, 73,
	81, 104, 114,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'$'", "'['", "']'", "':'", "','", "'?('", "')'", "'('", "'@.'", "'true'",
	"'false'", "'null'", "'=~'", "", "", "", "'length()'", "", "", "", "",
	"", "'*'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "OP", "COMP", "Identifier",
	"Length", "INT", "NUMBER", "QUOTED", "REGEX", "DOTS", "STAR", "WS",
}

var ruleNames = []string{
	"jsonpath", "path", "nodeExpr", "dotExpr", "selectExpr", "rangeExpr", "itemsExpr",
	"namesExpr", "starExpr", "filterExpr", "scriptExpr", "queryExpr", "valueExpr",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type JsonpathParser struct {
	*antlr.BaseParser
}

func NewJsonpathParser(input antlr.TokenStream) *JsonpathParser {
	this := new(JsonpathParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Jsonpath.g4"

	return this
}

// JsonpathParser tokens.
const (
	JsonpathParserEOF        = antlr.TokenEOF
	JsonpathParserT__0       = 1
	JsonpathParserT__1       = 2
	JsonpathParserT__2       = 3
	JsonpathParserT__3       = 4
	JsonpathParserT__4       = 5
	JsonpathParserT__5       = 6
	JsonpathParserT__6       = 7
	JsonpathParserT__7       = 8
	JsonpathParserT__8       = 9
	JsonpathParserT__9       = 10
	JsonpathParserT__10      = 11
	JsonpathParserT__11      = 12
	JsonpathParserT__12      = 13
	JsonpathParserOP         = 14
	JsonpathParserCOMP       = 15
	JsonpathParserIdentifier = 16
	JsonpathParserLength     = 17
	JsonpathParserINT        = 18
	JsonpathParserNUMBER     = 19
	JsonpathParserQUOTED     = 20
	JsonpathParserREGEX      = 21
	JsonpathParserDOTS       = 22
	JsonpathParserSTAR       = 23
	JsonpathParserWS         = 24
)

// JsonpathParser rules.
const (
	JsonpathParserRULE_jsonpath   = 0
	JsonpathParserRULE_path       = 1
	JsonpathParserRULE_nodeExpr   = 2
	JsonpathParserRULE_dotExpr    = 3
	JsonpathParserRULE_selectExpr = 4
	JsonpathParserRULE_rangeExpr  = 5
	JsonpathParserRULE_itemsExpr  = 6
	JsonpathParserRULE_namesExpr  = 7
	JsonpathParserRULE_starExpr   = 8
	JsonpathParserRULE_filterExpr = 9
	JsonpathParserRULE_scriptExpr = 10
	JsonpathParserRULE_queryExpr  = 11
	JsonpathParserRULE_valueExpr  = 12
)

// IJsonpathContext is an interface to support dynamic dispatch.
type IJsonpathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJsonpathContext differentiates from other interfaces.
	IsJsonpathContext()
}

type JsonpathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonpathContext() *JsonpathContext {
	var p = new(JsonpathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_jsonpath
	return p
}

func (*JsonpathContext) IsJsonpathContext() {}

func NewJsonpathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonpathContext {
	var p = new(JsonpathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_jsonpath

	return p
}

func (s *JsonpathContext) GetParser() antlr.Parser { return s.parser }

func (s *JsonpathContext) Path() IPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *JsonpathContext) EOF() antlr.TerminalNode {
	return s.GetToken(JsonpathParserEOF, 0)
}

func (s *JsonpathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonpathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JsonpathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterJsonpath(s)
	}
}

func (s *JsonpathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitJsonpath(s)
	}
}

func (p *JsonpathParser) Jsonpath() (localctx IJsonpathContext) {
	localctx = NewJsonpathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JsonpathParserRULE_jsonpath)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.Match(JsonpathParserT__0)
	}
	{
		p.SetState(27)
		p.Path()
	}
	{
		p.SetState(28)
		p.Match(JsonpathParserEOF)
	}

	return localctx
}

// IPathContext is an interface to support dynamic dispatch.
type IPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPathContext differentiates from other interfaces.
	IsPathContext()
}

type PathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathContext() *PathContext {
	var p = new(PathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_path
	return p
}

func (*PathContext) IsPathContext() {}

func NewPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathContext {
	var p = new(PathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_path

	return p
}

func (s *PathContext) GetParser() antlr.Parser { return s.parser }

func (s *PathContext) AllNodeExpr() []INodeExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INodeExprContext)(nil)).Elem())
	var tst = make([]INodeExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INodeExprContext)
		}
	}

	return tst
}

func (s *PathContext) NodeExpr(i int) INodeExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INodeExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INodeExprContext)
}

func (s *PathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterPath(s)
	}
}

func (s *PathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitPath(s)
	}
}

func (p *JsonpathParser) Path() (localctx IPathContext) {
	localctx = NewPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JsonpathParserRULE_path)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == JsonpathParserT__1 || _la == JsonpathParserDOTS {
		{
			p.SetState(30)
			p.NodeExpr()
		}

		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// INodeExprContext is an interface to support dynamic dispatch.
type INodeExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNodeExprContext differentiates from other interfaces.
	IsNodeExprContext()
}

type NodeExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNodeExprContext() *NodeExprContext {
	var p = new(NodeExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_nodeExpr
	return p
}

func (*NodeExprContext) IsNodeExprContext() {}

func NewNodeExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NodeExprContext {
	var p = new(NodeExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_nodeExpr

	return p
}

func (s *NodeExprContext) GetParser() antlr.Parser { return s.parser }

func (s *NodeExprContext) DotExpr() IDotExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDotExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDotExprContext)
}

func (s *NodeExprContext) SelectExpr() ISelectExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectExprContext)
}

func (s *NodeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NodeExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NodeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterNodeExpr(s)
	}
}

func (s *NodeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitNodeExpr(s)
	}
}

func (p *JsonpathParser) NodeExpr() (localctx INodeExprContext) {
	localctx = NewNodeExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JsonpathParserRULE_nodeExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(37)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JsonpathParserDOTS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(35)
			p.DotExpr()
		}

	case JsonpathParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(36)
			p.SelectExpr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDotExprContext is an interface to support dynamic dispatch.
type IDotExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDotExprContext differentiates from other interfaces.
	IsDotExprContext()
}

type DotExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDotExprContext() *DotExprContext {
	var p = new(DotExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_dotExpr
	return p
}

func (*DotExprContext) IsDotExprContext() {}

func NewDotExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DotExprContext {
	var p = new(DotExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_dotExpr

	return p
}

func (s *DotExprContext) GetParser() antlr.Parser { return s.parser }

func (s *DotExprContext) DOTS() antlr.TerminalNode {
	return s.GetToken(JsonpathParserDOTS, 0)
}

func (s *DotExprContext) Identifier() antlr.TerminalNode {
	return s.GetToken(JsonpathParserIdentifier, 0)
}

func (s *DotExprContext) STAR() antlr.TerminalNode {
	return s.GetToken(JsonpathParserSTAR, 0)
}

func (s *DotExprContext) Length() antlr.TerminalNode {
	return s.GetToken(JsonpathParserLength, 0)
}

func (s *DotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DotExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterDotExpr(s)
	}
}

func (s *DotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitDotExpr(s)
	}
}

func (p *JsonpathParser) DotExpr() (localctx IDotExprContext) {
	localctx = NewDotExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JsonpathParserRULE_dotExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(39)
			p.Match(JsonpathParserDOTS)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(40)
			p.Match(JsonpathParserDOTS)
		}
		{
			p.SetState(41)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JsonpathParserIdentifier)|(1<<JsonpathParserLength)|(1<<JsonpathParserSTAR))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}

// ISelectExprContext is an interface to support dynamic dispatch.
type ISelectExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectExprContext differentiates from other interfaces.
	IsSelectExprContext()
}

type SelectExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectExprContext() *SelectExprContext {
	var p = new(SelectExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_selectExpr
	return p
}

func (*SelectExprContext) IsSelectExprContext() {}

func NewSelectExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectExprContext {
	var p = new(SelectExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_selectExpr

	return p
}

func (s *SelectExprContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectExprContext) StarExpr() IStarExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStarExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStarExprContext)
}

func (s *SelectExprContext) RangeExpr() IRangeExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangeExprContext)
}

func (s *SelectExprContext) ItemsExpr() IItemsExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IItemsExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IItemsExprContext)
}

func (s *SelectExprContext) NamesExpr() INamesExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamesExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamesExprContext)
}

func (s *SelectExprContext) FilterExpr() IFilterExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilterExprContext)
}

func (s *SelectExprContext) ScriptExpr() IScriptExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScriptExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScriptExprContext)
}

func (s *SelectExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterSelectExpr(s)
	}
}

func (s *SelectExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitSelectExpr(s)
	}
}

func (p *JsonpathParser) SelectExpr() (localctx ISelectExprContext) {
	localctx = NewSelectExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JsonpathParserRULE_selectExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Match(JsonpathParserT__1)
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(45)
			p.StarExpr()
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(46)
			p.RangeExpr()
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 3 {
		{
			p.SetState(47)
			p.ItemsExpr()
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 4 {
		{
			p.SetState(48)
			p.NamesExpr()
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 5 {
		{
			p.SetState(49)
			p.FilterExpr()
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 6 {
		{
			p.SetState(50)
			p.ScriptExpr()
		}

	}
	{
		p.SetState(53)
		p.Match(JsonpathParserT__2)
	}

	return localctx
}

// IRangeExprContext is an interface to support dynamic dispatch.
type IRangeExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStartIndex returns the startIndex token.
	GetStartIndex() antlr.Token

	// GetEndIndex returns the endIndex token.
	GetEndIndex() antlr.Token

	// GetStep returns the step token.
	GetStep() antlr.Token

	// SetStartIndex sets the startIndex token.
	SetStartIndex(antlr.Token)

	// SetEndIndex sets the endIndex token.
	SetEndIndex(antlr.Token)

	// SetStep sets the step token.
	SetStep(antlr.Token)

	// IsRangeExprContext differentiates from other interfaces.
	IsRangeExprContext()
}

type RangeExprContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	startIndex antlr.Token
	endIndex   antlr.Token
	step       antlr.Token
}

func NewEmptyRangeExprContext() *RangeExprContext {
	var p = new(RangeExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_rangeExpr
	return p
}

func (*RangeExprContext) IsRangeExprContext() {}

func NewRangeExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeExprContext {
	var p = new(RangeExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_rangeExpr

	return p
}

func (s *RangeExprContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeExprContext) GetStartIndex() antlr.Token { return s.startIndex }

func (s *RangeExprContext) GetEndIndex() antlr.Token { return s.endIndex }

func (s *RangeExprContext) GetStep() antlr.Token { return s.step }

func (s *RangeExprContext) SetStartIndex(v antlr.Token) { s.startIndex = v }

func (s *RangeExprContext) SetEndIndex(v antlr.Token) { s.endIndex = v }

func (s *RangeExprContext) SetStep(v antlr.Token) { s.step = v }

func (s *RangeExprContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(JsonpathParserINT)
}

func (s *RangeExprContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(JsonpathParserINT, i)
}

func (s *RangeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterRangeExpr(s)
	}
}

func (s *RangeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitRangeExpr(s)
	}
}

func (p *JsonpathParser) RangeExpr() (localctx IRangeExprContext) {
	localctx = NewRangeExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JsonpathParserRULE_rangeExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JsonpathParserINT {
		{
			p.SetState(55)

			var _m = p.Match(JsonpathParserINT)

			localctx.(*RangeExprContext).startIndex = _m
		}

	}
	{
		p.SetState(58)
		p.Match(JsonpathParserT__3)
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JsonpathParserINT {
		{
			p.SetState(59)

			var _m = p.Match(JsonpathParserINT)

			localctx.(*RangeExprContext).endIndex = _m
		}
		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonpathParserT__3 {
			{
				p.SetState(60)
				p.Match(JsonpathParserT__3)
			}
			{
				p.SetState(61)

				var _m = p.Match(JsonpathParserINT)

				localctx.(*RangeExprContext).step = _m
			}

		}

	}

	return localctx
}

// IItemsExprContext is an interface to support dynamic dispatch.
type IItemsExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsItemsExprContext differentiates from other interfaces.
	IsItemsExprContext()
}

type ItemsExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyItemsExprContext() *ItemsExprContext {
	var p = new(ItemsExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_itemsExpr
	return p
}

func (*ItemsExprContext) IsItemsExprContext() {}

func NewItemsExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ItemsExprContext {
	var p = new(ItemsExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_itemsExpr

	return p
}

func (s *ItemsExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ItemsExprContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(JsonpathParserINT)
}

func (s *ItemsExprContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(JsonpathParserINT, i)
}

func (s *ItemsExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ItemsExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ItemsExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterItemsExpr(s)
	}
}

func (s *ItemsExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitItemsExpr(s)
	}
}

func (p *JsonpathParser) ItemsExpr() (localctx IItemsExprContext) {
	localctx = NewItemsExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, JsonpathParserRULE_itemsExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Match(JsonpathParserINT)
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == JsonpathParserT__4 {
		{
			p.SetState(67)
			p.Match(JsonpathParserT__4)
		}
		{
			p.SetState(68)
			p.Match(JsonpathParserINT)
		}

		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// INamesExprContext is an interface to support dynamic dispatch.
type INamesExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamesExprContext differentiates from other interfaces.
	IsNamesExprContext()
}

type NamesExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamesExprContext() *NamesExprContext {
	var p = new(NamesExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_namesExpr
	return p
}

func (*NamesExprContext) IsNamesExprContext() {}

func NewNamesExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamesExprContext {
	var p = new(NamesExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_namesExpr

	return p
}

func (s *NamesExprContext) GetParser() antlr.Parser { return s.parser }

func (s *NamesExprContext) AllQUOTED() []antlr.TerminalNode {
	return s.GetTokens(JsonpathParserQUOTED)
}

func (s *NamesExprContext) QUOTED(i int) antlr.TerminalNode {
	return s.GetToken(JsonpathParserQUOTED, i)
}

func (s *NamesExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamesExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamesExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterNamesExpr(s)
	}
}

func (s *NamesExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitNamesExpr(s)
	}
}

func (p *JsonpathParser) NamesExpr() (localctx INamesExprContext) {
	localctx = NewNamesExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, JsonpathParserRULE_namesExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Match(JsonpathParserQUOTED)
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == JsonpathParserT__4 {
		{
			p.SetState(75)
			p.Match(JsonpathParserT__4)
		}
		{
			p.SetState(76)
			p.Match(JsonpathParserQUOTED)
		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStarExprContext is an interface to support dynamic dispatch.
type IStarExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStarExprContext differentiates from other interfaces.
	IsStarExprContext()
}

type StarExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStarExprContext() *StarExprContext {
	var p = new(StarExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_starExpr
	return p
}

func (*StarExprContext) IsStarExprContext() {}

func NewStarExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StarExprContext {
	var p = new(StarExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_starExpr

	return p
}

func (s *StarExprContext) GetParser() antlr.Parser { return s.parser }

func (s *StarExprContext) STAR() antlr.TerminalNode {
	return s.GetToken(JsonpathParserSTAR, 0)
}

func (s *StarExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StarExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StarExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterStarExpr(s)
	}
}

func (s *StarExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitStarExpr(s)
	}
}

func (p *JsonpathParser) StarExpr() (localctx IStarExprContext) {
	localctx = NewStarExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, JsonpathParserRULE_starExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Match(JsonpathParserSTAR)
	}

	return localctx
}

// IFilterExprContext is an interface to support dynamic dispatch.
type IFilterExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilterExprContext differentiates from other interfaces.
	IsFilterExprContext()
}

type FilterExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterExprContext() *FilterExprContext {
	var p = new(FilterExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_filterExpr
	return p
}

func (*FilterExprContext) IsFilterExprContext() {}

func NewFilterExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterExprContext {
	var p = new(FilterExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_filterExpr

	return p
}

func (s *FilterExprContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterExprContext) QueryExpr() IQueryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryExprContext)
}

func (s *FilterExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterFilterExpr(s)
	}
}

func (s *FilterExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitFilterExpr(s)
	}
}

func (p *JsonpathParser) FilterExpr() (localctx IFilterExprContext) {
	localctx = NewFilterExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, JsonpathParserRULE_filterExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Match(JsonpathParserT__5)
	}
	{
		p.SetState(85)
		p.QueryExpr()
	}
	{
		p.SetState(86)
		p.Match(JsonpathParserT__6)
	}

	return localctx
}

// IScriptExprContext is an interface to support dynamic dispatch.
type IScriptExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScriptExprContext differentiates from other interfaces.
	IsScriptExprContext()
}

type ScriptExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScriptExprContext() *ScriptExprContext {
	var p = new(ScriptExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_scriptExpr
	return p
}

func (*ScriptExprContext) IsScriptExprContext() {}

func NewScriptExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScriptExprContext {
	var p = new(ScriptExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_scriptExpr

	return p
}

func (s *ScriptExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ScriptExprContext) ValueExpr() IValueExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueExprContext)
}

func (s *ScriptExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScriptExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScriptExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterScriptExpr(s)
	}
}

func (s *ScriptExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitScriptExpr(s)
	}
}

func (p *JsonpathParser) ScriptExpr() (localctx IScriptExprContext) {
	localctx = NewScriptExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, JsonpathParserRULE_scriptExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Match(JsonpathParserT__7)
	}
	{
		p.SetState(89)
		p.ValueExpr()
	}
	{
		p.SetState(90)
		p.Match(JsonpathParserT__6)
	}

	return localctx
}

// IQueryExprContext is an interface to support dynamic dispatch.
type IQueryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExists returns the exists token.
	GetExists() antlr.Token

	// GetName returns the name token.
	GetName() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetExists sets the exists token.
	SetExists(antlr.Token)

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsQueryExprContext differentiates from other interfaces.
	IsQueryExprContext()
}

type QueryExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	exists antlr.Token
	name   antlr.Token
	op     antlr.Token
	value  antlr.Token
}

func NewEmptyQueryExprContext() *QueryExprContext {
	var p = new(QueryExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_queryExpr
	return p
}

func (*QueryExprContext) IsQueryExprContext() {}

func NewQueryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryExprContext {
	var p = new(QueryExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_queryExpr

	return p
}

func (s *QueryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryExprContext) GetExists() antlr.Token { return s.exists }

func (s *QueryExprContext) GetName() antlr.Token { return s.name }

func (s *QueryExprContext) GetOp() antlr.Token { return s.op }

func (s *QueryExprContext) GetValue() antlr.Token { return s.value }

func (s *QueryExprContext) SetExists(v antlr.Token) { s.exists = v }

func (s *QueryExprContext) SetName(v antlr.Token) { s.name = v }

func (s *QueryExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *QueryExprContext) SetValue(v antlr.Token) { s.value = v }

func (s *QueryExprContext) Identifier() antlr.TerminalNode {
	return s.GetToken(JsonpathParserIdentifier, 0)
}

func (s *QueryExprContext) COMP() antlr.TerminalNode {
	return s.GetToken(JsonpathParserCOMP, 0)
}

func (s *QueryExprContext) INT() antlr.TerminalNode {
	return s.GetToken(JsonpathParserINT, 0)
}

func (s *QueryExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(JsonpathParserNUMBER, 0)
}

func (s *QueryExprContext) QUOTED() antlr.TerminalNode {
	return s.GetToken(JsonpathParserQUOTED, 0)
}

func (s *QueryExprContext) REGEX() antlr.TerminalNode {
	return s.GetToken(JsonpathParserREGEX, 0)
}

func (s *QueryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterQueryExpr(s)
	}
}

func (s *QueryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitQueryExpr(s)
	}
}

func (p *JsonpathParser) QueryExpr() (localctx IQueryExprContext) {
	localctx = NewQueryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, JsonpathParserRULE_queryExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.Match(JsonpathParserT__8)
		}
		{
			p.SetState(93)

			var _m = p.Match(JsonpathParserIdentifier)

			localctx.(*QueryExprContext).exists = _m
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(94)
			p.Match(JsonpathParserT__8)
		}
		{
			p.SetState(95)

			var _m = p.Match(JsonpathParserIdentifier)

			localctx.(*QueryExprContext).name = _m
		}
		{
			p.SetState(96)

			var _m = p.Match(JsonpathParserCOMP)

			localctx.(*QueryExprContext).op = _m
		}
		{
			p.SetState(97)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*QueryExprContext).value = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JsonpathParserT__9)|(1<<JsonpathParserT__10)|(1<<JsonpathParserT__11)|(1<<JsonpathParserINT)|(1<<JsonpathParserNUMBER)|(1<<JsonpathParserQUOTED))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*QueryExprContext).value = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(98)
			p.Match(JsonpathParserT__8)
		}
		{
			p.SetState(99)

			var _m = p.Match(JsonpathParserIdentifier)

			localctx.(*QueryExprContext).name = _m
		}
		{
			p.SetState(100)

			var _m = p.Match(JsonpathParserT__12)

			localctx.(*QueryExprContext).op = _m
		}
		{
			p.SetState(101)

			var _m = p.Match(JsonpathParserREGEX)

			localctx.(*QueryExprContext).value = _m
		}

	}

	return localctx
}

// IValueExprContext is an interface to support dynamic dispatch.
type IValueExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsValueExprContext differentiates from other interfaces.
	IsValueExprContext()
}

type ValueExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	op     antlr.Token
	value  antlr.Token
}

func NewEmptyValueExprContext() *ValueExprContext {
	var p = new(ValueExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_valueExpr
	return p
}

func (*ValueExprContext) IsValueExprContext() {}

func NewValueExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueExprContext {
	var p = new(ValueExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_valueExpr

	return p
}

func (s *ValueExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueExprContext) GetName() antlr.Token { return s.name }

func (s *ValueExprContext) GetOp() antlr.Token { return s.op }

func (s *ValueExprContext) GetValue() antlr.Token { return s.value }

func (s *ValueExprContext) SetName(v antlr.Token) { s.name = v }

func (s *ValueExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ValueExprContext) SetValue(v antlr.Token) { s.value = v }

func (s *ValueExprContext) Identifier() antlr.TerminalNode {
	return s.GetToken(JsonpathParserIdentifier, 0)
}

func (s *ValueExprContext) OP() antlr.TerminalNode {
	return s.GetToken(JsonpathParserOP, 0)
}

func (s *ValueExprContext) INT() antlr.TerminalNode {
	return s.GetToken(JsonpathParserINT, 0)
}

func (s *ValueExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(JsonpathParserNUMBER, 0)
}

func (s *ValueExprContext) Length() antlr.TerminalNode {
	return s.GetToken(JsonpathParserLength, 0)
}

func (s *ValueExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterValueExpr(s)
	}
}

func (s *ValueExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitValueExpr(s)
	}
}

func (p *JsonpathParser) ValueExpr() (localctx IValueExprContext) {
	localctx = NewValueExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, JsonpathParserRULE_valueExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(104)
			p.Match(JsonpathParserT__8)
		}
		{
			p.SetState(105)

			var _m = p.Match(JsonpathParserIdentifier)

			localctx.(*ValueExprContext).name = _m
		}
		{
			p.SetState(106)

			var _m = p.Match(JsonpathParserOP)

			localctx.(*ValueExprContext).op = _m
		}
		{
			p.SetState(107)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ValueExprContext).value = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == JsonpathParserINT || _la == JsonpathParserNUMBER) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ValueExprContext).value = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(108)
			p.Match(JsonpathParserT__8)
		}
		{
			p.SetState(109)

			var _m = p.Match(JsonpathParserLength)

			localctx.(*ValueExprContext).name = _m
		}
		{
			p.SetState(110)

			var _m = p.Match(JsonpathParserOP)

			localctx.(*ValueExprContext).op = _m
		}
		{
			p.SetState(111)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ValueExprContext).value = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == JsonpathParserINT || _la == JsonpathParserNUMBER) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ValueExprContext).value = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}
