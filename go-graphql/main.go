package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/graphql-go/graphql/language/parser"
	"github.com/graphql-go/graphql/language/source"
)

func main() {
	argsWithoutProg := os.Args[1:]
	content, _ := ioutil.ReadFile(argsWithoutProg[0])
	start := time.Now()
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
	fmt.Printf("%s\n", elapsed)
}
