package main

func mergeNegativeNumber(expression []string) []string {
	temp := []string{}
	for index, item := range expression {
		if index > 0 && isNum(item) && (temp[len(temp)-1] == "-") && (len(temp) == 1 || temp[len(temp)-2] == "(") {
			temp = temp[:len(temp)-1]
			number := asFloat(item)
			temp = append(temp, asString(-number))
		} else {
			temp = append(temp, item)
		}
	}
	return temp
}