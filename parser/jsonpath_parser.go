// Generated from Jsonpath.g4 by ANTLR 4.7.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 23, 105,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 6, 3, 32, 10, 3, 13, 3, 14, 3, 33,
	3, 4, 3, 4, 5, 4, 38, 10, 4, 3, 5, 3, 5, 3, 5, 5, 5, 43, 10, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 52, 10, 6, 3, 6, 3, 6, 3, 7, 5,
	7, 57, 10, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 63, 10, 7, 5, 7, 65, 10, 7,
	3, 8, 3, 8, 3, 8, 7, 8, 70, 10, 8, 12, 8, 14, 8, 73, 11, 8, 3, 9, 3, 9,
	3, 9, 7, 9, 78, 10, 9, 12, 9, 14, 9, 81, 11, 9, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 103, 10, 13, 3, 13, 2,
	2, 14, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 2, 4, 4, 2, 13, 13,
	19, 20, 3, 2, 15, 16, 2, 108, 2, 26, 3, 2, 2, 2, 4, 31, 3, 2, 2, 2, 6,
	37, 3, 2, 2, 2, 8, 42, 3, 2, 2, 2, 10, 44, 3, 2, 2, 2, 12, 56, 3, 2, 2,
	2, 14, 66, 3, 2, 2, 2, 16, 74, 3, 2, 2, 2, 18, 82, 3, 2, 2, 2, 20, 84,
	3, 2, 2, 2, 22, 88, 3, 2, 2, 2, 24, 102, 3, 2, 2, 2, 26, 27, 7, 3, 2, 2,
	27, 28, 5, 4, 3, 2, 28, 29, 7, 2, 2, 3, 29, 3, 3, 2, 2, 2, 30, 32, 5, 6,
	4, 2, 31, 30, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 31, 3, 2, 2, 2, 33, 34,
	3, 2, 2, 2, 34, 5, 3, 2, 2, 2, 35, 38, 5, 8, 5, 2, 36, 38, 5, 10, 6, 2,
	37, 35, 3, 2, 2, 2, 37, 36, 3, 2, 2, 2, 38, 7, 3, 2, 2, 2, 39, 43, 7, 18,
	2, 2, 40, 41, 7, 18, 2, 2, 41, 43, 9, 2, 2, 2, 42, 39, 3, 2, 2, 2, 42,
	40, 3, 2, 2, 2, 43, 9, 3, 2, 2, 2, 44, 51, 7, 4, 2, 2, 45, 52, 5, 18, 10,
	2, 46, 52, 5, 12, 7, 2, 47, 52, 5, 14, 8, 2, 48, 52, 5, 16, 9, 2, 49, 52,
	5, 20, 11, 2, 50, 52, 5, 22, 12, 2, 51, 45, 3, 2, 2, 2, 51, 46, 3, 2, 2,
	2, 51, 47, 3, 2, 2, 2, 51, 48, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 51, 50,
	3, 2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 54, 7, 5, 2, 2,
	54, 11, 3, 2, 2, 2, 55, 57, 7, 14, 2, 2, 56, 55, 3, 2, 2, 2, 56, 57, 3,
	2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 64, 7, 6, 2, 2, 59, 62, 7, 14, 2, 2, 60,
	61, 7, 6, 2, 2, 61, 63, 7, 14, 2, 2, 62, 60, 3, 2, 2, 2, 62, 63, 3, 2,
	2, 2, 63, 65, 3, 2, 2, 2, 64, 59, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 13,
	3, 2, 2, 2, 66, 71, 7, 14, 2, 2, 67, 68, 7, 7, 2, 2, 68, 70, 7, 14, 2,
	2, 69, 67, 3, 2, 2, 2, 70, 73, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 72,
	3, 2, 2, 2, 72, 15, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 74, 79, 7, 16, 2, 2,
	75, 76, 7, 7, 2, 2, 76, 78, 7, 16, 2, 2, 77, 75, 3, 2, 2, 2, 78, 81, 3,
	2, 2, 2, 79, 77, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 17, 3, 2, 2, 2, 81,
	79, 3, 2, 2, 2, 82, 83, 7, 19, 2, 2, 83, 19, 3, 2, 2, 2, 84, 85, 7, 8,
	2, 2, 85, 86, 5, 24, 13, 2, 86, 87, 7, 9, 2, 2, 87, 21, 3, 2, 2, 2, 88,
	89, 7, 10, 2, 2, 89, 90, 5, 24, 13, 2, 90, 91, 7, 9, 2, 2, 91, 23, 3, 2,
	2, 2, 92, 93, 7, 11, 2, 2, 93, 103, 7, 13, 2, 2, 94, 95, 7, 11, 2, 2, 95,
	96, 7, 13, 2, 2, 96, 97, 7, 22, 2, 2, 97, 103, 9, 3, 2, 2, 98, 99, 7, 11,
	2, 2, 99, 100, 7, 13, 2, 2, 100, 101, 7, 12, 2, 2, 101, 103, 7, 17, 2,
	2, 102, 92, 3, 2, 2, 2, 102, 94, 3, 2, 2, 2, 102, 98, 3, 2, 2, 2, 103,
	25, 3, 2, 2, 2, 12, 33, 37, 42, 51, 56, 62, 64, 71, 79, 102,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'$'", "'['", "']'", "':'", "','", "'?('", "')'", "'('", "'@.'", "'=~'",
	"", "", "", "", "", "", "'*'", "'length()'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "Identifier", "INT", "NUMBER",
	"QUOTED", "REGEX", "DOTS", "STAR", "Length", "OP", "COMP", "WS",
}

