// Code generated from searchgrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // searchgrammar

import "github.com/antlr4-go/antlr/v4"

type BasesearchgrammarVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasesearchgrammarVisitor) VisitClause(ctx *ClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesearchgrammarVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesearchgrammarVisitor) VisitTermExpr(ctx *TermExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesearchgrammarVisitor) VisitOrExpr(ctx *OrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesearchgrammarVisitor) VisitParensExpr(ctx *ParensExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesearchgrammarVisitor) VisitAndExpr(ctx *AndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesearchgrammarVisitor) VisitCmpExpr(ctx *CmpExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesearchgrammarVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesearchgrammarVisitor) VisitReference(ctx *ReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesearchgrammarVisitor) VisitLiteral_number(ctx *Literal_numberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasesearchgrammarVisitor) VisitLiteral_string(ctx *Literal_stringContext) interface{} {
	return v.VisitChildren(ctx)
}
