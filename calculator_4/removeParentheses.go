package main

import (
	"errors"
)

func findInnermostParenthesesIndices(expression []string) (int, int, error) {
	var leftParenthesesIndex int
	var rightParenthesesIndex int
	for index, item := range expression {
		if item == "(" {
			leftParenthesesIndex = index
		}
		if item == ")" {
			rightParenthesesIndex = index
			return leftParenthesesIndex, rightParenthesesIndex, nil
		}
	}
	return leftParenthesesIndex, rightParenthesesIndex, errors.New("Parentheses Not find")
}

func removeSingleParentheses(expression []string, leftParenthesesIndex int, rightParenthesesIndex int) []string {
	expressionInParentheses := expression[leftParenthesesIndex+1 : rightParenthesesIndex]
	result, err := noParanthesesCalc(expressionInParentheses)
	if err != nil {
		panic(err)
	}

	expressionNew := append(expression[:leftParenthesesIndex], asString(result))
	expressionNew = append(expressionNew, expression[rightParenthesesIndex+1:]...)
	return expressionNew
}

func removeAllParentheses(expression []string) []string {
	expressionNew := expression
	for {
		leftParenthesesIndex, rightParenthesesIndex, err := findInnermostParenthesesIndices(expressionNew)
		if err != nil {
			break
		}
	
		expressionNew = removeSingleParentheses(expressionNew, leftParenthesesIndex, rightParenthesesIndex)
	}

	return expressionNew
}