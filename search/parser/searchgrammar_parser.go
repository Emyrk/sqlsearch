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
		"", "'('", "')'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "OR", "AND", "NOT", "INTEGER", "DECIMAL", "CMP", "STRING",
		"WS",
	}
	staticData.ruleNames = []string{
		"clause", "expr", "term",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 10, 38, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 4, 0, 8, 8, 0,
		11, 0, 12, 0, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1,
		20, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 31,
		8, 1, 10, 1, 12, 1, 34, 9, 1, 1, 2, 1, 2, 1, 2, 0, 1, 2, 3, 0, 2, 4, 0,
		1, 2, 0, 6, 7, 9, 9, 40, 0, 7, 1, 0, 0, 0, 2, 19, 1, 0, 0, 0, 4, 35, 1,
		0, 0, 0, 6, 8, 3, 2, 1, 0, 7, 6, 1, 0, 0, 0, 8, 9, 1, 0, 0, 0, 9, 7, 1,
		0, 0, 0, 9, 10, 1, 0, 0, 0, 10, 1, 1, 0, 0, 0, 11, 12, 6, 1, -1, 0, 12,
		20, 3, 4, 2, 0, 13, 14, 5, 1, 0, 0, 14, 15, 3, 2, 1, 0, 15, 16, 5, 2, 0,
		0, 16, 20, 1, 0, 0, 0, 17, 18, 5, 5, 0, 0, 18, 20, 3, 2, 1, 1, 19, 11,
		1, 0, 0, 0, 19, 13, 1, 0, 0, 0, 19, 17, 1, 0, 0, 0, 20, 32, 1, 0, 0, 0,
		21, 22, 10, 4, 0, 0, 22, 23, 5, 3, 0, 0, 23, 31, 3, 2, 1, 5, 24, 25, 10,
		3, 0, 0, 25, 26, 5, 4, 0, 0, 26, 31, 3, 2, 1, 4, 27, 28, 10, 2, 0, 0, 28,
		29, 5, 8, 0, 0, 29, 31, 3, 2, 1, 3, 30, 21, 1, 0, 0, 0, 30, 24, 1, 0, 0,
		0, 30, 27, 1, 0, 0, 0, 31, 34, 1, 0, 0, 0, 32, 30, 1, 0, 0, 0, 32, 33,
		1, 0, 0, 0, 33, 3, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 35, 36, 7, 0, 0, 0,
		36, 5, 1, 0, 0, 0, 4, 9, 19, 30, 32,
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
	searchgrammarParserEOF     = antlr.TokenEOF
	searchgrammarParserT__0    = 1
	searchgrammarParserT__1    = 2
	searchgrammarParserOR      = 3
	searchgrammarParserAND     = 4
	searchgrammarParserNOT     = 5
	searchgrammarParserINTEGER = 6
	searchgrammarParserDECIMAL = 7
	searchgrammarParserCMP     = 8
	searchgrammarParserSTRING  = 9
	searchgrammarParserWS      = 10
)

// searchgrammarParser rules.
const (
	searchgrammarParserRULE_clause = 0
	searchgrammarParserRULE_expr   = 1
	searchgrammarParserRULE_term   = 2
)

// IClauseContext is an interface to support dynamic dispatch.
type IClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsClauseContext differentiates from other interfaces.
	IsClauseContext()
}

type ClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClauseContext() *ClauseContext {
	var p = new(ClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = searchgrammarParserRULE_clause
	return p
}

func (*ClauseContext) IsClauseContext() {}

func NewClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClauseContext {
	var p = new(ClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = searchgrammarParserRULE_clause

	return p
}

func (s *ClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ClauseContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ClauseContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case searchgrammarVisitor:
		return t.VisitClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *searchgrammarParser) Clause() (localctx IClauseContext) {
	this := p
	_ = this

	localctx = NewClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, searchgrammarParserRULE_clause)
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
	p.SetState(7)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&738) != 0) {
		{
			p.SetState(6)
			p.expr(0)
		}

		p.SetState(9)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Term() ITermContext
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	NOT() antlr.TerminalNode
	OR() antlr.TerminalNode
	AND() antlr.TerminalNode
	CMP() antlr.TerminalNode

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

func (s *ExprContext) Term() ITermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(searchgrammarParserNOT, 0)
}

func (s *ExprContext) OR() antlr.TerminalNode {
	return s.GetToken(searchgrammarParserOR, 0)
}

func (s *ExprContext) AND() antlr.TerminalNode {
	return s.GetToken(searchgrammarParserAND, 0)
}

func (s *ExprContext) CMP() antlr.TerminalNode {
	return s.GetToken(searchgrammarParserCMP, 0)
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
	return p.expr(0)
}

func (p *searchgrammarParser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, searchgrammarParserRULE_expr, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(19)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case searchgrammarParserINTEGER, searchgrammarParserDECIMAL, searchgrammarParserSTRING:
		{
			p.SetState(12)
			p.Term()
		}

	case searchgrammarParserT__0:
		{
			p.SetState(13)
			p.Match(searchgrammarParserT__0)
		}
		{
			p.SetState(14)
			p.expr(0)
		}
		{
			p.SetState(15)
			p.Match(searchgrammarParserT__1)
		}

	case searchgrammarParserNOT:
		{
			p.SetState(17)
			p.Match(searchgrammarParserNOT)
		}
		{
			p.SetState(18)
			p.expr(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(30)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, searchgrammarParserRULE_expr)
				p.SetState(21)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(22)
					p.Match(searchgrammarParserOR)
				}
				{
					p.SetState(23)
					p.expr(5)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, searchgrammarParserRULE_expr)
				p.SetState(24)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(25)
					p.Match(searchgrammarParserAND)
				}
				{
					p.SetState(26)
					p.expr(4)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, searchgrammarParserRULE_expr)
				p.SetState(27)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(28)
					p.Match(searchgrammarParserCMP)
				}
				{
					p.SetState(29)
					p.expr(3)
				}

			}

		}
		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTEGER() antlr.TerminalNode
	DECIMAL() antlr.TerminalNode
	STRING() antlr.TerminalNode

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

func (s *TermContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(searchgrammarParserINTEGER, 0)
}

func (s *TermContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(searchgrammarParserDECIMAL, 0)
}

func (s *TermContext) STRING() antlr.TerminalNode {
	return s.GetToken(searchgrammarParserSTRING, 0)
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
	p.EnterRule(localctx, 4, searchgrammarParserRULE_term)
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
		p.SetState(35)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&704) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *searchgrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *searchgrammarParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
