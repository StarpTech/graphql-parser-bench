package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/vektah/gqlparser/ast"
	"github.com/vektah/gqlparser/parser"
)

func main() {
	argsWithoutProg := os.Args[1:]
	content, _ := ioutil.ReadFile(argsWithoutProg[0])
	start := time.Now()
	if argsWithoutProg[1] == "schema" {
		_, err := parser.ParseSchema(&ast.Source{Input: string(content)})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		_, err := parser.ParseQuery(&ast.Source{Input: string(content)})
		if err != nil {
			log.Fatal(err)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("%s\n", elapsed)
}
