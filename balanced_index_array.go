package main

import "fmt"

//given an array of integers, find the first
// index such that the sum of all the integers before
// the index is equal to the sum of all the integers after the index.
func sum(arr []int) int {
	count := 0
	for _, v := range arr {
		count += v
	}
	return count
}
func nthIndex(arr []int) int {
	// Traverse the array
	var i int
	for i := 0; i < len(arr)-1; i++ {
		left_array := arr[:i] //array slice
		right_array := arr[i+1:]
		if sum(left_array) == sum(right_array) {
			return i
		}
	}
	return i
}
func main() {
	arr := []int{1, 4, 1, 2, 2, 1}
	// arr := []int{0, 0, 0}
	result := nthIndex(arr)
	if result != 0 {
		fmt.Println("Index := ", result)
	} else {
		fmt.Println("Index Not Found := ", result)
	}
}
