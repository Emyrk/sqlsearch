// Code generated from searchgrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // searchgrammar

import "github.com/antlr4-go/antlr/v4"

// BasesearchgrammarListener is a complete listener for a parse tree produced by searchgrammarParser.
type BasesearchgrammarListener struct{}

var _ searchgrammarListener = &BasesearchgrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasesearchgrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasesearchgrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasesearchgrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasesearchgrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterClause is called when production clause is entered.
func (s *BasesearchgrammarListener) EnterClause(ctx *ClauseContext) {}

// ExitClause is called when production clause is exited.
func (s *BasesearchgrammarListener) ExitClause(ctx *ClauseContext) {}

// EnterNotExpr is called when production notExpr is entered.
func (s *BasesearchgrammarListener) EnterNotExpr(ctx *NotExprContext) {}

// ExitNotExpr is called when production notExpr is exited.
func (s *BasesearchgrammarListener) ExitNotExpr(ctx *NotExprContext) {}

// EnterTermExpr is called when production termExpr is entered.
func (s *BasesearchgrammarListener) EnterTermExpr(ctx *TermExprContext) {}

// ExitTermExpr is called when production termExpr is exited.
func (s *BasesearchgrammarListener) ExitTermExpr(ctx *TermExprContext) {}

// EnterOrExpr is called when production orExpr is entered.
func (s *BasesearchgrammarListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production orExpr is exited.
func (s *BasesearchgrammarListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterParensExpr is called when production parensExpr is entered.
func (s *BasesearchgrammarListener) EnterParensExpr(ctx *ParensExprContext) {}

// ExitParensExpr is called when production parensExpr is exited.
func (s *BasesearchgrammarListener) ExitParensExpr(ctx *ParensExprContext) {}

// EnterAndExpr is called when production andExpr is entered.
func (s *BasesearchgrammarListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production andExpr is exited.
func (s *BasesearchgrammarListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterCmpExpr is called when production cmpExpr is entered.
func (s *BasesearchgrammarListener) EnterCmpExpr(ctx *CmpExprContext) {}

// ExitCmpExpr is called when production cmpExpr is exited.
func (s *BasesearchgrammarListener) ExitCmpExpr(ctx *CmpExprContext) {}

// EnterTerm is called when production term is entered.
func (s *BasesearchgrammarListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BasesearchgrammarListener) ExitTerm(ctx *TermContext) {}

// EnterReference is called when production reference is entered.
func (s *BasesearchgrammarListener) EnterReference(ctx *ReferenceContext) {}

// ExitReference is called when production reference is exited.
func (s *BasesearchgrammarListener) ExitReference(ctx *ReferenceContext) {}

// EnterLiteral_number is called when production literal_number is entered.
func (s *BasesearchgrammarListener) EnterLiteral_number(ctx *Literal_numberContext) {}

// ExitLiteral_number is called when production literal_number is exited.
func (s *BasesearchgrammarListener) ExitLiteral_number(ctx *Literal_numberContext) {}

// EnterLiteral_string is called when production literal_string is entered.
func (s *BasesearchgrammarListener) EnterLiteral_string(ctx *Literal_stringContext) {}

// ExitLiteral_string is called when production literal_string is exited.
func (s *BasesearchgrammarListener) ExitLiteral_string(ctx *Literal_stringContext) {}
