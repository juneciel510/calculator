package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
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

//extract operators from expression
func extractOperators(expression string) []string{
	var operators  []string
	reg, err := regexp.Compile("[0-9.]+")

	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(expression, ",")
	operators_list:= strings.Split(processedString, ",")

	for _, operator := range operators_list {
		if operator!="" {
			operators = append(operators, operator)
		}
	}

	return operators
}

//extract numbers from expression
func extractNumbers(expression string)[]float64{
	var numbers  []float64
	reg, err := regexp.Compile("[^0-9.]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(expression, ",")
	numbers_list := strings.Split(processedString, ",")

	for _, value := range numbers_list {

			number, _ := strconv.ParseFloat(value, 64)
			numbers = append(numbers, number)
		}

	return numbers
}

//set level for operators
func opLevel(op string) int {
	switch op {
	case "*", "/":
		return 1
	default:
		return 0
	}
}

//compare the levels of operators, and return a number
func opCmp(op1 string, op2 string) int {
	return opLevel(op1) - opLevel(op2)
}

//calculate the expression without Parentheses
func noParenthesesCalculation(numbers []float64, operators []string) float64{
	var result float64
	opStack := CreateStack()
	numberStack := CreateStack()
	for index, operator := range operators{
		//if the stack is empty,push two number & one operator to the stacks
		if opStack.Depth()==0 {
			numberStack.Push(numbers[index])
			numberStack.Push(numbers[index+1])
			opStack.Push(operator)
		}else{
			if opCmp(operator,opStack.Top().(string))>0{
			//if the level of operator is higher than the top of the stack
			//pop top of number stack with the number_i+1 to calculate using the current operator
				result=calculateTwoNumber(operator,numbers[index+1],numberStack.Pop().(float64))
				numberStack.Push(result)
			} else{
				result=calculateTwoNumber(opStack.Pop().(string),numberStack.Pop().(float64),numberStack.Pop().(float64))
				numberStack.Push(result)
				numberStack.Push(numbers[index+1])
				opStack.Push(operator)
			}
		}
	}
	result=calculateTwoNumber(opStack.Pop().(string),numberStack.Pop().(float64),numberStack.Pop().(float64))
	return result
}

func calculateTwoNumber(operator string, num1 float64,num2 float64) float64 {
	var result float64
	switch operator {
	case "+":
		result=num1+num2
		break
	case "-":
		result=num2-num1
		break
	case "*":
		result=num1*num2
		break
	case "/":
		result=num2/num1
		break
	}
	return result
}

