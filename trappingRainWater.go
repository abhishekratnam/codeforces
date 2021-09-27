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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func trappingRainWater(arr []int, length int) int {

	left := []int{}
	right := []int{}
	left = append(left, arr[0])
	minima := arr[length-1]
	maxima := arr[0]

	for i := 1; i < length; i++ {
		arr_max := max(arr[i-1], arr[i])
		if maxima < arr_max {
			maxima = arr_max
		}
		left = append(left, maxima)
	}
	right = append(right, arr[length-1])
	for i := length - 2; i >= 0; i-- {
		arr_min := max(arr[i], arr[i+1])
		if minima < arr_min {
			minima = arr_min
		}
		right = append(right, minima)
	}
	//mistake
	answer := 0
	for i := 0; i < length-1; i++ {
		answer += min(left[i], right[i]) - arr[i]
	}
	fmt.Println(left, right)
	return answer

}
func main() {
	arr := []int{3, 1, 2, 4, 0, 1, 3, 2}
	fmt.Println(trappingRainWater(arr, len(arr)))
}
