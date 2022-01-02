package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func runesToString(runes []rune) (outString string) {
	// don't need index so _
	for _, v := range runes {
		outString += string(v)
	}
	return
}
func method(str, op string) string {
	if op == "S" {
		letter1 := []rune{}
		var upperVal rune
		for _, i := range str {
			if unicode.IsUpper(i) && unicode.IsLetter(i) {
				upperVal = i
				break
			}
			if unicode.IsLower(i) && unicode.IsLetter(i) {
				letter1 = append(letter1, i)
			}
		}
		var letter2 string
		for j, i := range str {
			if unicode.IsUpper(i) && unicode.IsLetter(i) {
				if i == upperVal {
					letter2 = str[j : len(str)-2]
					break
				}
			}
		}

		return runesToString(letter1) + " " + strings.ToLower(letter2)
	}
	M := strings.Split(str, " ")
	var res []string
	for i, val := range M {
		var notAppend bool
		if i != 0 {
			for i, j := range val {
				if i == 0 {
					v := string(unicode.ToUpper(j)) + val[i+1:]
					res = append(res, v)
					notAppend = true
					break
				}
			}
		} else {
			if notAppend {
				continue
			}
			res = append(res, val)

		}

	}
	result := strings.Join(res, "")
	return result + "()"

}
func class(str, op string) string {
	if op == "S" {
		var res []string
		res = append(res, " ")
		for _, val := range str {
			if unicode.IsUpper(val) && unicode.IsLetter(val) {
				res = append(res, " ")
			}
			res = append(res, string(val))
		}
		var result []string
		for _, val := range res {
			var notAppend bool
			for i, v := range val {
				if i == 0 {
					notAppend = true
					result = append(result, string(unicode.ToLower(v)))
					break
				}
			}
			if notAppend {
				continue
			}
			result = append(result, val)
		}
		return strings.Join(result, "")
	}

	//else C
	M := strings.Split(str, " ")
	var res []string
	for _, val := range M {
		var notAppend bool
		for i, j := range val {
			if i == 0 {
				v := string(unicode.ToUpper(j)) + val[i+1:]
				res = append(res, v)
				notAppend = true
				break
			}
		}
		if notAppend {
			continue
		}
		res = append(res, val)

	}
	return strings.Join(res, "")

}
func variable(str, op string) string {
	if op == "S" {
		letter1 := []rune{}
		var upperVal rune
		for _, i := range str {
			if unicode.IsUpper(i) && unicode.IsLetter(i) {
				upperVal = i
				break
			}
			if unicode.IsLower(i) && unicode.IsLetter(i) {
				letter1 = append(letter1, i)
			}
		}
		var letter2 string

		for j, i := range str {

			if unicode.IsUpper(i) && unicode.IsLetter(i) {
				if i == upperVal {
					letter2 = str[j:]
					break
				}
			}
		}
		return runesToString(letter1) + " " + strings.ToLower(letter2)
	}
	M := strings.Split(str, " ")
	var res []string
	for i, val := range M {
		if i == 0 {
			res = append(res, M[0])
			continue
		}
		var notAppend bool
		for i, j := range val {
			if i == 0 {
				v := string(unicode.ToUpper(j)) + val[i+1:]
				res = append(res, v)
				notAppend = true
				break
			}
		}
		if notAppend {
			continue
		}
		res = append(res, val)

	}
	return strings.Join(res, "")
	// return " variable "
}
func camelCase(data []string) (string, bool) {
	//Operation Split S or C combine followed by M, C, or V
	op1 := data[0]
	op2 := data[1]
	str := data[2]
	if op1 == "S" {
		switch op2 {
		case "M":
			return method(str, op1), true
		case "C":
			return class(str, op1), true
		case "V":
			return variable(str, op1), true
		}
	}
	if op1 == "C" {
		switch op2 {
		case "M":
			return method(str, op1), true
		case "C":
			return class(str, op1), true
		case "V":
			return variable(str, op1), true

		}
	}

	return "", true

}
func main() {

	scanner := bufio.NewScanner(os.Stdin)
	data := make([]string, 0)
	for {
		// reads user input until \n by default
		scanner.Scan()
		// Holds the string that was scanned
		text := scanner.Text()
		if len(text) != 0 {
			data = append(data, text)
		} else {
			// exit if user entered an empty string
			break
		}
	}
	checkError(scanner.Err())
	for _, val := range data {
		iter := strings.Split(val, ";")
		var result []string
		v, _ := camelCase(iter)
		result = append(result, v)

		fmt.Println(strings.TrimLeft(strings.Join(result, ""), "\t \n"))
	}

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
