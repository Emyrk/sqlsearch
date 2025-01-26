package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Emyrk/sqlsearch/search"
)

func main() {
	log.SetOutput(os.Stderr)
	s := search.New(map[string]string{
		"id":           "workspaces.id",
		"owner_id":     "workspaces.owner_id",
		"template_id":  "workspaces.template_id",
		"name":         "workspaces.name",
		"deleted":      "workspaces.deleted",
		"organization": "workspaces.organization",
	})
	flag.Parse()

	if len(flag.Args()) == 0 {
		log.Fatal("Missing search expression")
		return
	}

	sqlExpr, err := s.Convert(flag.Arg(0))
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Println("Output is SQL to place in WHERE clause:")
	fmt.Println(sqlExpr)
}
