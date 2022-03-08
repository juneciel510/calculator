package main

import (
	"errors"
	"fmt"
	"math"
)

//calculate expression have no parentheses
func noParanthesesCalc(expression []string) (float64, error) {
	var err error
	expression,err = removeLetterOperators(expression)
	if err != nil {
		return math.NaN(), err
	}
	expression,err = removeSymbolOperators(expression, []string{"^"})
	if err != nil {
		return math.NaN(), err
	}
	expression,err = removeSymbolOperators(expression, []string{"*", "/"})	
	if err != nil {
		return math.NaN(), err
	}
	expression,err = removeSymbolOperators(expression, []string{"+", "-"})
	if err != nil {
		return math.NaN(), err
	}
	if len(expression) == 1 {
		return asFloat(expression[0]), nil
	}
	return math.NaN(), errors.New("Please check the input")

}

//calculate and remove log, sin, cos from expression
func removeLetterOperators(expression []string) ([]string,error) {
	err:=errors.New("Please check the input")
	temp := []string{}
	index := 0
	var result float64
	for index < len(expression) {
		switch expression[index] {
		case "ln":
			if index+1> len(expression)-1{
				return temp,err
			}
			result = math.Log(asFloat(expression[index+1]))
			temp = append(temp, asString(result))
			index += 2
			break
		case "log":
			if index+2> len(expression)-1{
				return temp,err
			}
			result =math.Log(asFloat(expression[index+2]))/ math.Log(asFloat(expression[index+1]))
			temp = append(temp, asString(result))
			index += 3
			break
		case "sin":
			if index+1> len(expression)-1{
				return temp,err
			}
			result = math.Sin(asFloat(expression[index+1]))
			index += 2
			temp = append(temp, asString(result))
			break
		case "cos":
			if index+1> len(expression)-1{
				return temp,err
			}
			result = math.Cos(asFloat(expression[index+1]))
			index += 2
			temp = append(temp, asString(result))
			break
		default:
			temp = append(temp, expression[index])
			index += 1
		}
	}
	return temp,nil
}

//calculate and remove + - * / ^ from expression
func removeSymbolOperators(expression []string, SymbolOperators []string) ([]string,error){
	temp := []string{}
	index := 0
	var result float64
	for index < len(expression) {
		if contains(SymbolOperators, expression[index]) {
			if index+1>len(expression)-1{
				return temp,errors.New("Please check the input")
			}
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
	return temp,nil
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