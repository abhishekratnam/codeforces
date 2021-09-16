package main

import "fmt"

func bubblesort(arr []int) []int {
	run := true
	for run {
		run = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				run = true
			}
		}
	}
	return arr
}
func main() {
	arr := []int{1, 3, 2, 5, 8}
	fmt.Println(bubblesort(arr))
}
