package search_test

import (
	"testing"

	"github.com/Emyrk/sqlsearch/search"
	"github.com/stretchr/testify/require"
)

func TestConvert(t *testing.T) {
	tc := []struct {
		input  string
		expSql string
	}{
		{
			input:  "42 < 10",
			expSql: "42 < 10",
		},
	}

	for _, c := range tc {
		t.Run(c.input, func(t *testing.T) {
			sql, err := search.Convert(c.input)
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
