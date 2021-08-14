package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/jensneuse/graphql-go-tools/pkg/astparser"
)

func main() {
	start := time.Now()
	content, _ := ioutil.ReadFile("./../schema.graphql")
	_, defReport := astparser.ParseGraphqlDocumentBytes(content)
	if defReport.HasErrors() {
		log.Fatal(defReport.Error())
	}
	elapsed := time.Since(start)
	fmt.Printf("graphql-go-tools: %s\n", elapsed)
}
