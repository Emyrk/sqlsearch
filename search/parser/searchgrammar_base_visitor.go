// Code generated from searchgrammar.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // searchgrammar

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BasesearchgrammarVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasesearchgrammarVisitor) VisitClause(ctx *ClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesearchgrammarVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesearchgrammarVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}
