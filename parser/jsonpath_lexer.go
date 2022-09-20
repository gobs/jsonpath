// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type JsonpathLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var jsonpathlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jsonpathlexerLexerInit() {
	staticData := &jsonpathlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
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
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "OP", "COMP", "Identifier", "Length", "INT",
		"NUMBER", "QUOTED", "REGEX", "DOTS", "STAR", "FilterTrue", "FilterFalse",
		"WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 25, 177, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1,
		12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 3, 13, 100, 8, 13, 1, 14, 1, 14, 5, 14, 104, 8, 14, 10, 14, 12,
		14, 107, 9, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 16, 3, 16, 119, 8, 16, 1, 16, 4, 16, 122, 8, 16, 11, 16, 12,
		16, 123, 1, 17, 1, 17, 1, 17, 4, 17, 129, 8, 17, 11, 17, 12, 17, 130, 3,
		17, 133, 8, 17, 1, 18, 1, 18, 5, 18, 137, 8, 18, 10, 18, 12, 18, 140, 9,
		18, 1, 18, 1, 18, 1, 19, 1, 19, 5, 19, 146, 8, 19, 10, 19, 12, 19, 149,
		9, 19, 1, 19, 1, 19, 5, 19, 153, 8, 19, 10, 19, 12, 19, 156, 9, 19, 1,
		20, 1, 20, 1, 20, 3, 20, 161, 8, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22,
		1, 23, 1, 23, 1, 23, 1, 24, 4, 24, 172, 8, 24, 11, 24, 12, 24, 173, 1,
		24, 1, 24, 0, 0, 25, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8,
		17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 1, 0, 7,
		3, 0, 42, 43, 45, 45, 47, 47, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48,
		57, 65, 90, 95, 95, 97, 122, 1, 0, 39, 39, 1, 0, 47, 47, 4, 0, 83, 83,
		105, 105, 109, 109, 117, 117, 3, 0, 9, 10, 13, 13, 32, 32, 191, 0, 1, 1,
		0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1,
		0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17,
		1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0,
		25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0,
		0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0,
		0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0,
		0, 0, 0, 49, 1, 0, 0, 0, 1, 51, 1, 0, 0, 0, 3, 53, 1, 0, 0, 0, 5, 55, 1,
		0, 0, 0, 7, 57, 1, 0, 0, 0, 9, 59, 1, 0, 0, 0, 11, 61, 1, 0, 0, 0, 13,
		63, 1, 0, 0, 0, 15, 65, 1, 0, 0, 0, 17, 68, 1, 0, 0, 0, 19, 73, 1, 0, 0,
		0, 21, 79, 1, 0, 0, 0, 23, 84, 1, 0, 0, 0, 25, 87, 1, 0, 0, 0, 27, 99,
		1, 0, 0, 0, 29, 101, 1, 0, 0, 0, 31, 108, 1, 0, 0, 0, 33, 118, 1, 0, 0,
		0, 35, 125, 1, 0, 0, 0, 37, 134, 1, 0, 0, 0, 39, 143, 1, 0, 0, 0, 41, 160,
		1, 0, 0, 0, 43, 162, 1, 0, 0, 0, 45, 164, 1, 0, 0, 0, 47, 167, 1, 0, 0,
		0, 49, 171, 1, 0, 0, 0, 51, 52, 5, 36, 0, 0, 52, 2, 1, 0, 0, 0, 53, 54,
		5, 91, 0, 0, 54, 4, 1, 0, 0, 0, 55, 56, 5, 93, 0, 0, 56, 6, 1, 0, 0, 0,
		57, 58, 5, 58, 0, 0, 58, 8, 1, 0, 0, 0, 59, 60, 5, 44, 0, 0, 60, 10, 1,
		0, 0, 0, 61, 62, 5, 41, 0, 0, 62, 12, 1, 0, 0, 0, 63, 64, 5, 40, 0, 0,
		64, 14, 1, 0, 0, 0, 65, 66, 5, 64, 0, 0, 66, 67, 5, 46, 0, 0, 67, 16, 1,
		0, 0, 0, 68, 69, 5, 116, 0, 0, 69, 70, 5, 114, 0, 0, 70, 71, 5, 117, 0,
		0, 71, 72, 5, 101, 0, 0, 72, 18, 1, 0, 0, 0, 73, 74, 5, 102, 0, 0, 74,
		75, 5, 97, 0, 0, 75, 76, 5, 108, 0, 0, 76, 77, 5, 115, 0, 0, 77, 78, 5,
		101, 0, 0, 78, 20, 1, 0, 0, 0, 79, 80, 5, 110, 0, 0, 80, 81, 5, 117, 0,
		0, 81, 82, 5, 108, 0, 0, 82, 83, 5, 108, 0, 0, 83, 22, 1, 0, 0, 0, 84,
		85, 5, 61, 0, 0, 85, 86, 5, 126, 0, 0, 86, 24, 1, 0, 0, 0, 87, 88, 7, 0,
		0, 0, 88, 26, 1, 0, 0, 0, 89, 100, 5, 62, 0, 0, 90, 91, 5, 62, 0, 0, 91,
		100, 5, 61, 0, 0, 92, 100, 5, 60, 0, 0, 93, 94, 5, 60, 0, 0, 94, 100, 5,
		61, 0, 0, 95, 96, 5, 61, 0, 0, 96, 100, 5, 61, 0, 0, 97, 98, 5, 33, 0,
		0, 98, 100, 5, 61, 0, 0, 99, 89, 1, 0, 0, 0, 99, 90, 1, 0, 0, 0, 99, 92,
		1, 0, 0, 0, 99, 93, 1, 0, 0, 0, 99, 95, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0,
		100, 28, 1, 0, 0, 0, 101, 105, 7, 1, 0, 0, 102, 104, 7, 2, 0, 0, 103, 102,
		1, 0, 0, 0, 104, 107, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 105, 106, 1, 0,
		0, 0, 106, 30, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 108, 109, 5, 108, 0, 0,
		109, 110, 5, 101, 0, 0, 110, 111, 5, 110, 0, 0, 111, 112, 5, 103, 0, 0,
		112, 113, 5, 116, 0, 0, 113, 114, 5, 104, 0, 0, 114, 115, 5, 40, 0, 0,
		115, 116, 5, 41, 0, 0, 116, 32, 1, 0, 0, 0, 117, 119, 5, 45, 0, 0, 118,
		117, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 121, 1, 0, 0, 0, 120, 122,
		2, 48, 57, 0, 121, 120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 121, 1,
		0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 34, 1, 0, 0, 0, 125, 132, 3, 33, 16,
		0, 126, 128, 5, 46, 0, 0, 127, 129, 2, 48, 57, 0, 128, 127, 1, 0, 0, 0,
		129, 130, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131,
		133, 1, 0, 0, 0, 132, 126, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 36, 1,
		0, 0, 0, 134, 138, 5, 39, 0, 0, 135, 137, 8, 3, 0, 0, 136, 135, 1, 0, 0,
		0, 137, 140, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139,
		141, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 141, 142, 5, 39, 0, 0, 142, 38,
		1, 0, 0, 0, 143, 147, 5, 47, 0, 0, 144, 146, 8, 4, 0, 0, 145, 144, 1, 0,
		0, 0, 146, 149, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0,
		148, 150, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 150, 154, 5, 47, 0, 0, 151,
		153, 7, 5, 0, 0, 152, 151, 1, 0, 0, 0, 153, 156, 1, 0, 0, 0, 154, 152,
		1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 40, 1, 0, 0, 0, 156, 154, 1, 0,
		0, 0, 157, 161, 5, 46, 0, 0, 158, 159, 5, 46, 0, 0, 159, 161, 5, 46, 0,
		0, 160, 157, 1, 0, 0, 0, 160, 158, 1, 0, 0, 0, 161, 42, 1, 0, 0, 0, 162,
		163, 5, 42, 0, 0, 163, 44, 1, 0, 0, 0, 164, 165, 5, 63, 0, 0, 165, 166,
		5, 40, 0, 0, 166, 46, 1, 0, 0, 0, 167, 168, 5, 33, 0, 0, 168, 169, 5, 40,
		0, 0, 169, 48, 1, 0, 0, 0, 170, 172, 7, 6, 0, 0, 171, 170, 1, 0, 0, 0,
		172, 173, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174,
		175, 1, 0, 0, 0, 175, 176, 6, 24, 0, 0, 176, 50, 1, 0, 0, 0, 12, 0, 99,
		105, 118, 123, 130, 132, 138, 147, 154, 160, 173, 1, 6, 0, 0,
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

