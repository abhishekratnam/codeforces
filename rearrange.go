package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{-5, -2, 5, 2, 4, 7, 1, 8, 0, -8}
	sort.Ints(arr)
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}	
	fmt.Println(arr)
}
