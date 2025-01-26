package visitor

import (
	"strconv"

	"github.com/Emyrk/sqlsearch/search/parser"
	"github.com/antlr4-go/antlr/v4"
	pg_query "github.com/pganalyze/pg_query_go/v6"
)

var _ = pg_query.MakeAConstStrNode

type SearchVisitor struct {
	parser.BasesearchgrammarListener

	stack *Stack[*pg_query.Node]
}

func New() *SearchVisitor {
	return &SearchVisitor{
		stack: NewStack[*pg_query.Node](),
	}
}

func (s *SearchVisitor) Stack() *Stack[*pg_query.Node] {
	return s.stack
}

//func (s *SearchVisitor) VisitTerminal(node antlr.TerminalNode) {}
//
//// VisitErrorNode is called when an error node is visited.
//func (s *SearchVisitor) VisitErrorNode(node antlr.ErrorNode) {}

//// EnterEveryRule is called when any rule is entered.
//func (s *BasesearchgrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}
//
//// ExitEveryRule is called when any rule is exited.
//func (s *BasesearchgrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterClause is called when production clause is entered.
func (s *SearchVisitor) EnterClause(ctx *parser.ClauseContext) {
}

// ExitClause is called when production clause is exited.
func (s *SearchVisitor) ExitClause(ctx *parser.ClauseContext) {}

// EnterCmpExpr is called when production cmpExpr is entered.
func (s *SearchVisitor) EnterCmpExpr(ctx *parser.CmpExprContext) {}

// ExitCmpExpr is called when production cmpExpr is exited.
func (s *SearchVisitor) ExitCmpExpr(ctx *parser.CmpExprContext) {
	c := ctx.GetChild(1).(antlr.ParseTree)
	l, r := s.stack.Pop(), s.stack.Pop()

	n := pg_query.MakeAExprNode(pg_query.A_Expr_Kind_AEXPR_OP,
		[]*pg_query.Node{pg_query.MakeStrNode(c.GetText())},
		l, r, 0,
	)
	s.stack.Push(n)
}

// EnterTerm is called when production term is entered.
func (s *SearchVisitor) EnterTerm(ctx *parser.TermContext) {}

// ExitTerm is called when production term is exited.
func (s *SearchVisitor) ExitTerm(ctx *parser.TermContext) {}

// EnterLiteral_number is called when production literal_number is entered.
func (s *SearchVisitor) EnterLiteral_number(ctx *parser.Literal_numberContext) {}

// ExitLiteral_number is called when production literal_number is exited.
func (s *SearchVisitor) ExitLiteral_number(ctx *parser.Literal_numberContext) {
	num, err := strconv.ParseInt(ctx.GetText(), 10, 64)
	if err != nil {
		panic(err)
	}
	node := pg_query.MakeAConstIntNode(num, 0)
	s.stack.Push(node)
}

// EnterLiteral_string is called when production literal_string is entered.
func (s *SearchVisitor) EnterLiteral_string(ctx *parser.Literal_stringContext) {}

// ExitLiteral_string is called when production literal_string is exited.
func (s *SearchVisitor) ExitLiteral_string(ctx *parser.Literal_stringContext) {
	node := pg_query.MakeAConstStrNode(ctx.GetText(), 0)
	s.stack.Push(node)
}
