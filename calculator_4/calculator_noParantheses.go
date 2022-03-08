package main

import (
	"errors"
	"fmt"
	"math"
)

//calculate expression have no parentheses
func noParanthesesCalc(expression []string) (float64, error) {
	expression = removeLetterOperators(expression)
	expression = removeSymbolOperators(expression, []string{"^"})
	expression = removeSymbolOperators(expression, []string{"*", "/"})	
	expression = removeSymbolOperators(expression, []string{"+", "-"})
	if len(expression) == 1 {
		return asFloat(expression[0]), nil
	}
	return asFloat(expression[0]), errors.New("errors in noParanthesesCalc")

}

//calculate and remove log, sin, cos from expression
func removeLetterOperators(expression []string) []string {
	temp := []string{}
	index := 0
	var result float64
	for index < len(expression) {
		switch expression[index] {
		case "ln":
			result = math.Log(asFloat(expression[index+1]))
			temp = append(temp, asString(result))
			index += 2
			break
		case "log":
			result =math.Log(asFloat(expression[index+2]))/ math.Log(asFloat(expression[index+1]))
			temp = append(temp, asString(result))
			index += 3
			break
		case "sin":
			result = math.Sin(asFloat(expression[index+1]))
			index += 2
			temp = append(temp, asString(result))
			break
		case "cos":
			result = math.Cos(asFloat(expression[index+1]))
			index += 2
			temp = append(temp, asString(result))
			break
		default:
			temp = append(temp, expression[index])
			index += 1
		}
	}
	return temp
}

//calculate and remove + - * / ^ from expression
func removeSymbolOperators(expression []string, SymbolOperators []string) []string {
	temp := []string{}
	index := 0
	var result float64
	for index < len(expression) {
		if contains(SymbolOperators, expression[index]) {
			result = calculateSymbolOperators(expression[index],
				asFloat(temp[len(temp)-1]),
				asFloat(expression[index+1]))
			//remove last element
			temp = temp[:len(temp)-1]
			temp = append(temp, asString(result))
			index += 2
		} else {
			temp = append(temp, expression[index])
			index += 1
		}
	}
	return temp
}

func calculateSymbolOperators(operator string, num1 float64, num2 float64) float64 {
	var result float64
	switch operator {
	case "+":
		result = num1 + num2
		break
	case "-":
		result = num1 - num2
		break
	case "*":
		result = num1 * num2
		break
	case "/":
		result = num1 / num2
		break
	case "^":
		result = math.Pow(num1, num2)
		break
	default:
		fmt.Println("input false operators")
	}
	return result
}