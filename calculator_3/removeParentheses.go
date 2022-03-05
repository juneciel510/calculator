package main

import (
	"errors"
)

func findFirstParenthesesIndices(numbers []float64, operators []string) (int, int, error) {
	var leftParenthesesIndex int
	var rightParenthesesIndex int
	for index, operator := range operators {
		if operator == "(" {
			leftParenthesesIndex = index
		}
		if operator == ")" {
			rightParenthesesIndex = index
			return leftParenthesesIndex, rightParenthesesIndex, nil
		}
	}
	return leftParenthesesIndex, rightParenthesesIndex, errors.New("Parentheses Not find")
}

func removeSingleParentheses(numbers []float64, operators []string, leftParenthesesIndex int, rightParenthesesIndex int) ([]float64, []string) {
	copyNumbers:= make([]float64, len(numbers))
	copy(copyNumbers, numbers)
	copyOperators:= make([]string, len(operators))
	copy(copyOperators, operators)

	operatorsInParentheses := copyOperators[leftParenthesesIndex+1 : rightParenthesesIndex]
	numbersInParentheses := copyNumbers[leftParenthesesIndex:rightParenthesesIndex]
	
	result := noParenthesesCalculation(numbersInParentheses,operatorsInParentheses)

	operatorsNew := append(copyOperators[:leftParenthesesIndex], copyOperators[rightParenthesesIndex+1:]...)
	
	numbersNew := append(copyNumbers[:leftParenthesesIndex], result)
	numbersNew = append(numbersNew,copyNumbers[rightParenthesesIndex:]...)
	return numbersNew, operatorsNew
}

func removeAllParentheses(numbers []float64, operators []string) ([]float64, []string) {
	operatorsNew := operators
	numbersNew := numbers
	for  {
		leftParenthesesIndex, rightParenthesesIndex, err := findFirstParenthesesIndices(numbersNew,operatorsNew)
		if err != nil {
			break	
		}	
		numbersNew,operatorsNew  = removeSingleParentheses(numbersNew, operatorsNew, leftParenthesesIndex, rightParenthesesIndex)		
	}

	return numbersNew, operatorsNew
}

