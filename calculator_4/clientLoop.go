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
		fmt.Println("Enter expression e.g. 7*8/2+(9-3)^8:")
		scanner.Scan()
		expression := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading text:", err)
		}

		processedExpression:=parse(expression)
		fmt.Println("****************",processedExpression)
		for index,str := range processedExpression{
			fmt.Println("****************",index,str)
		}
		processedExpression=removeAllParentheses(processedExpression)
		fmt.Println("removeAllParentheses",processedExpression)
		result,err:=noParanthesesCalc(processedExpression)
		fmt.Println("result,err",result,err)

	}
}