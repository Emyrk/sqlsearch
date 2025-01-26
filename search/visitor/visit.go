package visitor

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Emyrk/sqlsearch/search/parser"
	"github.com/antlr4-go/antlr/v4"
	pg_query "github.com/pganalyze/pg_query_go/v6"
)

var _ = pg_query.MakeAConstStrNode

var _ antlr.ParseTreeListener = &SearchVisitor{}

type SearchVisitor struct {
	parser.BasesearchgrammarListener

	stack *Stack[*pg_query.Node]
	refs  map[string]*pg_query.Node
}

func (s *SearchVisitor) VisitTerminal(node antlr.TerminalNode) {}

func (s *SearchVisitor) VisitErrorNode(node antlr.ErrorNode) {}

func (s *SearchVisitor) EnterEveryRule(ctx antlr.ParserRuleContext) {}

func (s *SearchVisitor) ExitEveryRule(ctx antlr.ParserRuleContext) {}
func New(refs map[string]string) *SearchVisitor {
	cref := make(map[string]*pg_query.Node, len(refs))
	for k, v := range refs {
		parts := strings.Split(v, ".")
		fields := make([]*pg_query.Node, 0, len(parts))
		for _, p := range parts {
			fields = append(fields, pg_query.MakeStrNode(p))
		}
		cref[k] = pg_query.MakeColumnRefNode(fields, 0)
	}

	return &SearchVisitor{
		stack: NewStack[*pg_query.Node](),
		refs:  cref,
	}
}

func (s *SearchVisitor) Stack() *Stack[*pg_query.Node] {
	return s.stack
}

// EnterClause is called when production clause is entered.
func (s *SearchVisitor) EnterClause(ctx *parser.ClauseContext) {
}

// ExitClause is called when production clause is exited.
func (s *SearchVisitor) ExitClause(ctx *parser.ClauseContext) {
	nodes := make([]*pg_query.Node, 0, s.Stack().Len())
	for !s.Stack().IsEmpty() {
		nodes = append(nodes, s.Stack().Pop())
	}
	n := pg_query.MakeBoolExprNode(pg_query.BoolExprType_OR_EXPR, nodes, 0)
	s.stack.Push(nodes[0])
	var _ = n
}

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
	ref := ctx.GetChild(1).(antlr.ParseTree)
	text := ref.GetText()
	node := pg_query.MakeAConstStrNode(text, 0)
	s.stack.Push(node)
}

// ExitReference is called when production reference is exited.
func (s *SearchVisitor) ExitReference(ctx *parser.ReferenceContext) {
	ref := ctx.GetText()
	columnRef, ok := s.refs[ref]
	if !ok {
		panic(fmt.Sprintf("%q reference not found", ref))
	}

	s.stack.Push(columnRef)
}
