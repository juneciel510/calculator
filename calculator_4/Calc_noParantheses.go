package main

import (
	"errors"
	"fmt"
	"math"
)

func noParanthesesCalc(expressionProcessed []string) (float64, error) {
	fmt.Println("********before calc", expressionProcessed)
	expressionProcessed = removeLetterOperators(expressionProcessed)
	fmt.Println("********removeLetterOperators", expressionProcessed)
	expressionProcessed = removeSymbolOperators(expressionProcessed, []string{"^"})
	fmt.Println("********remove ^", expressionProcessed)
	expressionProcessed = removeSymbolOperators(expressionProcessed, []string{"*", "/"})
	fmt.Println("********remove * /", expressionProcessed)
	expressionProcessed = removeSymbolOperators(expressionProcessed, []string{"+", "-"})
	fmt.Println("********remove + -", expressionProcessed)
	if len(expressionProcessed) == 1 {
		return asFloat(expressionProcessed[0]), nil
	}
	fmt.Println(fmt.Sprintf("The result is", expressionProcessed))
	return asFloat(expressionProcessed[0]), errors.New("errors in noParanthesesCalc")

}

func removeLetterOperators(expressionProcessed []string) []string {
	fmt.Println("-===================removeLetterOperators")
	temp := []string{}
	index := 0
	var result float64
	for index < len(expressionProcessed) {
		switch expressionProcessed[index] {
		case "log":
			result = math.Log(asFloat(expressionProcessed[index+1]))
			fmt.Println("log", expressionProcessed[index])
			temp = append(temp, asString(result))
			index += 2
			break
		case "sin":
			result = math.Sin(asFloat(expressionProcessed[index+1]))
			index += 2
			temp = append(temp, asString(result))
			break
		case "cos":
			result = math.Cos(asFloat(expressionProcessed[index+1]))
			index += 2
			temp = append(temp, asString(result))
			break
		default:

			temp = append(temp, expressionProcessed[index])
			fmt.Println("default", expressionProcessed[index], temp)
			index += 1
		}
	}
	return temp
}

func removeSymbolOperators(expressionProcessed []string, SymbolOperators []string) []string {
	fmt.Println("-===================removeSymbolOperators")
	temp := []string{}
	index := 0
	var result float64
	for index < len(expressionProcessed) {
		fmt.Println("---index", index, expressionProcessed[index])
		if contains(SymbolOperators, expressionProcessed[index]) {
			fmt.Println("---index", index, expressionProcessed[index])
			result = calculateSymbolOperators(expressionProcessed[index],
				asFloat(temp[len(temp)-1]),
				asFloat(expressionProcessed[index+1]))
			temp = temp[:len(temp)-1]
			temp = append(temp, asString(result))
			index += 2
		} else {

			temp = append(temp, expressionProcessed[index])

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