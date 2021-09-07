package main

import (
	"fmt"
	"sort"
)

func main() {

	numbers := make([]int, 3)
	var userinput rune

	for {
		userinput = 0
		fmt.Println("Enter a number: ")
		fmt.Scan(&userinput)
		if userinput == 0 {
			fmt.Println("Finished")
			break
		}
		numbers = append(numbers, int(userinput))

	}
	sort.Ints(numbers)
	newNumbers := append(numbers[3:])
	fmt.Println(newNumbers)
}
