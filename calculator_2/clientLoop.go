package main

import (
	"bufio"
	"fmt"
	"os"
)

func clientLoop() {
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	//wait for the input from the user, then print result
	for {
		fmt.Println("Enter expression (e.g. 3*4.5+9-8/7):")
		scanner.Scan()
		expression := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading text:", err)
		}
		
		operators:=extractOperators(expression)
		numbers:=extractNumbers(expression)
		
		result:=noParenthesesCalculation(numbers,operators)
		fmt.Println("the result of expression ",expression, ":",result)

	}
}