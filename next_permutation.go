package main

import (
	"fmt"
)

func next_permutation(arr []int) []int {
	length := len(arr)
	for i := length - 1; i > 0; i-- {
		if arr[i] > arr[i-1] {
			for j := i - 1; j < length; j++ {
				if arr[j] < arr[i] {
					arr[j], arr[i] = arr[i], arr[j]
				}
			}
			// Reverse the element
			for j := i + 1; j < i+1+(length-i-1)/2; j++ {
				arr[j], arr[length-j+i] = arr[length-j+i], arr[j]
			}
			return arr
		}
	}
	return nil
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(next_permutation(arr))
}
