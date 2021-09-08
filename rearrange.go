package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{-5, -2, 5, 2, 4, 7, 1, 8, 0, -8}
	sort.Ints(arr)
	fmt.Println("After Sorting")
	fmt.Println(arr)

	i, j := 1, 1
	for j < len(arr) {
		if arr[j] > 0 {
			break
		}
		j += 1
	}
	for arr[i] < 0 && j < len(arr) {
		arr[i], arr[j] = arr[j], arr[i]
		i += 1
		j += 1
	}
	fmt.Println("Result Array")
	fmt.Println(arr)
}
