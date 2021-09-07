package main

import (
	"fmt"
	"strings"
)

func main() {
	var data string
	fmt.Println("Enter String")
	fmt.Scanln(&data)
	fmt.Println("Vowels Count for the ", data, "are = ")
	data = strings.ToLower(data)
	checklist := map[string]bool{"a": true, "e": true, "i": true, "o": true, "u": true}
	split := strings.Split(data, "")

	var count = 0
	for _, val := range split {
		if _, ok := checklist[val]; ok {
			count += 1
		}
	}
	fmt.Println(count)
}
