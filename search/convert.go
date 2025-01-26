package search

import (
	"fmt"
	"go/parser"


	"github.com/antlr4-go/antlr/v4"
)


// "42 < 10 && word < \"done\""
func convert(expr string) error {
	tree, err := parser.ParseExpr(expr)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	antlr.NewParseTreeWalker().Walk()

	return err
}
