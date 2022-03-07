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
		fmt.Println("----------Read before use-----------")
		fmt.Println("Operators: + - * / ^ log(number)--natural logarithm, log(number,base)")
		fmt.Println("Nested parentheses not allowed")
		fmt.Println("Enter expression e.g. 7*(8.7/2-3)+log(9.5-3/8)^8:")

		scanner.Scan()
		expression := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading text:", err)
		}

		expressionProcessed:=removeAllLogarithm(expression)
		result:=compoundCalculation(expressionProcessed)
		fmt.Println("========================================")
		fmt.Println("The result of  ",expression," is:",result)
		fmt.Println("=========================================")


	}
}