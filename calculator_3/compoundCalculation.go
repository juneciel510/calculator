package main

import "fmt"

//calculation can handle Parentheses, power, and +,-,*,/
func compoundCalculation(expression string) float64 {
	var operators  []string
	var numbers  []float64
	numbers=extractNumbers(expression)	
	if len(numbers) == 1 {
		return numbers[0]
	}
	operators=extractOperators(expression)
	fmt.Println("numbers,operators:", numbers,operators)
	numbers, operators = removeAllParentheses(numbers, operators)
	fmt.Println("numbers, operators:", numbers, operators)
	numbers, operators = removeAllPower(numbers, operators)
	fmt.Println("numbers, operators:", numbers, operators)
	result := simpleCalculation(numbers, operators)
	return result
}
