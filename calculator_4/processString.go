package main

import (
	"fmt"
	"strings"
)

//seperate operators and numbers of the input string
//e.g log(3.8*8)+3/8
//["log", "(", "3.8","*","8",")" ,"+", "3","/","8"]
func parse(expressionInput string) []string {
	expression := []string{}
	var numPart strings.Builder
	var operatorPart strings.Builder

	for _, charRune := range expressionInput {
		char := string(charRune)
		//the start of number means the end of operator
		if isNum(char) {
			numPart.WriteString(char)
			if operatorPart.Len() > 0 {
				expression = append(expression, operatorPart.String())
				operatorPart.Reset()
			}

		//the start of operator means the end of another operator(e.g 'log(' , '*(')) or a number
		} else if isSymbolOp(char) || isLetterOp(char) {
			if numPart.Len() > 0 {
				expression = append(expression, numPart.String())
				numPart.Reset()
			}

			if isSymbolOp(char) {
				if operatorPart.Len() > 0 {
					expression = append(expression, operatorPart.String())
					operatorPart.Reset()
				}
				operatorPart.WriteString(char)
				expression = append(expression, operatorPart.String())
				operatorPart.Reset()

			}

			if isLetterOp(char) {
				operatorPart.WriteString(char)
			}

		} else {
			fmt.Println(fmt.Sprintf("error: Invalid character: %s.\n", char))
		}
	}
	if numPart.Len() > 0 {
		expression = append(expression, numPart.String())
	}
	return expression
}
