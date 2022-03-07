package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func regexpression() {
	expression := "(2+(7*8+9))+9"
	logreg := regexp.MustCompile(`\((.*?)\)`)
	logreg = regexp.MustCompile(`\(((?:"\(|\)"|[^()])+)\)`)
	logreg = regexp.MustCompile(`\(([^\(\)\"]*\"[^\"]*\")+[^\(\)\"]*\)|\([^\(\)]*\)`)

	found := logreg.FindString(expression)
	indexRange := logreg.FindStringIndex(expression)

	fmt.Println("Hello", found, indexRange)

	reg, err := regexp.Compile("\\+|\\-|\\*|/|\\^")
	if err != nil {
		log.Fatal(err)
	}
	expression = "log2+6+log8.3/9*8.9-6^8*9"
	processedString := reg.ReplaceAllString(expression, ",")

	operator_list := strings.Split(processedString, ",")
	fmt.Println("Hello", operator_list)

}

// output:

// Hello (7*8+9) [3 10]
// Hello [log2 6 log8.3 9 8.9 6 8 9]