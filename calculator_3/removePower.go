package main

import (
	"errors"
	"math"
)

func findFirstPowerIndices(numbers []float64, operators []string) (int, error) {
	var index int
	for index, operator := range operators {
		if operator == "^" {
			return index, nil
		}
	}
	return index, errors.New("Power Not find")
}

func removeSinglePower(numbers []float64, operators []string, powerIndex int) ([]float64, []string) {
	copyNumbers:= make([]float64, len(numbers))
	copy(copyNumbers, numbers)
	copyOperators:= make([]string, len(operators))
	copy(copyOperators, operators)

	
	result := math.Pow(numbers[powerIndex], numbers[powerIndex+1])

	operatorsNew := append(copyOperators[:powerIndex], copyOperators[powerIndex+1:]...)
	
	numbersNew := append(copyNumbers[:powerIndex], result)
	numbersNew = append(numbersNew,copyNumbers[powerIndex+2:]...)
	return numbersNew, operatorsNew
}

func removeAllPower(numbers []float64, operators []string) ([]float64, []string) {
	operatorsNew := operators
	numbersNew := numbers
	for  {
		powerIndex, err := findFirstPowerIndices(numbersNew,operatorsNew)
		if err != nil {
			break	
		}	
		numbersNew,operatorsNew  = removeSinglePower(numbersNew, operatorsNew, powerIndex)		
	}

	return numbersNew, operatorsNew
}

