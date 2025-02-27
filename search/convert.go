package search

import (
	"fmt"
	"strings"

	"github.com/Emyrk/sqlsearch/search/parser"
	"github.com/Emyrk/sqlsearch/search/visitor"
	"github.com/antlr4-go/antlr/v4"
	pg_query "github.com/pganalyze/pg_query_go/v6"
)

type Converter struct {
	Refences map[string]string
}

func New(ref map[string]string) *Converter {
	return &Converter{
		Refences: ref,
	}
}

func (c *Converter) Convert(expr string) (string, error) {
	input := antlr.NewInputStream(expr)
	lexer := parser.NewsearchgrammarLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewsearchgrammarParser(stream)
	clause := p.Clause()

	v := visitor.New(c.Refences)
	antlr.ParseTreeWalkerDefault.Walk(v, clause)

	//fmt.Println(antlr.TreesStringTree(clause, nil, nil))

	var n *pg_query.Node
	if v.Stack().Len() == 1 {
		n = v.Stack().Pop()
	} else {
		nodes := make([]*pg_query.Node, 0, v.Stack().Len())
		for !v.Stack().IsEmpty() {
			nodes = append(nodes, v.Stack().Pop())
		}
		n = pg_query.MakeBoolExprNode(pg_query.BoolExprType_OR_EXPR, nodes, 0)
	}

	dsql, err := pg_query.Deparse(&pg_query.ParseResult{
		Version: 0,
		Stmts: []*pg_query.RawStmt{
			{
				Stmt: &pg_query.Node{
					Node: &pg_query.Node_SelectStmt{
						SelectStmt: &pg_query.SelectStmt{
							LimitOption:    pg_query.LimitOption_LIMIT_OPTION_DEFAULT,
							Op:             pg_query.SetOperation_SETOP_NONE,
							DistinctClause: nil,
							IntoClause:     nil,
							TargetList:     nil,
							FromClause:     nil,
							WhereClause:    n,
							GroupClause:    nil,
							GroupDistinct:  false,
							HavingClause:   nil,
							WindowClause:   nil,
							ValuesLists:    nil,
							SortClause:     nil,
							LimitOffset:    nil,
							LimitCount:     nil,
							LockingClause:  nil,
							WithClause:     nil,
							All:            false,
							Larg:           nil,
							Rarg:           nil,
						},
					},
				},
				StmtLocation: 0,
				StmtLen:      0,
			},
		},
	})
	if err != nil {
		return "", fmt.Errorf("deparse: %w", err)
	}

	withoutSelect := strings.TrimPrefix(dsql, "SELECT WHERE")

	// TODO: It would be best to find the args and pass them in as parameters, not strings.
	return strings.TrimSpace(withoutSelect), nil
	//// Attempt #2
	//data, err := proto.Marshal(n)
	//if err != nil {
	//	return "", fmt.Errorf("marshal node: %w", err)
	//}
	//
	//sqlResult, err := pgparser.DeparseFromProtobuf(data)
	//if err != nil {
	//	return "", fmt.Errorf("deparse: %w", err)
	//}
	//
	//return sqlResult, nil
}
