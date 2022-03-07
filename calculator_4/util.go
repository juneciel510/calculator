package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func isNum(str string) bool {
	res, err := regexp.MatchString("^[\\d\\.]+$", str)
	if err != nil {
		panic(fmt.Sprintf("regexp.MatchString() error: %s.\n", err.Error()))
	}
	return res
}

//operator: + - * / ^ ( ) 
func isSymbolOp(str string) bool {
	// res, err := regexp.MatchString("^[\\+|\\-|\\*|/|\\^|\\)|\\()|!]$", str)
	res, err := regexp.MatchString("^[\\+|\\-|\\*|/|\\^|\\)|\\(]$", str)
	if err != nil {
		panic(fmt.Sprintf("regexp.MatchString() error: %s.\n", err.Error()))
	}
	return res
}

//operators: log sin cos
func isLetterOp(str string) bool {
	letterOperators := "logsincos"
	res := strings.Contains(letterOperators, str)
	return res
}

// contains checks if a string is present in a slice
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func asFloat(value interface{}) float64 {
	switch value := value.(type) {
	case float64:
		return value
	case string:
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			panic(fmt.Sprintf("strconv.ParseFloat() error: %s.\n", err.Error()))
		}
		return floatValue
	default:
		panic(fmt.Sprintf("calc.asFloat() error: Invalid type %T.", value))
	}
}

func asString(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 32)
}
