package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)


func extractOperators(expression string) []string{
	var operators  []string
	reg, err := regexp.Compile("[0-9.]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(expression, ",")
	
	operator_list:= strings.Split(processedString, ",")
	
	for _, v := range operator_list {
		if v!="" {
			for _, charRune := range v {
				char := string(charRune)
				// fmt.Println("char:", char)
				operators = append(operators, char)	
			}
				
		}
	}

	return operators
}

func extractNumbers(expression string)[]float64{
	var numbers  []float64
	reg, err := regexp.Compile("[^0-9.]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(expression, ",")
	numbers_string := strings.Split(processedString, ",")
	
	for _, v := range numbers_string {
		if v!="" {	
			number, _ := strconv.ParseFloat(v, 64)
			numbers = append(numbers, number)}
		}
	return numbers
}
func opLevel(op string) int {
	switch op {
	case "^":
		return 2
	case "*", "/":
		return 1
	default:
		return 0
	}
}

func opCmp(op1 string, op2 string) int {
	return opLevel(op1) - opLevel(op2)
}

//calculate the expression only contains numbers and +,-,*,/
func simpleCalculation(numbers []float64, operators []string) float64{
	var result float64
	opStack := CreateStack()
	numberStack := CreateStack()
	for index, operator := range operators{
		fmt.Println("---------index, operator",index, operator)
		if opStack.Depth()==0 {
			numberStack.Push(numbers[index])
			numberStack.Push(numbers[index+1])
			opStack.Push(operator)
		}else{
			if opCmp(operator,opStack.Top().(string))>0{
				fmt.Println("------   >0:",opStack.Top().(string))
				result=calculateTwoNumber(operator,numbers[index+1],numberStack.Pop().(float64))
				fmt.Println("result:",result)
				numberStack.Push(result)
			} else{
				fmt.Println("------   <=0:",opStack.Top().(string))
				result=calculateTwoNumber(opStack.Pop().(string),numberStack.Pop().(float64),numberStack.Pop().(float64))
				fmt.Println("result:",result)
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
	default:
		fmt.Println("input false operators")
	}
	return result
}

