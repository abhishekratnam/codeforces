package main

import "fmt"

func CountInversions(arr []int, n int) int {
	count := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				count += 1
			}
		}
	}
	return count
}
func main() {
	arr := []int{2, 3, 4, 5, 6}
	n := len(arr)
	fmt.Println(CountInversions(arr, n))

}
