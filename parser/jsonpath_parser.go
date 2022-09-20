// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Jsonpath
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type JsonpathParser struct {
	*antlr.BaseParser
}

var jsonpathParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jsonpathParserInit() {
	staticData := &jsonpathParserStaticData
	staticData.literalNames = []string{
		"", "'$'", "'['", "']'", "':'", "','", "')'", "'('", "'@.'", "'true'",
		"'false'", "'null'", "'=~'", "", "", "", "'length()'", "", "", "", "",
		"", "'*'", "'?('", "'!('",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "OP", "COMP", "Identifier",
		"Length", "INT", "NUMBER", "QUOTED", "REGEX", "DOTS", "STAR", "FilterTrue",
		"FilterFalse", "WS",
	}
	staticData.ruleNames = []string{
		"jsonpath", "path", "nodeExpr", "dotExpr", "selectExpr", "rangeExpr",
		"itemsExpr", "namesExpr", "starExpr", "filterExpr", "scriptExpr", "queryExpr",
		"valueExpr",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 25, 121, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 4, 1, 32,
		8, 1, 11, 1, 12, 1, 33, 1, 2, 1, 2, 3, 2, 38, 8, 2, 1, 3, 1, 3, 1, 3, 3,
		3, 43, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 52, 8, 4,
		1, 4, 1, 4, 1, 5, 3, 5, 57, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 63, 8,
		5, 3, 5, 65, 8, 5, 1, 6, 1, 6, 1, 6, 5, 6, 70, 8, 6, 10, 6, 12, 6, 73,
		9, 6, 1, 7, 1, 7, 1, 7, 5, 7, 78, 8, 7, 10, 7, 12, 7, 81, 9, 7, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 93, 8, 9, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 3, 11, 109, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 3, 12, 119, 8, 12, 1, 12, 0, 0, 13, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 0, 3, 2, 0, 15, 16, 22, 22, 2, 0, 9, 11,
		17, 19, 1, 0, 17, 18, 125, 0, 26, 1, 0, 0, 0, 2, 31, 1, 0, 0, 0, 4, 37,
		1, 0, 0, 0, 6, 42, 1, 0, 0, 0, 8, 44, 1, 0, 0, 0, 10, 56, 1, 0, 0, 0, 12,
		66, 1, 0, 0, 0, 14, 74, 1, 0, 0, 0, 16, 82, 1, 0, 0, 0, 18, 92, 1, 0, 0,
		0, 20, 94, 1, 0, 0, 0, 22, 108, 1, 0, 0, 0, 24, 118, 1, 0, 0, 0, 26, 27,
		5, 1, 0, 0, 27, 28, 3, 2, 1, 0, 28, 29, 5, 0, 0, 1, 29, 1, 1, 0, 0, 0,
		30, 32, 3, 4, 2, 0, 31, 30, 1, 0, 0, 0, 32, 33, 1, 0, 0, 0, 33, 31, 1,
		0, 0, 0, 33, 34, 1, 0, 0, 0, 34, 3, 1, 0, 0, 0, 35, 38, 3, 6, 3, 0, 36,
		38, 3, 8, 4, 0, 37, 35, 1, 0, 0, 0, 37, 36, 1, 0, 0, 0, 38, 5, 1, 0, 0,
		0, 39, 43, 5, 21, 0, 0, 40, 41, 5, 21, 0, 0, 41, 43, 7, 0, 0, 0, 42, 39,
		1, 0, 0, 0, 42, 40, 1, 0, 0, 0, 43, 7, 1, 0, 0, 0, 44, 51, 5, 2, 0, 0,
		45, 52, 3, 16, 8, 0, 46, 52, 3, 10, 5, 0, 47, 52, 3, 12, 6, 0, 48, 52,
		3, 14, 7, 0, 49, 52, 3, 18, 9, 0, 50, 52, 3, 20, 10, 0, 51, 45, 1, 0, 0,
		0, 51, 46, 1, 0, 0, 0, 51, 47, 1, 0, 0, 0, 51, 48, 1, 0, 0, 0, 51, 49,
		1, 0, 0, 0, 51, 50, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0,
		53, 54, 5, 3, 0, 0, 54, 9, 1, 0, 0, 0, 55, 57, 5, 17, 0, 0, 56, 55, 1,
		0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 64, 5, 4, 0, 0, 59,
		62, 5, 17, 0, 0, 60, 61, 5, 4, 0, 0, 61, 63, 5, 17, 0, 0, 62, 60, 1, 0,
		0, 0, 62, 63, 1, 0, 0, 0, 63, 65, 1, 0, 0, 0, 64, 59, 1, 0, 0, 0, 64, 65,
		1, 0, 0, 0, 65, 11, 1, 0, 0, 0, 66, 71, 5, 17, 0, 0, 67, 68, 5, 5, 0, 0,
		68, 70, 5, 17, 0, 0, 69, 67, 1, 0, 0, 0, 70, 73, 1, 0, 0, 0, 71, 69, 1,
		0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 13, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 74,
		79, 5, 19, 0, 0, 75, 76, 5, 5, 0, 0, 76, 78, 5, 19, 0, 0, 77, 75, 1, 0,
		0, 0, 78, 81, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 15,
		1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 82, 83, 5, 22, 0, 0, 83, 17, 1, 0, 0, 0,
		84, 85, 5, 23, 0, 0, 85, 86, 3, 22, 11, 0, 86, 87, 5, 6, 0, 0, 87, 93,
		1, 0, 0, 0, 88, 89, 5, 24, 0, 0, 89, 90, 3, 22, 11, 0, 90, 91, 5, 6, 0,
		0, 91, 93, 1, 0, 0, 0, 92, 84, 1, 0, 0, 0, 92, 88, 1, 0, 0, 0, 93, 19,
		1, 0, 0, 0, 94, 95, 5, 7, 0, 0, 95, 96, 3, 24, 12, 0, 96, 97, 5, 6, 0,
		0, 97, 21, 1, 0, 0, 0, 98, 99, 5, 8, 0, 0, 99, 109, 5, 15, 0, 0, 100, 101,
		5, 8, 0, 0, 101, 102, 5, 15, 0, 0, 102, 103, 5, 14, 0, 0, 103, 109, 7,
		1, 0, 0, 104, 105, 5, 8, 0, 0, 105, 106, 5, 15, 0, 0, 106, 107, 5, 12,
		0, 0, 107, 109, 5, 20, 0, 0, 108, 98, 1, 0, 0, 0, 108, 100, 1, 0, 0, 0,
		108, 104, 1, 0, 0, 0, 109, 23, 1, 0, 0, 0, 110, 111, 5, 8, 0, 0, 111, 112,
		5, 15, 0, 0, 112, 113, 5, 13, 0, 0, 113, 119, 7, 2, 0, 0, 114, 115, 5,
		8, 0, 0, 115, 116, 5, 16, 0, 0, 116, 117, 5, 13, 0, 0, 117, 119, 7, 2,
		0, 0, 118, 110, 1, 0, 0, 0, 118, 114, 1, 0, 0, 0, 119, 25, 1, 0, 0, 0,
		12, 33, 37, 42, 51, 56, 62, 64, 71, 79, 92, 108, 118,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// JsonpathParserInit initializes any static state used to implement JsonpathParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewJsonpathParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func JsonpathParserInit() {
	staticData := &jsonpathParserStaticData
	staticData.once.Do(jsonpathParserInit)
}

// NewJsonpathParser produces a new parser instance for the optional input antlr.TokenStream.
func NewJsonpathParser(input antlr.TokenStream) *JsonpathParser {
	JsonpathParserInit()
	this := new(JsonpathParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &jsonpathParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// JsonpathParser tokens.
const (
	JsonpathParserEOF         = antlr.TokenEOF
	JsonpathParserT__0        = 1
	JsonpathParserT__1        = 2
	JsonpathParserT__2        = 3
	JsonpathParserT__3        = 4
	JsonpathParserT__4        = 5
	JsonpathParserT__5        = 6
	JsonpathParserT__6        = 7
	JsonpathParserT__7        = 8
	JsonpathParserT__8        = 9
	JsonpathParserT__9        = 10
	JsonpathParserT__10       = 11
	JsonpathParserT__11       = 12
	JsonpathParserOP          = 13
	JsonpathParserCOMP        = 14
	JsonpathParserIdentifier  = 15
	JsonpathParserLength      = 16
	JsonpathParserINT         = 17
	JsonpathParserNUMBER      = 18
	JsonpathParserQUOTED      = 19
	JsonpathParserREGEX       = 20
	JsonpathParserDOTS        = 21
	JsonpathParserSTAR        = 22
	JsonpathParserFilterTrue  = 23
	JsonpathParserFilterFalse = 24
	JsonpathParserWS          = 25
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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	this := p
	_ = this

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INodeExprContext); ok {
			len++
		}
	}

	tst := make([]INodeExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INodeExprContext); ok {
			tst[i] = t.(INodeExprContext)
			i++
		}
	}

	return tst
}

