// Code generated from searchgrammar.g4 by ANTLR 4.12.0. DO NOT EDIT.

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

type searchgrammarLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var searchgrammarlexerLexerStaticData struct {
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

func searchgrammarlexerLexerInit() {
	staticData := &searchgrammarlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'('", "')'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "OR", "AND", "NOT", "INTEGER", "DECIMAL", "CMP", "STRING",
		"WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "A", "N", "D", "O", "R", "T", "OR", "AND", "NOT", "INTEGER",
		"DECIMAL", "CMP", "LOWERCASE", "UPPERCASE", "STRING", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 10, 135, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 3, 8, 59, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3,
		9, 67, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 74, 8, 10, 1, 11,
		3, 11, 77, 8, 11, 1, 11, 4, 11, 80, 8, 11, 11, 11, 12, 11, 81, 1, 12, 3,
		12, 85, 8, 12, 1, 12, 4, 12, 88, 8, 12, 11, 12, 12, 12, 89, 1, 12, 1, 12,
		4, 12, 94, 8, 12, 11, 12, 12, 12, 95, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 109, 8, 13, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 16, 1, 16, 4, 16, 117, 8, 16, 11, 16, 12, 16, 118,
		1, 16, 1, 16, 1, 16, 5, 16, 124, 8, 16, 10, 16, 12, 16, 127, 9, 16, 1,
		17, 4, 17, 130, 8, 17, 11, 17, 12, 17, 131, 1, 17, 1, 17, 0, 0, 18, 1,
		1, 3, 2, 5, 0, 7, 0, 9, 0, 11, 0, 13, 0, 15, 0, 17, 3, 19, 4, 21, 5, 23,
		6, 25, 7, 27, 8, 29, 0, 31, 0, 33, 9, 35, 10, 1, 0, 11, 2, 0, 65, 65, 97,
		97, 2, 0, 78, 78, 110, 110, 2, 0, 68, 68, 100, 100, 2, 0, 79, 79, 111,
		111, 2, 0, 82, 82, 114, 114, 2, 0, 84, 84, 116, 116, 1, 0, 48, 57, 1, 0,
		97, 122, 1, 0, 65, 90, 3, 0, 45, 45, 48, 57, 95, 95, 3, 0, 9, 10, 13, 13,
		32, 32, 146, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0,
		19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0,
		0, 27, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 1, 37, 1, 0, 0,
		0, 3, 39, 1, 0, 0, 0, 5, 41, 1, 0, 0, 0, 7, 43, 1, 0, 0, 0, 9, 45, 1, 0,
		0, 0, 11, 47, 1, 0, 0, 0, 13, 49, 1, 0, 0, 0, 15, 51, 1, 0, 0, 0, 17, 58,
		1, 0, 0, 0, 19, 66, 1, 0, 0, 0, 21, 73, 1, 0, 0, 0, 23, 76, 1, 0, 0, 0,
		25, 84, 1, 0, 0, 0, 27, 108, 1, 0, 0, 0, 29, 110, 1, 0, 0, 0, 31, 112,
		1, 0, 0, 0, 33, 116, 1, 0, 0, 0, 35, 129, 1, 0, 0, 0, 37, 38, 5, 40, 0,
		0, 38, 2, 1, 0, 0, 0, 39, 40, 5, 41, 0, 0, 40, 4, 1, 0, 0, 0, 41, 42, 7,
		0, 0, 0, 42, 6, 1, 0, 0, 0, 43, 44, 7, 1, 0, 0, 44, 8, 1, 0, 0, 0, 45,
		46, 7, 2, 0, 0, 46, 10, 1, 0, 0, 0, 47, 48, 7, 3, 0, 0, 48, 12, 1, 0, 0,
		0, 49, 50, 7, 4, 0, 0, 50, 14, 1, 0, 0, 0, 51, 52, 7, 5, 0, 0, 52, 16,
		1, 0, 0, 0, 53, 54, 3, 11, 5, 0, 54, 55, 3, 13, 6, 0, 55, 59, 1, 0, 0,
		0, 56, 57, 5, 124, 0, 0, 57, 59, 5, 124, 0, 0, 58, 53, 1, 0, 0, 0, 58,
		56, 1, 0, 0, 0, 59, 18, 1, 0, 0, 0, 60, 61, 3, 5, 2, 0, 61, 62, 3, 7, 3,
		0, 62, 63, 3, 9, 4, 0, 63, 67, 1, 0, 0, 0, 64, 65, 5, 38, 0, 0, 65, 67,
		5, 38, 0, 0, 66, 60, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 67, 20, 1, 0, 0, 0,
		68, 69, 3, 7, 3, 0, 69, 70, 3, 11, 5, 0, 70, 71, 3, 15, 7, 0, 71, 74, 1,
		0, 0, 0, 72, 74, 5, 33, 0, 0, 73, 68, 1, 0, 0, 0, 73, 72, 1, 0, 0, 0, 74,
		22, 1, 0, 0, 0, 75, 77, 5, 45, 0, 0, 76, 75, 1, 0, 0, 0, 76, 77, 1, 0,
		0, 0, 77, 79, 1, 0, 0, 0, 78, 80, 7, 6, 0, 0, 79, 78, 1, 0, 0, 0, 80, 81,
		1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 24, 1, 0, 0, 0,
		83, 85, 5, 45, 0, 0, 84, 83, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 87, 1,
		0, 0, 0, 86, 88, 7, 6, 0, 0, 87, 86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89,
		87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 93, 5, 46,
		0, 0, 92, 94, 7, 6, 0, 0, 93, 92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 93,
		1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 26, 1, 0, 0, 0, 97, 109, 5, 61, 0,
		0, 98, 99, 5, 61, 0, 0, 99, 109, 5, 61, 0, 0, 100, 101, 5, 33, 0, 0, 101,
		109, 5, 61, 0, 0, 102, 109, 5, 62, 0, 0, 103, 104, 5, 62, 0, 0, 104, 109,
		5, 61, 0, 0, 105, 109, 5, 60, 0, 0, 106, 107, 5, 60, 0, 0, 107, 109, 5,
		61, 0, 0, 108, 97, 1, 0, 0, 0, 108, 98, 1, 0, 0, 0, 108, 100, 1, 0, 0,
		0, 108, 102, 1, 0, 0, 0, 108, 103, 1, 0, 0, 0, 108, 105, 1, 0, 0, 0, 108,
		106, 1, 0, 0, 0, 109, 28, 1, 0, 0, 0, 110, 111, 7, 7, 0, 0, 111, 30, 1,
		0, 0, 0, 112, 113, 7, 8, 0, 0, 113, 32, 1, 0, 0, 0, 114, 117, 3, 29, 14,
		0, 115, 117, 3, 31, 15, 0, 116, 114, 1, 0, 0, 0, 116, 115, 1, 0, 0, 0,
		117, 118, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119,
		125, 1, 0, 0, 0, 120, 124, 3, 29, 14, 0, 121, 124, 3, 31, 15, 0, 122, 124,
		7, 9, 0, 0, 123, 120, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 123, 122, 1, 0,
		0, 0, 124, 127, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0,
		126, 34, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 128, 130, 7, 10, 0, 0, 129,
		128, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 131, 132,
		1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 134, 6, 17, 0, 0, 134, 36, 1, 0,
		0, 0, 15, 0, 58, 66, 73, 76, 81, 84, 89, 95, 108, 116, 118, 123, 125, 131,
		1, 6, 0, 0,
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

// searchgrammarLexerInit initializes any static state used to implement searchgrammarLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewsearchgrammarLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SearchgrammarLexerInit() {
	staticData := &searchgrammarlexerLexerStaticData
	staticData.once.Do(searchgrammarlexerLexerInit)
}

// NewsearchgrammarLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewsearchgrammarLexer(input antlr.CharStream) *searchgrammarLexer {
	SearchgrammarLexerInit()
	l := new(searchgrammarLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &searchgrammarlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "searchgrammar.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// searchgrammarLexer tokens.
const (
	searchgrammarLexerT__0    = 1
	searchgrammarLexerT__1    = 2
	searchgrammarLexerOR      = 3
	searchgrammarLexerAND     = 4
	searchgrammarLexerNOT     = 5
	searchgrammarLexerINTEGER = 6
	searchgrammarLexerDECIMAL = 7
	searchgrammarLexerCMP     = 8
	searchgrammarLexerSTRING  = 9
	searchgrammarLexerWS      = 10
)
