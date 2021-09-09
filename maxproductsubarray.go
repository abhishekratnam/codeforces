package main

import (
	"fmt"
)

func maxProduct(arr []int) int {
	max_ending_at_i := arr[0]
	min_ending_at_i := arr[0]
	temp := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < 0 {
			max_ending_at_i, min_ending_at_i = min_ending_at_i, max_ending_at_i
		}
		max_ending_at_i *= arr[i]
		if arr[i] > max_ending_at_i {
			max_ending_at_i = arr[i]
		}
		min_ending_at_i *= arr[i]
		if arr[i] < min_ending_at_i {
			min_ending_at_i = arr[i]
		}
		if max_ending_at_i > temp {
			temp = max_ending_at_i
		}
	}
	return temp
}
func main() {
	slice := []int{6, -3, -10, 0, 2}
	fmt.Println(slice)
	fmt.Println(maxProduct(slice))
}
