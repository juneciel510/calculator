package main

import (
	"fmt"
	"strings"
)

func parse(expression string) []string {
	processedExpression := []string{}
	var numPart strings.Builder
	var operatorPart strings.Builder

	for _, charRune := range expression {
		char := string(charRune)
		if isNum(char) {
			numPart.WriteString(char)
			if operatorPart.Len() > 0 {
				fmt.Println("operatorPart,operatorPart.Len()", operatorPart.Len())
				processedExpression = append(processedExpression, operatorPart.String())
				operatorPart.Reset()
			}

		} else if isSymbolOp(char) || isLetterOp(char) {
			fmt.Println("***********", char)

			if numPart.Len() > 0 {
				processedExpression = append(processedExpression, numPart.String())
				numPart.Reset()
			}

			if isSymbolOp(char) {
				if operatorPart.Len() > 0 {
					fmt.Println("operatorPart,operatorPart.Len()", operatorPart.Len())
					processedExpression = append(processedExpression, operatorPart.String())
					operatorPart.Reset()
				}
				operatorPart.WriteString(char)
				processedExpression = append(processedExpression, operatorPart.String())
				operatorPart.Reset()
				fmt.Println("1***********", char, processedExpression)

			}

			if isLetterOp(char) {
				operatorPart.WriteString(char)
			}

		} else {
			panic(fmt.Sprintf("error: Invalid character: %s.\n", char))
		}
	}
	if numPart.Len() > 0 {
		processedExpression = append(processedExpression, numPart.String())
	}
	return processedExpression
}
