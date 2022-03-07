package main

//calculation can handle Parentheses, power, and +,-,*,/
func compoundCalculation(expression string) float64 {
	var operators  []string
	var numbers  []float64
	numbers=extractNumbers(expression)	
	if len(numbers) == 1 {
		return numbers[0]
	}
	operators=extractOperators(expression)
	numbers, operators = removeAllParentheses(numbers, operators)
	numbers, operators = removeAllPower(numbers, operators)
	result := simpleCalculation(numbers, operators)
	return result
}
