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
		"", "'('", "')'", "'.'", "'\"'", "'''", "", "", "", "", "", "", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "OR", "AND", "NOT", "INTEGER", "DECIMAL", "CMP",
		"SEMICOLON", "STRING", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "A", "N", "D", "O", "R", "T",
		"OR", "AND", "NOT", "INTEGER", "DECIMAL", "CMP", "SEMICOLON", "LOWERCASE",
		"UPPERCASE", "STRING", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 14, 151, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 73, 8, 11, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 81, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 3, 13, 88, 8, 13, 1, 14, 3, 14, 91, 8, 14, 1, 14, 4, 14, 94, 8,
		14, 11, 14, 12, 14, 95, 1, 15, 3, 15, 99, 8, 15, 1, 15, 4, 15, 102, 8,
		15, 11, 15, 12, 15, 103, 1, 15, 1, 15, 4, 15, 108, 8, 15, 11, 15, 12, 15,
		109, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 3, 16, 123, 8, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 20, 1, 20, 4, 20, 133, 8, 20, 11, 20, 12, 20, 134, 1, 20, 1, 20, 1,
		20, 5, 20, 140, 8, 20, 10, 20, 12, 20, 143, 9, 20, 1, 21, 4, 21, 146, 8,
		21, 11, 21, 12, 21, 147, 1, 21, 1, 21, 0, 0, 22, 1, 1, 3, 2, 5, 3, 7, 4,
		9, 5, 11, 0, 13, 0, 15, 0, 17, 0, 19, 0, 21, 0, 23, 6, 25, 7, 27, 8, 29,
		9, 31, 10, 33, 11, 35, 12, 37, 0, 39, 0, 41, 13, 43, 14, 1, 0, 11, 2, 0,
		65, 65, 97, 97, 2, 0, 78, 78, 110, 110, 2, 0, 68, 68, 100, 100, 2, 0, 79,
		79, 111, 111, 2, 0, 82, 82, 114, 114, 2, 0, 84, 84, 116, 116, 1, 0, 48,
		57, 1, 0, 97, 122, 1, 0, 65, 90, 3, 0, 45, 45, 48, 57, 95, 95, 3, 0, 9,
		10, 13, 13, 32, 32, 162, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0,
		0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1,
		0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33,
		1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 1,
		45, 1, 0, 0, 0, 3, 47, 1, 0, 0, 0, 5, 49, 1, 0, 0, 0, 7, 51, 1, 0, 0, 0,
		9, 53, 1, 0, 0, 0, 11, 55, 1, 0, 0, 0, 13, 57, 1, 0, 0, 0, 15, 59, 1, 0,
		0, 0, 17, 61, 1, 0, 0, 0, 19, 63, 1, 0, 0, 0, 21, 65, 1, 0, 0, 0, 23, 72,
		1, 0, 0, 0, 25, 80, 1, 0, 0, 0, 27, 87, 1, 0, 0, 0, 29, 90, 1, 0, 0, 0,
		31, 98, 1, 0, 0, 0, 33, 122, 1, 0, 0, 0, 35, 124, 1, 0, 0, 0, 37, 126,
		1, 0, 0, 0, 39, 128, 1, 0, 0, 0, 41, 132, 1, 0, 0, 0, 43, 145, 1, 0, 0,
		0, 45, 46, 5, 40, 0, 0, 46, 2, 1, 0, 0, 0, 47, 48, 5, 41, 0, 0, 48, 4,
		1, 0, 0, 0, 49, 50, 5, 46, 0, 0, 50, 6, 1, 0, 0, 0, 51, 52, 5, 34, 0, 0,
		52, 8, 1, 0, 0, 0, 53, 54, 5, 39, 0, 0, 54, 10, 1, 0, 0, 0, 55, 56, 7,
		0, 0, 0, 56, 12, 1, 0, 0, 0, 57, 58, 7, 1, 0, 0, 58, 14, 1, 0, 0, 0, 59,
		60, 7, 2, 0, 0, 60, 16, 1, 0, 0, 0, 61, 62, 7, 3, 0, 0, 62, 18, 1, 0, 0,
		0, 63, 64, 7, 4, 0, 0, 64, 20, 1, 0, 0, 0, 65, 66, 7, 5, 0, 0, 66, 22,
		1, 0, 0, 0, 67, 68, 3, 17, 8, 0, 68, 69, 3, 19, 9, 0, 69, 73, 1, 0, 0,
		0, 70, 71, 5, 124, 0, 0, 71, 73, 5, 124, 0, 0, 72, 67, 1, 0, 0, 0, 72,
		70, 1, 0, 0, 0, 73, 24, 1, 0, 0, 0, 74, 75, 3, 11, 5, 0, 75, 76, 3, 13,
		6, 0, 76, 77, 3, 15, 7, 0, 77, 81, 1, 0, 0, 0, 78, 79, 5, 38, 0, 0, 79,
		81, 5, 38, 0, 0, 80, 74, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 81, 26, 1, 0,
		0, 0, 82, 83, 3, 13, 6, 0, 83, 84, 3, 17, 8, 0, 84, 85, 3, 21, 10, 0, 85,
		88, 1, 0, 0, 0, 86, 88, 5, 33, 0, 0, 87, 82, 1, 0, 0, 0, 87, 86, 1, 0,
		0, 0, 88, 28, 1, 0, 0, 0, 89, 91, 5, 45, 0, 0, 90, 89, 1, 0, 0, 0, 90,
		91, 1, 0, 0, 0, 91, 93, 1, 0, 0, 0, 92, 94, 7, 6, 0, 0, 93, 92, 1, 0, 0,
		0, 94, 95, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 30,
		1, 0, 0, 0, 97, 99, 5, 45, 0, 0, 98, 97, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0,
		99, 101, 1, 0, 0, 0, 100, 102, 7, 6, 0, 0, 101, 100, 1, 0, 0, 0, 102, 103,
		1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 105, 1, 0,
		0, 0, 105, 107, 5, 46, 0, 0, 106, 108, 7, 6, 0, 0, 107, 106, 1, 0, 0, 0,
		108, 109, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110,
		32, 1, 0, 0, 0, 111, 123, 5, 61, 0, 0, 112, 113, 5, 61, 0, 0, 113, 123,
		5, 61, 0, 0, 114, 115, 5, 33, 0, 0, 115, 123, 5, 61, 0, 0, 116, 123, 5,
		62, 0, 0, 117, 118, 5, 62, 0, 0, 118, 123, 5, 61, 0, 0, 119, 123, 5, 60,
		0, 0, 120, 121, 5, 60, 0, 0, 121, 123, 5, 61, 0, 0, 122, 111, 1, 0, 0,
		0, 122, 112, 1, 0, 0, 0, 122, 114, 1, 0, 0, 0, 122, 116, 1, 0, 0, 0, 122,
		117, 1, 0, 0, 0, 122, 119, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 123, 34, 1,
		0, 0, 0, 124, 125, 5, 59, 0, 0, 125, 36, 1, 0, 0, 0, 126, 127, 7, 7, 0,
		0, 127, 38, 1, 0, 0, 0, 128, 129, 7, 8, 0, 0, 129, 40, 1, 0, 0, 0, 130,
		133, 3, 37, 18, 0, 131, 133, 3, 39, 19, 0, 132, 130, 1, 0, 0, 0, 132, 131,
		1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0, 134, 135, 1, 0,
		0, 0, 135, 141, 1, 0, 0, 0, 136, 140, 3, 37, 18, 0, 137, 140, 3, 39, 19,
		0, 138, 140, 7, 9, 0, 0, 139, 136, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139,
		138, 1, 0, 0, 0, 140, 143, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141, 142,
		1, 0, 0, 0, 142, 42, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 144, 146, 7, 10,
		0, 0, 145, 144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0,
		147, 148, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 150, 6, 21, 0, 0, 150,
		44, 1, 0, 0, 0, 15, 0, 72, 80, 87, 90, 95, 98, 103, 109, 122, 132, 134,
		139, 141, 147, 1, 6, 0, 0,
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
	searchgrammarLexerT__0      = 1
	searchgrammarLexerT__1      = 2
	searchgrammarLexerT__2      = 3
	searchgrammarLexerT__3      = 4
	searchgrammarLexerT__4      = 5
	searchgrammarLexerOR        = 6
	searchgrammarLexerAND       = 7
	searchgrammarLexerNOT       = 8
	searchgrammarLexerINTEGER   = 9
	searchgrammarLexerDECIMAL   = 10
	searchgrammarLexerCMP       = 11
	searchgrammarLexerSEMICOLON = 12
	searchgrammarLexerSTRING    = 13
	searchgrammarLexerWS        = 14
)
