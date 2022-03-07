package main

import (
	"errors"
	"fmt"
)

func findInnermostParenthesesIndices(expressionProcessed []string) (int, int, error) {
	var leftParenthesesIndex int
	var rightParenthesesIndex int
	for index, item := range expressionProcessed {
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

func removeSingleParentheses(expressionProcessed []string, leftParenthesesIndex int, rightParenthesesIndex int) []string {
	// copyOperators:= make([]string, len(operators))
	// copy(copyOperators, operators)

	expressionInParentheses := expressionProcessed[leftParenthesesIndex+1 : rightParenthesesIndex]
	fmt.Println("-------removeSingleParentheses------")
	fmt.Println("expressionInParentheses:", expressionInParentheses)
	result, err := noParanthesesCalc(expressionInParentheses)
	if err != nil {
		panic(err)
	}

	expressionNew := append(expressionProcessed[:leftParenthesesIndex], asString(result))
	expressionNew = append(expressionNew, expressionProcessed[rightParenthesesIndex+1:]...)
	return expressionNew
}

func removeAllParentheses(expressionProcessed []string) []string {
	expressionNew := expressionProcessed
	for {
		leftParenthesesIndex, rightParenthesesIndex, err := findInnermostParenthesesIndices(expressionNew)
		if err != nil {
			break
		}
		fmt.Println("==========")
		fmt.Println("expressionNew, leftParenthesesIndex, rightParenthesesIndex:", expressionNew, leftParenthesesIndex, rightParenthesesIndex)
		expressionNew = removeSingleParentheses(expressionNew, leftParenthesesIndex, rightParenthesesIndex)
	}

	return expressionNew
}