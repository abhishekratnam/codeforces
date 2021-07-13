package main

import (
	"fmt"
	"strings"
)

func check(arr []string) bool {
	var (
		i, a  int
		count int = 0
	)
	for k := 0; k < len(arr); k++ {
		if arr[k] == "i" {
			count += 1
			i = k
			break
		}
	}
	for q := i + 1; q < len(arr); q++ {
		if arr[q] == "a" {
			count += 1
			a = q
			break
		}

	}

	for w := a + 1; w < len(arr); w++ {
		if arr[w] == "n" {
			count += 1
			break
		}
	}
	return true
}
func main() {
	fmt.Println("Enter the stirng")
	var word string
	fmt.Scanln(&word)
	arr := strings.Split(word, "")
	//Found := []string{"ian", "Ian", "iuiygaygn", "I d skd a efju N"}
	// NotFound := []string{"ihhhhhn", "ina","xian"}
	//should strat with i and ends with n with a in between
	res := check(arr)
	if res {
		fmt.Println("Found")
	}

}