// JsonpathLexerInit initializes any static state used to implement JsonpathLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewJsonpathLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func JsonpathLexerInit() {
	staticData := &jsonpathlexerLexerStaticData
	staticData.once.Do(jsonpathlexerLexerInit)
}

// NewJsonpathLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewJsonpathLexer(input antlr.CharStream) *JsonpathLexer {
	JsonpathLexerInit()
	l := new(JsonpathLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &jsonpathlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Jsonpath.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// JsonpathLexer tokens.
const (
	JsonpathLexerT__0        = 1
	JsonpathLexerT__1        = 2
	JsonpathLexerT__2        = 3
	JsonpathLexerT__3        = 4
	JsonpathLexerT__4        = 5
	JsonpathLexerT__5        = 6
	JsonpathLexerT__6        = 7
	JsonpathLexerT__7        = 8
	JsonpathLexerT__8        = 9
	JsonpathLexerT__9        = 10
	JsonpathLexerT__10       = 11
	JsonpathLexerT__11       = 12
	JsonpathLexerOP          = 13
	JsonpathLexerCOMP        = 14
	JsonpathLexerIdentifier  = 15
	JsonpathLexerLength      = 16
	JsonpathLexerINT         = 17
	JsonpathLexerNUMBER      = 18
	JsonpathLexerQUOTED      = 19
	JsonpathLexerREGEX       = 20
	JsonpathLexerDOTS        = 21
	JsonpathLexerSTAR        = 22
	JsonpathLexerFilterTrue  = 23
	JsonpathLexerFilterFalse = 24
	JsonpathLexerWS          = 25
)
