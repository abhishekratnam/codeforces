package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func trappingRainWater(arr []int, length int) int {

	left := []int{}
	right := []int{}
	left = append(left, arr[0])
	right = append(right, arr[length-1])
	maxima := arr[0]
	minima := arr[length-1]
	final := 0
	for i := 1; i < length; i++ {
		arr_max := max(arr[i-1], arr[i])
		if maxima < arr_max {
			maxima = arr_max
		}
		left = append(left, maxima)

	}
	var min int
	for i := length - 2; i >= 1; i-- {
		arr_min := max(arr[i], arr[i+1])
		if minima < arr_min {
			minima = arr_min
		}
		right = append(right, minima)
	}
	//mistake
	for i, j := 0, 0; i < len(left) && j < len(right); j++ {
		if left[i] > right[j] {
			min = right[j]
			final += min
			fmt.Println(min)
		} else if left[i] < right[j] {
			min = left[i]
			final += min
			fmt.Println(min)
		}

		i += 1
	}
	fmt.Println(left, right)
	return final

}
func main() {
	arr := []int{3, 1, 2, 4, 0, 1, 3, 2}
	fmt.Println(trappingRainWater(arr, len(arr)))
}
