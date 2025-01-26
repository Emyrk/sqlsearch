// Code generated from searchgrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
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

var SearchgrammarLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func searchgrammarlexerLexerInit() {
	staticData := &SearchgrammarLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'('", "')'", "'\"'", "'''",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "OR", "AND", "NOT", "INTEGER", "DECIMAL", "CMP",
		"STRING", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "A", "N", "D", "O", "R", "T", "OR",
		"AND", "NOT", "INTEGER", "DECIMAL", "CMP", "LOWERCASE", "UPPERCASE",
		"STRING", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 12, 143, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		3, 10, 67, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 75,
		8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 82, 8, 12, 1, 13, 3, 13,
		85, 8, 13, 1, 13, 4, 13, 88, 8, 13, 11, 13, 12, 13, 89, 1, 14, 3, 14, 93,
		8, 14, 1, 14, 4, 14, 96, 8, 14, 11, 14, 12, 14, 97, 1, 14, 1, 14, 4, 14,
		102, 8, 14, 11, 14, 12, 14, 103, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 117, 8, 15, 1, 16, 1, 16,
		1, 17, 1, 17, 1, 18, 1, 18, 4, 18, 125, 8, 18, 11, 18, 12, 18, 126, 1,
		18, 1, 18, 1, 18, 5, 18, 132, 8, 18, 10, 18, 12, 18, 135, 9, 18, 1, 19,
		4, 19, 138, 8, 19, 11, 19, 12, 19, 139, 1, 19, 1, 19, 0, 0, 20, 1, 1, 3,
		2, 5, 3, 7, 4, 9, 0, 11, 0, 13, 0, 15, 0, 17, 0, 19, 0, 21, 5, 23, 6, 25,
		7, 27, 8, 29, 9, 31, 10, 33, 0, 35, 0, 37, 11, 39, 12, 1, 0, 11, 2, 0,
		65, 65, 97, 97, 2, 0, 78, 78, 110, 110, 2, 0, 68, 68, 100, 100, 2, 0, 79,
		79, 111, 111, 2, 0, 82, 82, 114, 114, 2, 0, 84, 84, 116, 116, 1, 0, 48,
		57, 1, 0, 97, 122, 1, 0, 65, 90, 3, 0, 45, 45, 48, 57, 95, 95, 3, 0, 9,
		10, 13, 13, 32, 32, 154, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0,
		0, 0, 0, 7, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1,
		0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 37,
		1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 1, 41, 1, 0, 0, 0, 3, 43, 1, 0, 0, 0, 5,
		45, 1, 0, 0, 0, 7, 47, 1, 0, 0, 0, 9, 49, 1, 0, 0, 0, 11, 51, 1, 0, 0,
		0, 13, 53, 1, 0, 0, 0, 15, 55, 1, 0, 0, 0, 17, 57, 1, 0, 0, 0, 19, 59,
		1, 0, 0, 0, 21, 66, 1, 0, 0, 0, 23, 74, 1, 0, 0, 0, 25, 81, 1, 0, 0, 0,
		27, 84, 1, 0, 0, 0, 29, 92, 1, 0, 0, 0, 31, 116, 1, 0, 0, 0, 33, 118, 1,
		0, 0, 0, 35, 120, 1, 0, 0, 0, 37, 124, 1, 0, 0, 0, 39, 137, 1, 0, 0, 0,
		41, 42, 5, 40, 0, 0, 42, 2, 1, 0, 0, 0, 43, 44, 5, 41, 0, 0, 44, 4, 1,
		0, 0, 0, 45, 46, 5, 34, 0, 0, 46, 6, 1, 0, 0, 0, 47, 48, 5, 39, 0, 0, 48,
		8, 1, 0, 0, 0, 49, 50, 7, 0, 0, 0, 50, 10, 1, 0, 0, 0, 51, 52, 7, 1, 0,
		0, 52, 12, 1, 0, 0, 0, 53, 54, 7, 2, 0, 0, 54, 14, 1, 0, 0, 0, 55, 56,
		7, 3, 0, 0, 56, 16, 1, 0, 0, 0, 57, 58, 7, 4, 0, 0, 58, 18, 1, 0, 0, 0,
		59, 60, 7, 5, 0, 0, 60, 20, 1, 0, 0, 0, 61, 62, 3, 15, 7, 0, 62, 63, 3,
		17, 8, 0, 63, 67, 1, 0, 0, 0, 64, 65, 5, 124, 0, 0, 65, 67, 5, 124, 0,
		0, 66, 61, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 67, 22, 1, 0, 0, 0, 68, 69,
		3, 9, 4, 0, 69, 70, 3, 11, 5, 0, 70, 71, 3, 13, 6, 0, 71, 75, 1, 0, 0,
		0, 72, 73, 5, 38, 0, 0, 73, 75, 5, 38, 0, 0, 74, 68, 1, 0, 0, 0, 74, 72,
		1, 0, 0, 0, 75, 24, 1, 0, 0, 0, 76, 77, 3, 11, 5, 0, 77, 78, 3, 15, 7,
		0, 78, 79, 3, 19, 9, 0, 79, 82, 1, 0, 0, 0, 80, 82, 5, 33, 0, 0, 81, 76,
		1, 0, 0, 0, 81, 80, 1, 0, 0, 0, 82, 26, 1, 0, 0, 0, 83, 85, 5, 45, 0, 0,
		84, 83, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 87, 1, 0, 0, 0, 86, 88, 7,
		6, 0, 0, 87, 86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89,
		90, 1, 0, 0, 0, 90, 28, 1, 0, 0, 0, 91, 93, 5, 45, 0, 0, 92, 91, 1, 0,
		0, 0, 92, 93, 1, 0, 0, 0, 93, 95, 1, 0, 0, 0, 94, 96, 7, 6, 0, 0, 95, 94,
		1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0,
		98, 99, 1, 0, 0, 0, 99, 101, 5, 46, 0, 0, 100, 102, 7, 6, 0, 0, 101, 100,
		1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103, 104, 1, 0,
		0, 0, 104, 30, 1, 0, 0, 0, 105, 117, 5, 61, 0, 0, 106, 107, 5, 61, 0, 0,
		107, 117, 5, 61, 0, 0, 108, 109, 5, 33, 0, 0, 109, 117, 5, 61, 0, 0, 110,
		117, 5, 62, 0, 0, 111, 112, 5, 62, 0, 0, 112, 117, 5, 61, 0, 0, 113, 117,
		5, 60, 0, 0, 114, 115, 5, 60, 0, 0, 115, 117, 5, 61, 0, 0, 116, 105, 1,
		0, 0, 0, 116, 106, 1, 0, 0, 0, 116, 108, 1, 0, 0, 0, 116, 110, 1, 0, 0,
		0, 116, 111, 1, 0, 0, 0, 116, 113, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 117,
		32, 1, 0, 0, 0, 118, 119, 7, 7, 0, 0, 119, 34, 1, 0, 0, 0, 120, 121, 7,
		8, 0, 0, 121, 36, 1, 0, 0, 0, 122, 125, 3, 33, 16, 0, 123, 125, 3, 35,
		17, 0, 124, 122, 1, 0, 0, 0, 124, 123, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0,
		126, 124, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 133, 1, 0, 0, 0, 128,
		132, 3, 33, 16, 0, 129, 132, 3, 35, 17, 0, 130, 132, 7, 9, 0, 0, 131, 128,
		1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 131, 130, 1, 0, 0, 0, 132, 135, 1, 0,
		0, 0, 133, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 38, 1, 0, 0, 0,
		135, 133, 1, 0, 0, 0, 136, 138, 7, 10, 0, 0, 137, 136, 1, 0, 0, 0, 138,
		139, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 141,
		1, 0, 0, 0, 141, 142, 6, 19, 0, 0, 142, 40, 1, 0, 0, 0, 15, 0, 66, 74,
		81, 84, 89, 92, 97, 103, 116, 124, 126, 131, 133, 139, 1, 6, 0, 0,
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
	staticData := &SearchgrammarLexerLexerStaticData
	staticData.once.Do(searchgrammarlexerLexerInit)
}

// NewsearchgrammarLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewsearchgrammarLexer(input antlr.CharStream) *searchgrammarLexer {
	SearchgrammarLexerInit()
	l := new(searchgrammarLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &SearchgrammarLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "searchgrammar.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// searchgrammarLexer tokens.
const (
	searchgrammarLexerT__0    = 1
	searchgrammarLexerT__1    = 2
	searchgrammarLexerT__2    = 3
	searchgrammarLexerT__3    = 4
	searchgrammarLexerOR      = 5
	searchgrammarLexerAND     = 6
	searchgrammarLexerNOT     = 7
	searchgrammarLexerINTEGER = 8
	searchgrammarLexerDECIMAL = 9
	searchgrammarLexerCMP     = 10
	searchgrammarLexerSTRING  = 11
	searchgrammarLexerWS      = 12
)
