package search_test

import (
	"fmt"
	"testing"

	"github.com/Emyrk/sqlsearch/search"
	pg_query "github.com/pganalyze/pg_query_go/v6"
	"github.com/stretchr/testify/require"
)

func TestConvert(t *testing.T) {
	tc := []struct {
		input  string
		expSql string
		refs   map[string]string
	}{
		{
			input:  "42 < 10",
			expSql: "42 < 10",
		},
		{
			input:  "color = 'red'",
			expSql: "users.favorite_color = 'red'",
			refs: map[string]string{
				"color": "users.favorite_color",
			},
		},
	}

	for _, c := range tc {
		t.Run(c.input, func(t *testing.T) {
			sch := search.New(c.refs)
			sql, err := sch.Convert(c.input)
			require.NoError(t, err)
			require.Equal(t, c.expSql, sql)
		})
	}

	//sql, err := search.Convert("42 < 10 && word < \"done\"")
	//require.NoError(t, err)
	//
	//tree, err := pg_query.Parse("SELECT * WHERE 42 < 10")
	//require.NoError(t, err)
	//
	//out, err := pg_query.Deparse(tree)
	//require.NoError(t, err)
	//
	//fmt.Println(out)
	//
	//fmt.Println(sql)
	//fmt.Println(tree.String())
}

func TestEx(t *testing.T) {
	tree, err := pg_query.Parse("SELECT * WHERE workspaces.count < 10")
	require.NoError(t, err)

	fmt.Println(tree.String())
}
