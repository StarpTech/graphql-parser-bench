package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/graphql-go/graphql/language/parser"
	"github.com/graphql-go/graphql/language/source"
)

func main() {
	start := time.Now()
	content, _ := ioutil.ReadFile("./../schema.graphql")
	_, err := parser.Parse(parser.ParseParams{
		Source: &source.Source{
			Body: content,
			Name: "GraphQL",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(start)
	fmt.Printf("graphql: %s\n", elapsed)
}
