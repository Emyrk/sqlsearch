// Code generated from searchgrammar.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // searchgrammar

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

type searchgrammarParser struct {
	*antlr.BaseParser
}

var searchgrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func searchgrammarParserInit() {
	staticData := &searchgrammarParserStaticData
	staticData.literalNames = []string{
		"", "'+'", "'-'", "'*'", "'/'", "'('", "')'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "NUMBER", "WS",
	}
	staticData.ruleNames = []string{
		"expr", "term", "factor",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 8, 32, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 1, 0, 1, 0, 5, 0,
		10, 8, 0, 10, 0, 12, 0, 13, 9, 0, 1, 1, 1, 1, 1, 1, 5, 1, 18, 8, 1, 10,
		1, 12, 1, 21, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 30,
		8, 2, 1, 2, 0, 0, 3, 0, 2, 4, 0, 2, 1, 0, 1, 2, 1, 0, 3, 4, 32, 0, 6, 1,
		0, 0, 0, 2, 14, 1, 0, 0, 0, 4, 29, 1, 0, 0, 0, 6, 11, 3, 2, 1, 0, 7, 8,
		7, 0, 0, 0, 8, 10, 3, 2, 1, 0, 9, 7, 1, 0, 0, 0, 10, 13, 1, 0, 0, 0, 11,
		9, 1, 0, 0, 0, 11, 12, 1, 0, 0, 0, 12, 1, 1, 0, 0, 0, 13, 11, 1, 0, 0,
		0, 14, 19, 3, 4, 2, 0, 15, 16, 7, 1, 0, 0, 16, 18, 3, 4, 2, 0, 17, 15,
		1, 0, 0, 0, 18, 21, 1, 0, 0, 0, 19, 17, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0,
		20, 3, 1, 0, 0, 0, 21, 19, 1, 0, 0, 0, 22, 30, 5, 7, 0, 0, 23, 24, 5, 5,
		0, 0, 24, 25, 3, 0, 0, 0, 25, 26, 5, 6, 0, 0, 26, 30, 1, 0, 0, 0, 27, 28,
		7, 0, 0, 0, 28, 30, 3, 4, 2, 0, 29, 22, 1, 0, 0, 0, 29, 23, 1, 0, 0, 0,
		29, 27, 1, 0, 0, 0, 30, 5, 1, 0, 0, 0, 3, 11, 19, 29,
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

// searchgrammarParserInit initializes any static state used to implement searchgrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewsearchgrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SearchgrammarParserInit() {
	staticData := &searchgrammarParserStaticData
	staticData.once.Do(searchgrammarParserInit)
}

// NewsearchgrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewsearchgrammarParser(input antlr.TokenStream) *searchgrammarParser {
	SearchgrammarParserInit()
	this := new(searchgrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &searchgrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "searchgrammar.g4"

	return this
}

// searchgrammarParser tokens.
const (
	searchgrammarParserEOF    = antlr.TokenEOF
	searchgrammarParserT__0   = 1
	searchgrammarParserT__1   = 2
	searchgrammarParserT__2   = 3
	searchgrammarParserT__3   = 4
	searchgrammarParserT__4   = 5
	searchgrammarParserT__5   = 6
	searchgrammarParserNUMBER = 7
	searchgrammarParserWS     = 8
)

// searchgrammarParser rules.
const (
	searchgrammarParserRULE_expr   = 0
	searchgrammarParserRULE_term   = 1
	searchgrammarParserRULE_factor = 2
)

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTerm() []ITermContext
	Term(i int) ITermContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = searchgrammarParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = searchgrammarParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
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

	return t.(ITermContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case searchgrammarVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *searchgrammarParser) Expr() (localctx IExprContext) {
	this := p
	_ = this

	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, searchgrammarParserRULE_expr)
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
		p.SetState(6)
		p.Term()
	}
	p.SetState(11)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == searchgrammarParserT__0 || _la == searchgrammarParserT__1 {
		{
			p.SetState(7)
			_la = p.GetTokenStream().LA(1)

			if !(_la == searchgrammarParserT__0 || _la == searchgrammarParserT__1) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(8)
			p.Term()
		}

		p.SetState(13)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFactor() []IFactorContext
	Factor(i int) IFactorContext

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = searchgrammarParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = searchgrammarParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) AllFactor() []IFactorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFactorContext); ok {
			len++
		}
	}

	tst := make([]IFactorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFactorContext); ok {
			tst[i] = t.(IFactorContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) Factor(i int) IFactorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
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

	return t.(IFactorContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case searchgrammarVisitor:
		return t.VisitTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *searchgrammarParser) Term() (localctx ITermContext) {
	this := p
	_ = this

	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, searchgrammarParserRULE_term)
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
		p.SetState(14)
		p.Factor()
	}
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == searchgrammarParserT__2 || _la == searchgrammarParserT__3 {
		{
			p.SetState(15)
			_la = p.GetTokenStream().LA(1)

			if !(_la == searchgrammarParserT__2 || _la == searchgrammarParserT__3) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(16)
			p.Factor()
		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	Expr() IExprContext
	Factor() IFactorContext

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = searchgrammarParserRULE_factor
	return p
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = searchgrammarParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(searchgrammarParserNUMBER, 0)
}

func (s *FactorContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FactorContext) Factor() IFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case searchgrammarVisitor:
		return t.VisitFactor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *searchgrammarParser) Factor() (localctx IFactorContext) {
	this := p
	_ = this

	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, searchgrammarParserRULE_factor)
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

	p.SetState(29)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case searchgrammarParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(22)
			p.Match(searchgrammarParserNUMBER)
		}

	case searchgrammarParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(23)
			p.Match(searchgrammarParserT__4)
		}
		{
			p.SetState(24)
			p.Expr()
		}
		{
			p.SetState(25)
			p.Match(searchgrammarParserT__5)
		}

	case searchgrammarParserT__0, searchgrammarParserT__1:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(27)
			_la = p.GetTokenStream().LA(1)

			if !(_la == searchgrammarParserT__0 || _la == searchgrammarParserT__1) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(28)
			p.Factor()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
