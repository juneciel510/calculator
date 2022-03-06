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

		expressionProcessed:=removeAllLogarithm(expression)
		fmt.Println(" expressionProcessed ",expressionProcessed)	
		result:=compoundCalculation(expressionProcessed)
		fmt.Println("The result of expression ",expression," is:",result)


	}
}