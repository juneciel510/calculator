package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func clientLoop() {

	var operator string
	var result float64
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	//wait for the input from the user
	for {
		fmt.Println("Enter operator(+,-,*,/):")
		fmt.Scanln(&operator)

		fmt.Println("Enter numbers(separate with comma):")
		scanner.Scan()
		txt := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading text:", err)
		}
		s := strings.Split(txt, ",")
		numbers := []float64{}
		for _, v := range s {
			number, _ := strconv.ParseFloat(v,64)
			numbers = append(numbers, number)
		}
		// fmt.Println("numbers",numbers)
		switch operator {
		case "+":
			result=addition(numbers)
			break
		case "-":
			result=subtraction(numbers)
			break
		case "*":
			result=multiplication(numbers)
			break
		case "/":
			result=division(numbers)
			break
		}
		fmt.Println("the result is:",result)
	}
}

func addition(numbers []float64) float64{
	result := 0.0
	for _, v := range numbers {
	result += v
	}
	return result
}
func subtraction(numbers []float64) float64{
	var result float64
	for index, v := range numbers {
		if index==0{
			result=v
		}else{
			result -= v
		}
	}
	return result
}

func multiplication(numbers []float64) float64{
	result := 1.0
	for _, v := range numbers {
	result *= v
	}
	return result
}

func division(numbers []float64) float64{
	var result float64
	for index, v := range numbers {
		if index==0{
			result=v
		}else{
			result /= v
		}
	}
	return result
}