package main

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strings"
)

func removeSingleLogarithm(expression string) (string,error) {
	logreg := regexp.MustCompile(`log\((.*?)\)`)
	
	found := logreg.FindString(expression)
	// indexRange := logreg.FindStringIndex(expression)
	if found == ""{
		return expression, errors.New("Logrithm not found")
	}
	expressionInLog := strings.Trim(found, "log(")
	expressionInLog = strings.Trim(expressionInLog, ")")
	subExpressions := strings.Split(expressionInLog, ",")

	var result float64
	switch len(subExpressions) {
		case 1:
			number:=compoundCalculation(subExpressions[0])
			result=math.Log(number)
			break
		case 2:
			number:=compoundCalculation(subExpressions[0])
			base:=compoundCalculation(subExpressions[1])
			result=math.Log(number)/math.Log(base)
			break
		default:
			return expression, errors.New("the logrithm expression is not correct, must be log(x)--natural logarithm, log(number,base)")
	}

	resultStr:=fmt.Sprintf("%f", result) 
	expressionProcessed:=strings.Replace(expression, found, resultStr, 1)
	
	return expressionProcessed,nil
}

func removeAllLogarithm(expression string) string {
	expressionNew := expression
	var err error
	for  {
		expressionNew, err = removeSingleLogarithm(expressionNew)
		if err != nil {
			break	
		}		
	}
	return expressionNew
}