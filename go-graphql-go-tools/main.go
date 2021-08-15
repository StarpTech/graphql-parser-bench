package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/jensneuse/graphql-go-tools/pkg/astparser"
)

func main() {
	argsWithoutProg := os.Args[1:]
	content, _ := ioutil.ReadFile(argsWithoutProg[0])
	start := time.Now()
	_, defReport := astparser.ParseGraphqlDocumentBytes(content)
	if defReport.HasErrors() {
		log.Fatal(defReport.Error())
	}
	elapsed := time.Since(start)
	fmt.Printf("graphql-go-tools: %s\n", elapsed)
}
