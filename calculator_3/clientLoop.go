package main

import (
	"bufio"
	"fmt"
	"os"
)

func clientLoop() {
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

	//wait for the input from the user and print result
	for {
		var operators  []string
		var numbers  []float64
		fmt.Println("Enter expression e.g. 6*8+(7-8/9+6)*(6+9*8):")
		scanner.Scan()
		expression := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading text:", err)
		}
		// for _, charRune := range expression {
		// 	char := string(charRune)
		// 	fmt.Println("char:", char)
		// }

		numbers=extractNumbers(expression)
		operators=extractOperators(expression)
		numbersNew, operatorsNew:=removeAllParentheses(numbers, operators)
		result:=noParenthesesCalculation(numbersNew, operatorsNew)
		fmt.Println("The result of expression ",expression," is:",result)


	}
}