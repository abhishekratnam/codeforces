package main

import (
	"fmt"
	"sort"
)

//Naive Approach
func TripletSumArrayOptimized(arr []int, n, x int) int {
	sort.Slice(arr, func(a, b int) bool {
		return arr[a] < arr[b]
	})
	for i := 0; i < n-2; i++ {
		j := i + 1
		k := n - 1
		for j < k {
			computed_sum := arr[i] + arr[j] + arr[k]
			if computed_sum == x {
				return 1
			} else if computed_sum < x {
				j += 1
			} else if computed_sum > x {
				k -= 1
			}
		}
	}
	return 0

}

//Naive Approach
func TripletSumArray(arr []int, n, x int) int {
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if arr[i]+arr[j]+arr[k] == x {
					return 1
				}
			}
		}
	}
	return 0
}
func main() {
	array1 := []int{1, 4, 45, 6, 10, 8}
	// array2 := []int{1, 2, 4, 3, 6}
	n := 6
	x := 13

	fmt.Println(TripletSumArray(array1, n, x))
	fmt.Println(TripletSumArrayOptimized(array1, n, x))

}
