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
	fmt.Printf("%s\n", elapsed)
}

// Reuse AST

// func main() {
// 	argsWithoutProg := os.Args[1:]
// 	content, _ := ioutil.ReadFile(argsWithoutProg[0])
// 	parser := astparser.NewParser()

// 	doc := *ast.NewDocument()
// 	doc.Input.ResetInputBytes(content)
// 	report := operationreport.Report{}

// 	parser.Parse(&doc, &report)

// 	start := time.Now()

// 	report = operationreport.Report{}
// 	doc.Input.ResetInputBytes(content)
// 	parser.Parse(&doc, &report)
// 	if report.HasErrors() {
// 		log.Fatal(report.Error())
// 	}

// 	elapsed := time.Since(start)
// 	fmt.Printf("%s\n", elapsed)
// }
