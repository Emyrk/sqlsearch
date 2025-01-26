// Code generated from searchgrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // searchgrammar

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by searchgrammarParser.
type searchgrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by searchgrammarParser#clause.
	VisitClause(ctx *ClauseContext) interface{}

	// Visit a parse tree produced by searchgrammarParser#notExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by searchgrammarParser#termExpr.
	VisitTermExpr(ctx *TermExprContext) interface{}

	// Visit a parse tree produced by searchgrammarParser#orExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by searchgrammarParser#parensExpr.
	VisitParensExpr(ctx *ParensExprContext) interface{}

	// Visit a parse tree produced by searchgrammarParser#andExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by searchgrammarParser#cmpExpr.
	VisitCmpExpr(ctx *CmpExprContext) interface{}

	// Visit a parse tree produced by searchgrammarParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by searchgrammarParser#reference.
	VisitReference(ctx *ReferenceContext) interface{}

	// Visit a parse tree produced by searchgrammarParser#literal_number.
	VisitLiteral_number(ctx *Literal_numberContext) interface{}

	// Visit a parse tree produced by searchgrammarParser#literal_string.
	VisitLiteral_string(ctx *Literal_stringContext) interface{}
}
