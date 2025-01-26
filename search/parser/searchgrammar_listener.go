// Code generated from searchgrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // searchgrammar

import "github.com/antlr4-go/antlr/v4"

// searchgrammarListener is a complete listener for a parse tree produced by searchgrammarParser.
type searchgrammarListener interface {
	antlr.ParseTreeListener

	// EnterClause is called when entering the clause production.
	EnterClause(c *ClauseContext)

	// EnterNotExpr is called when entering the notExpr production.
	EnterNotExpr(c *NotExprContext)

	// EnterTermExpr is called when entering the termExpr production.
	EnterTermExpr(c *TermExprContext)

	// EnterOrExpr is called when entering the orExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterParensExpr is called when entering the parensExpr production.
	EnterParensExpr(c *ParensExprContext)

	// EnterAndExpr is called when entering the andExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterCmpExpr is called when entering the cmpExpr production.
	EnterCmpExpr(c *CmpExprContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterReference is called when entering the reference production.
	EnterReference(c *ReferenceContext)

	// EnterLiteral_number is called when entering the literal_number production.
	EnterLiteral_number(c *Literal_numberContext)

	// EnterLiteral_string is called when entering the literal_string production.
	EnterLiteral_string(c *Literal_stringContext)

	// ExitClause is called when exiting the clause production.
	ExitClause(c *ClauseContext)

	// ExitNotExpr is called when exiting the notExpr production.
	ExitNotExpr(c *NotExprContext)

	// ExitTermExpr is called when exiting the termExpr production.
	ExitTermExpr(c *TermExprContext)

	// ExitOrExpr is called when exiting the orExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitParensExpr is called when exiting the parensExpr production.
	ExitParensExpr(c *ParensExprContext)

	// ExitAndExpr is called when exiting the andExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitCmpExpr is called when exiting the cmpExpr production.
	ExitCmpExpr(c *CmpExprContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitReference is called when exiting the reference production.
	ExitReference(c *ReferenceContext)

	// ExitLiteral_number is called when exiting the literal_number production.
	ExitLiteral_number(c *Literal_numberContext)

	// ExitLiteral_string is called when exiting the literal_string production.
	ExitLiteral_string(c *Literal_stringContext)
}