var ruleNames = []string{
	"jsonpath", "path", "nodeExpr", "dotExpr", "selectExpr", "rangeExpr", "itemsExpr",
	"namesExpr", "starExpr", "filterExpr", "scriptExpr", "queryExpr",
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
	JsonpathParserIdentifier = 11
	JsonpathParserINT        = 12
	JsonpathParserNUMBER     = 13
	JsonpathParserQUOTED     = 14
	JsonpathParserREGEX      = 15
	JsonpathParserDOTS       = 16
	JsonpathParserSTAR       = 17
	JsonpathParserLength     = 18
	JsonpathParserOP         = 19
	JsonpathParserCOMP       = 20
	JsonpathParserWS         = 21
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
		p.SetState(24)
		p.Match(JsonpathParserT__0)
	}
	{
		p.SetState(25)
		p.Path()
	}
	{
		p.SetState(26)
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
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == JsonpathParserT__1 || _la == JsonpathParserDOTS {
		{
			p.SetState(28)
			p.NodeExpr()
		}

		p.SetState(31)
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

	p.SetState(35)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JsonpathParserDOTS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)
			p.DotExpr()
		}

	case JsonpathParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(34)
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

	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(37)
			p.Match(JsonpathParserDOTS)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(38)
			p.Match(JsonpathParserDOTS)
		}
		p.SetState(39)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JsonpathParserIdentifier)|(1<<JsonpathParserSTAR)|(1<<JsonpathParserLength))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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
		p.SetState(42)
		p.Match(JsonpathParserT__1)
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(43)
			p.StarExpr()
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(44)
			p.RangeExpr()
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 3 {
		{
			p.SetState(45)
			p.ItemsExpr()
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 4 {
		{
			p.SetState(46)
			p.NamesExpr()
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 5 {
		{
			p.SetState(47)
			p.FilterExpr()
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 6 {
		{
			p.SetState(48)
			p.ScriptExpr()
		}

	}
	{
		p.SetState(51)
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
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JsonpathParserINT {
		{
			p.SetState(53)

			var _m = p.Match(JsonpathParserINT)

			localctx.(*RangeExprContext).startIndex = _m
		}

	}
	{
		p.SetState(56)
		p.Match(JsonpathParserT__3)
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JsonpathParserINT {
		{
			p.SetState(57)

			var _m = p.Match(JsonpathParserINT)

			localctx.(*RangeExprContext).endIndex = _m
		}
		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonpathParserT__3 {
			{
				p.SetState(58)
				p.Match(JsonpathParserT__3)
			}
			{
				p.SetState(59)

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
		p.SetState(64)
		p.Match(JsonpathParserINT)
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == JsonpathParserT__4 {
		{
			p.SetState(65)
			p.Match(JsonpathParserT__4)
		}
		{
			p.SetState(66)
			p.Match(JsonpathParserINT)
		}

		p.SetState(71)
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
		p.SetState(72)
		p.Match(JsonpathParserQUOTED)
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == JsonpathParserT__4 {
		{
			p.SetState(73)
			p.Match(JsonpathParserT__4)
		}
		{
			p.SetState(74)
			p.Match(JsonpathParserQUOTED)
		}

		p.SetState(79)
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
		p.SetState(80)
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
		p.SetState(82)
		p.Match(JsonpathParserT__5)
	}
	{
		p.SetState(83)
		p.QueryExpr()
	}
	{
		p.SetState(84)
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

func (s *ScriptExprContext) QueryExpr() IQueryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryExprContext)
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
		p.SetState(86)
		p.Match(JsonpathParserT__7)
	}
	{
		p.SetState(87)
		p.QueryExpr()
	}
	{
		p.SetState(88)
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

	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(90)
			p.Match(JsonpathParserT__8)
		}
		{
			p.SetState(91)

			var _m = p.Match(JsonpathParserIdentifier)

			localctx.(*QueryExprContext).exists = _m
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(92)
			p.Match(JsonpathParserT__8)
		}
		{
			p.SetState(93)

			var _m = p.Match(JsonpathParserIdentifier)

			localctx.(*QueryExprContext).name = _m
		}
		{
			p.SetState(94)

			var _m = p.Match(JsonpathParserCOMP)

			localctx.(*QueryExprContext).op = _m
		}
		p.SetState(95)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*QueryExprContext).value = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == JsonpathParserNUMBER || _la == JsonpathParserQUOTED) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*QueryExprContext).value = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(96)
			p.Match(JsonpathParserT__8)
		}
		{
			p.SetState(97)

			var _m = p.Match(JsonpathParserIdentifier)

			localctx.(*QueryExprContext).name = _m
		}
		{
			p.SetState(98)

			var _m = p.Match(JsonpathParserT__9)

			localctx.(*QueryExprContext).op = _m
		}
		{
			p.SetState(99)

			var _m = p.Match(JsonpathParserREGEX)

			localctx.(*QueryExprContext).value = _m
		}

	}

	return localctx
}
