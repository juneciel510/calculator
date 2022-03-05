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

	// var ct string
	// var result float64
	// var operators  []string
	// var numbers  []float64
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	// var expression string
	//wait for the input from the user
	for {
		var operators  []string
		var numbers  []float64
		fmt.Println("Enter expression (e.g. 3*4.5+9-8/7):")
		scanner.Scan()
		expression := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading text:", err)
		}
		reg, err := regexp.Compile("[0-9.]+")
		if err != nil {
			log.Fatal(err)
		}
		processedString := reg.ReplaceAllString(expression, ",")

		operator_list:= strings.Split(processedString, ",")

		for _, operator := range operator_list {
			if operator!="" {
				operators = append(operators, operator)
			}
		}

		fmt.Println("operators,len(operators):", operators,len(operators))

		reg, err = regexp.Compile("[^0-9.]+")
		if err != nil {
			log.Fatal(err)
		}
		processedString = reg.ReplaceAllString(expression, ",")
		numbers_string := strings.Split(processedString, ",")

		for _, v := range numbers_string {

				number, _ := strconv.ParseFloat(v, 64)
				numbers = append(numbers, number)
			}
		fmt.Println("numbers,len(numbers):", numbers,len(numbers))
		result:=noParenthesesCalculation(numbers,operators)
		fmt.Println("the result of expression is:",expression,result)

	}
}

func opLevel(op string) int {
	switch op {
	case "*", "/":
		return 1
	default:
		return 0
	}
}

func opCmp(op1 string, op2 string) int {
	return opLevel(op1) - opLevel(op2)
}

func noParenthesesCalculation(numbers []float64, operators []string) float64{
	var result float64
	opStack := CreateStack()
	numberStack := CreateStack()
	for index, operator := range operators{
		fmt.Println("------index",index,operator)
		if opStack.Depth()==0 {
			numberStack.Push(numbers[index])
			numberStack.Push(numbers[index+1])
			opStack.Push(operator)
			fmt.Println("numberStack,opStack",numberStack.Top(),numberStack.Depth(),opStack.Depth(),opStack.Top())
		}else{
			//if the level of operator is higher than the top of the stack
			//pop top to calculate
			if opCmp(operator,opStack.Top().(string))>0{
				fmt.Println("operator,opStack.Top()",operator,opStack.Top().(string))
				result=calculateTwoNumber(operator,numbers[index+1],numberStack.Pop().(float64))
				fmt.Println("> result",result)
				numberStack.Push(result)
			} else if opCmp(operator,opStack.Top().(string))<=0{
				result=calculateTwoNumber(opStack.Pop().(string),numberStack.Pop().(float64),numberStack.Pop().(float64))
				fmt.Println("<=0 result",result)
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

