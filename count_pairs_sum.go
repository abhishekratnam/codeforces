package main

import (
	"fmt"
)

//naive approach
func getPairsCount(arr []int, n, k int) int {
	s := 0
	for l := 0; l < n; l++ {
		for j := l + 1; j < n; j++ {
			if k == arr[l]+arr[j] {
				s += 1
			}
		}
	}
	return s
}

func main() {
	arr := []int{1, 5, 7, 1}
	var k int
	fmt.Println("Enter K:- ")
	fmt.Scanln(&k)
	n := len(arr)
	fmt.Println(getPairsCount(arr, n, k))
}