func (s *PathContext) NodeExpr(i int) INodeExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INodeExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
	this := p
	_ = this

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDotExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDotExprContext)
}

func (s *NodeExprContext) SelectExpr() ISelectExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	this := p
	_ = this

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
	this := p
	_ = this

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

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4292608) != 0) {
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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStarExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStarExprContext)
}

func (s *SelectExprContext) RangeExpr() IRangeExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangeExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangeExprContext)
}

func (s *SelectExprContext) ItemsExpr() IItemsExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IItemsExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IItemsExprContext)
}

func (s *SelectExprContext) NamesExpr() INamesExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamesExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamesExprContext)
}

func (s *SelectExprContext) FilterExpr() IFilterExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilterExprContext)
}

func (s *SelectExprContext) ScriptExpr() IScriptExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScriptExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	this := p
	_ = this

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
	this := p
	_ = this

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
	this := p
	_ = this

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
	this := p
	_ = this

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
	this := p
	_ = this

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

func (s *FilterExprContext) FilterTrue() antlr.TerminalNode {
	return s.GetToken(JsonpathParserFilterTrue, 0)
}

func (s *FilterExprContext) QueryExpr() IQueryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQueryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQueryExprContext)
}

func (s *FilterExprContext) FilterFalse() antlr.TerminalNode {
	return s.GetToken(JsonpathParserFilterFalse, 0)
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
	this := p
	_ = this

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

	p.SetState(92)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JsonpathParserFilterTrue:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(84)
			p.Match(JsonpathParserFilterTrue)
		}
		{
			p.SetState(85)
			p.QueryExpr()
		}
		{
			p.SetState(86)
			p.Match(JsonpathParserT__5)
		}

	case JsonpathParserFilterFalse:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(88)
			p.Match(JsonpathParserFilterFalse)
		}
		{
			p.SetState(89)
			p.QueryExpr()
		}
		{
			p.SetState(90)
			p.Match(JsonpathParserT__5)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	this := p
	_ = this

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
		p.SetState(94)
		p.Match(JsonpathParserT__6)
	}
	{
		p.SetState(95)
		p.ValueExpr()
	}
	{
		p.SetState(96)
		p.Match(JsonpathParserT__5)
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
	this := p
	_ = this

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

	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(98)
			p.Match(JsonpathParserT__7)
		}
		{
			p.SetState(99)

			var _m = p.Match(JsonpathParserIdentifier)

			localctx.(*QueryExprContext).exists = _m
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(100)
			p.Match(JsonpathParserT__7)
		}
		{
			p.SetState(101)

			var _m = p.Match(JsonpathParserIdentifier)

			localctx.(*QueryExprContext).name = _m
		}
		{
			p.SetState(102)

			var _m = p.Match(JsonpathParserCOMP)

			localctx.(*QueryExprContext).op = _m
		}
		{
			p.SetState(103)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*QueryExprContext).value = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&921088) != 0) {
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
			p.SetState(104)
			p.Match(JsonpathParserT__7)
		}
		{
			p.SetState(105)

			var _m = p.Match(JsonpathParserIdentifier)

			localctx.(*QueryExprContext).name = _m
		}
		{
			p.SetState(106)

			var _m = p.Match(JsonpathParserT__11)

			localctx.(*QueryExprContext).op = _m
		}
		{
			p.SetState(107)

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
	this := p
	_ = this

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

	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(110)
			p.Match(JsonpathParserT__7)
		}
		{
			p.SetState(111)

			var _m = p.Match(JsonpathParserIdentifier)

			localctx.(*ValueExprContext).name = _m
		}
		{
			p.SetState(112)

			var _m = p.Match(JsonpathParserOP)

			localctx.(*ValueExprContext).op = _m
		}
		{
			p.SetState(113)

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
			p.SetState(114)
			p.Match(JsonpathParserT__7)
		}
		{
			p.SetState(115)

			var _m = p.Match(JsonpathParserLength)

			localctx.(*ValueExprContext).name = _m
		}
		{
			p.SetState(116)

			var _m = p.Match(JsonpathParserOP)

			localctx.(*ValueExprContext).op = _m
		}
		{
			p.SetState(117)

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
