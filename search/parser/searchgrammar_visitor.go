// Code generated from searchgrammar.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // searchgrammar

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by searchgrammarParser.
type searchgrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by searchgrammarParser#clause.
	VisitClause(ctx *ClauseContext) interface{}

	// Visit a parse tree produced by searchgrammarParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by searchgrammarParser#term.
	VisitTerm(ctx *TermContext) interface{}
}
