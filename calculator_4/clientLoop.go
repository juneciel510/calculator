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
		fmt.Println("Operators: + - * / ^ sin cos log(number)--natural logarithm")
		fmt.Println("Nested parentheses allowed")
		fmt.Println("Enter expression e.g. 7*(9.5+9-3+9^2/9*32+(log(3+8.9*3/100)^2+3-4+7)*9+sin9-3+9^2/9)*32:")
		fmt.Println("-------------------------------------")
		scanner.Scan()
		expressionInput := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading text:", err)
		}

		expression:=parse(expressionInput)
		expression=removeAllParentheses(expression)
		result,err:=noParanthesesCalc(expression)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("========================================")
		fmt.Println("The result of  ",expressionInput," is:",result)
		fmt.Println("=========================================")

	}
}