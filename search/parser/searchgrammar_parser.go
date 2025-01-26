// Code generated from searchgrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // searchgrammar

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type searchgrammarParser struct {
	*antlr.BaseParser
}

var SearchgrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func searchgrammarParserInit() {
	staticData := &SearchgrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "'('", "')'", "'\"'", "'''",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "OR", "AND", "NOT", "INTEGER", "DECIMAL", "CMP",
		"STRING", "WS",
	}
	staticData.RuleNames = []string{
		"clause", "expr", "term", "reference", "literal_number", "literal_string",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 12, 59, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 1, 0, 4, 0, 14, 8, 0, 11, 0, 12, 0, 15, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 26, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 37, 8, 1, 10, 1, 12, 1, 40, 9, 1, 1, 2,
		1, 2, 1, 2, 3, 2, 45, 8, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 3, 5, 57, 8, 5, 1, 5, 0, 1, 2, 6, 0, 2, 4, 6, 8, 10, 0,
		1, 1, 0, 8, 9, 61, 0, 13, 1, 0, 0, 0, 2, 25, 1, 0, 0, 0, 4, 44, 1, 0, 0,
		0, 6, 46, 1, 0, 0, 0, 8, 48, 1, 0, 0, 0, 10, 56, 1, 0, 0, 0, 12, 14, 3,
		2, 1, 0, 13, 12, 1, 0, 0, 0, 14, 15, 1, 0, 0, 0, 15, 13, 1, 0, 0, 0, 15,
		16, 1, 0, 0, 0, 16, 1, 1, 0, 0, 0, 17, 18, 6, 1, -1, 0, 18, 26, 3, 4, 2,
		0, 19, 20, 5, 1, 0, 0, 20, 21, 3, 2, 1, 0, 21, 22, 5, 2, 0, 0, 22, 26,
		1, 0, 0, 0, 23, 24, 5, 7, 0, 0, 24, 26, 3, 2, 1, 1, 25, 17, 1, 0, 0, 0,
		25, 19, 1, 0, 0, 0, 25, 23, 1, 0, 0, 0, 26, 38, 1, 0, 0, 0, 27, 28, 10,
		4, 0, 0, 28, 29, 5, 5, 0, 0, 29, 37, 3, 2, 1, 5, 30, 31, 10, 3, 0, 0, 31,
		32, 5, 6, 0, 0, 32, 37, 3, 2, 1, 4, 33, 34, 10, 2, 0, 0, 34, 35, 5, 10,
		0, 0, 35, 37, 3, 2, 1, 3, 36, 27, 1, 0, 0, 0, 36, 30, 1, 0, 0, 0, 36, 33,
		1, 0, 0, 0, 37, 40, 1, 0, 0, 0, 38, 36, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0,
		39, 3, 1, 0, 0, 0, 40, 38, 1, 0, 0, 0, 41, 45, 3, 8, 4, 0, 42, 45, 3, 6,
		3, 0, 43, 45, 3, 10, 5, 0, 44, 41, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 44,
		43, 1, 0, 0, 0, 45, 5, 1, 0, 0, 0, 46, 47, 5, 11, 0, 0, 47, 7, 1, 0, 0,
		0, 48, 49, 7, 0, 0, 0, 49, 9, 1, 0, 0, 0, 50, 51, 5, 3, 0, 0, 51, 52, 5,
		11, 0, 0, 52, 57, 5, 3, 0, 0, 53, 54, 5, 4, 0, 0, 54, 55, 5, 11, 0, 0,
		55, 57, 5, 4, 0, 0, 56, 50, 1, 0, 0, 0, 56, 53, 1, 0, 0, 0, 57, 11, 1,
		0, 0, 0, 6, 15, 25, 36, 38, 44, 56,
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
	staticData := &SearchgrammarParserStaticData
	staticData.once.Do(searchgrammarParserInit)
}

// NewsearchgrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewsearchgrammarParser(input antlr.TokenStream) *searchgrammarParser {
	SearchgrammarParserInit()
	this := new(searchgrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SearchgrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "searchgrammar.g4"

	return this
}

// searchgrammarParser tokens.
const (
	searchgrammarParserEOF     = antlr.TokenEOF
	searchgrammarParserT__0    = 1
	searchgrammarParserT__1    = 2
	searchgrammarParserT__2    = 3
	searchgrammarParserT__3    = 4
	searchgrammarParserOR      = 5
	searchgrammarParserAND     = 6
	searchgrammarParserNOT     = 7
	searchgrammarParserINTEGER = 8
	searchgrammarParserDECIMAL = 9
	searchgrammarParserCMP     = 10
	searchgrammarParserSTRING  = 11
	searchgrammarParserWS      = 12
)

// searchgrammarParser rules.
const (
	searchgrammarParserRULE_clause         = 0
	searchgrammarParserRULE_expr           = 1
	searchgrammarParserRULE_term           = 2
	searchgrammarParserRULE_reference      = 3
	searchgrammarParserRULE_literal_number = 4
	searchgrammarParserRULE_literal_string = 5
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClauseContext() *ClauseContext {
	var p = new(ClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = searchgrammarParserRULE_clause
	return p
}

func InitEmptyClauseContext(p *ClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = searchgrammarParserRULE_clause
}

func (*ClauseContext) IsClauseContext() {}

func NewClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClauseContext {
	var p = new(ClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

func (s *ClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(searchgrammarListener); ok {
		listenerT.EnterClause(s)
	}
}

func (s *ClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(searchgrammarListener); ok {
		listenerT.ExitClause(s)
	}
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
	localctx = NewClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, searchgrammarParserRULE_clause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(13)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2970) != 0) {
		{
			p.SetState(12)
			p.expr(0)
		}

		p.SetState(15)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = searchgrammarParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = searchgrammarParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = searchgrammarParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NotExprContext struct {
	ExprContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(searchgrammarParserNOT, 0)
}

func (s *NotExprContext) Expr() IExprContext {
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

func (s *NotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(searchgrammarListener); ok {
		listenerT.EnterNotExpr(s)
	}
}

func (s *NotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(searchgrammarListener); ok {
		listenerT.ExitNotExpr(s)
	}
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case searchgrammarVisitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type TermExprContext struct {
	ExprContext
}

func NewTermExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TermExprContext {
	var p = new(TermExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *TermExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermExprContext) Term() ITermContext {
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

func (s *TermExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(searchgrammarListener); ok {
		listenerT.EnterTermExpr(s)
	}
}

func (s *TermExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(searchgrammarListener); ok {
		listenerT.ExitTermExpr(s)
	}
}

func (s *TermExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case searchgrammarVisitor:
		return t.VisitTermExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrExprContext struct {
	ExprContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) AllExpr() []IExprContext {
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

func (s *OrExprContext) Expr(i int) IExprContext {
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

func (s *OrExprContext) OR() antlr.TerminalNode {
	return s.GetToken(searchgrammarParserOR, 0)
}

func (s *OrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(searchgrammarListener); ok {
		listenerT.EnterOrExpr(s)
	}
}

func (s *OrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(searchgrammarListener); ok {
		listenerT.ExitOrExpr(s)
	}
}

func (s *OrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case searchgrammarVisitor:
		return t.VisitOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParensExprContext struct {
	ExprContext
}

func NewParensExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensExprContext {
	var p = new(ParensExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParensExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensExprContext) Expr() IExprContext {
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

func (s *ParensExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(searchgrammarListener); ok {
		listenerT.EnterParensExpr(s)
	}
}

func (s *ParensExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(searchgrammarListener); ok {
		listenerT.ExitParensExpr(s)
	}
}

func (s *ParensExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case searchgrammarVisitor:
		return t.VisitParensExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndExprContext struct {
	ExprContext
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AllExpr() []IExprContext {
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

func (s *AndExprContext) Expr(i int) IExprContext {
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

func (s *AndExprContext) AND() antlr.TerminalNode {
	return s.GetToken(searchgrammarParserAND, 0)
}

func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(searchgrammarListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(searchgrammarListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

func (s *AndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case searchgrammarVisitor:
		return t.VisitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CmpExprContext struct {
	ExprContext
}

func NewCmpExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CmpExprContext {
	var p = new(CmpExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CmpExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmpExprContext) AllExpr() []IExprContext {
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

func (s *CmpExprContext) Expr(i int) IExprContext {
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

func (s *CmpExprContext) CMP() antlr.TerminalNode {
	return s.GetToken(searchgrammarParserCMP, 0)
}

func (s *CmpExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(searchgrammarListener); ok {
		listenerT.EnterCmpExpr(s)
	}
}

func (s *CmpExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(searchgrammarListener); ok {
		listenerT.ExitCmpExpr(s)
	}
}

func (s *CmpExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case searchgrammarVisitor:
		return t.VisitCmpExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *searchgrammarParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *searchgrammarParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, searchgrammarParserRULE_expr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case searchgrammarParserT__2, searchgrammarParserT__3, searchgrammarParserINTEGER, searchgrammarParserDECIMAL, searchgrammarParserSTRING:
		localctx = NewTermExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(18)
			p.Term()
		}

	case searchgrammarParserT__0:
		localctx = NewParensExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(19)
			p.Match(searchgrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(20)
			p.expr(0)
		}
		{
			p.SetState(21)
			p.Match(searchgrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case searchgrammarParserNOT:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(23)
			p.Match(searchgrammarParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(24)
			p.expr(1)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(36)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, searchgrammarParserRULE_expr)
				p.SetState(27)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(28)
					p.Match(searchgrammarParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(29)
					p.expr(5)
				}

			case 2:
				localctx = NewAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, searchgrammarParserRULE_expr)
				p.SetState(30)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(31)
					p.Match(searchgrammarParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(32)
					p.expr(4)
				}

			case 3:
				localctx = NewCmpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, searchgrammarParserRULE_expr)
				p.SetState(33)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(34)
					p.Match(searchgrammarParserCMP)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(35)
					p.expr(3)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Literal_number() ILiteral_numberContext
	Reference() IReferenceContext
	Literal_string() ILiteral_stringContext

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = searchgrammarParserRULE_term
	return p
}

func InitEmptyTermContext(p *TermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = searchgrammarParserRULE_term
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = searchgrammarParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) Literal_number() ILiteral_numberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteral_numberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteral_numberContext)
}

func (s *TermContext) Reference() IReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReferenceContext)
}

func (s *TermContext) Literal_string() ILiteral_stringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteral_stringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteral_stringContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(searchgrammarListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(searchgrammarListener); ok {
		listenerT.ExitTerm(s)
	}
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
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, searchgrammarParserRULE_term)
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case searchgrammarParserINTEGER, searchgrammarParserDECIMAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)
			p.Literal_number()
		}

	case searchgrammarParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(42)
			p.Reference()
		}

	case searchgrammarParserT__2, searchgrammarParserT__3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(43)
			p.Literal_string()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReferenceContext is an interface to support dynamic dispatch.
type IReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsReferenceContext differentiates from other interfaces.
	IsReferenceContext()
}

type ReferenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReferenceContext() *ReferenceContext {
	var p = new(ReferenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = searchgrammarParserRULE_reference
	return p
}

func InitEmptyReferenceContext(p *ReferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = searchgrammarParserRULE_reference
}

func (*ReferenceContext) IsReferenceContext() {}

func NewReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferenceContext {
	var p = new(ReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = searchgrammarParserRULE_reference

	return p
}

func (s *ReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ReferenceContext) STRING() antlr.TerminalNode {
	return s.GetToken(searchgrammarParserSTRING, 0)
}

func (s *ReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(searchgrammarListener); ok {
		listenerT.EnterReference(s)
	}
}

func (s *ReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(searchgrammarListener); ok {
		listenerT.ExitReference(s)
	}
}

func (s *ReferenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case searchgrammarVisitor:
		return t.VisitReference(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *searchgrammarParser) Reference() (localctx IReferenceContext) {
	localctx = NewReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, searchgrammarParserRULE_reference)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Match(searchgrammarParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteral_numberContext is an interface to support dynamic dispatch.
type ILiteral_numberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTEGER() antlr.TerminalNode
	DECIMAL() antlr.TerminalNode

	// IsLiteral_numberContext differentiates from other interfaces.
	IsLiteral_numberContext()
}

type Literal_numberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteral_numberContext() *Literal_numberContext {
	var p = new(Literal_numberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = searchgrammarParserRULE_literal_number
	return p
}

func InitEmptyLiteral_numberContext(p *Literal_numberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = searchgrammarParserRULE_literal_number
}

func (*Literal_numberContext) IsLiteral_numberContext() {}

func NewLiteral_numberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Literal_numberContext {
	var p = new(Literal_numberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = searchgrammarParserRULE_literal_number

	return p
}

func (s *Literal_numberContext) GetParser() antlr.Parser { return s.parser }

func (s *Literal_numberContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(searchgrammarParserINTEGER, 0)
}

func (s *Literal_numberContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(searchgrammarParserDECIMAL, 0)
}

func (s *Literal_numberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Literal_numberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Literal_numberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(searchgrammarListener); ok {
		listenerT.EnterLiteral_number(s)
	}
}

func (s *Literal_numberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(searchgrammarListener); ok {
		listenerT.ExitLiteral_number(s)
	}
}

func (s *Literal_numberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case searchgrammarVisitor:
		return t.VisitLiteral_number(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *searchgrammarParser) Literal_number() (localctx ILiteral_numberContext) {
	localctx = NewLiteral_numberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, searchgrammarParserRULE_literal_number)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		_la = p.GetTokenStream().LA(1)

		if !(_la == searchgrammarParserINTEGER || _la == searchgrammarParserDECIMAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteral_stringContext is an interface to support dynamic dispatch.
type ILiteral_stringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsLiteral_stringContext differentiates from other interfaces.
	IsLiteral_stringContext()
}

type Literal_stringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteral_stringContext() *Literal_stringContext {
	var p = new(Literal_stringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = searchgrammarParserRULE_literal_string
	return p
}

func InitEmptyLiteral_stringContext(p *Literal_stringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = searchgrammarParserRULE_literal_string
}

func (*Literal_stringContext) IsLiteral_stringContext() {}

func NewLiteral_stringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Literal_stringContext {
	var p = new(Literal_stringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = searchgrammarParserRULE_literal_string

	return p
}

func (s *Literal_stringContext) GetParser() antlr.Parser { return s.parser }

func (s *Literal_stringContext) STRING() antlr.TerminalNode {
	return s.GetToken(searchgrammarParserSTRING, 0)
}

func (s *Literal_stringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Literal_stringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Literal_stringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(searchgrammarListener); ok {
		listenerT.EnterLiteral_string(s)
	}
}

func (s *Literal_stringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(searchgrammarListener); ok {
		listenerT.ExitLiteral_string(s)
	}
}

func (s *Literal_stringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case searchgrammarVisitor:
		return t.VisitLiteral_string(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *searchgrammarParser) Literal_string() (localctx ILiteral_stringContext) {
	localctx = NewLiteral_stringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, searchgrammarParserRULE_literal_string)
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case searchgrammarParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(50)
			p.Match(searchgrammarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(51)
			p.Match(searchgrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(52)
			p.Match(searchgrammarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case searchgrammarParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(53)
			p.Match(searchgrammarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(54)
			p.Match(searchgrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(55)
			p.Match(searchgrammarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
