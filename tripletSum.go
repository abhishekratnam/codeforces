package main

import (
	"fmt"
	"sort"
)

func TripletSumArray(arr []int, n, x int) int {
	sort.Slice(arr, func(a, b int) bool {
		return arr[a] < arr[b]
	})

	for i := 0; i < n-2; i++ {
		j := i + 1
		k := n - 1
		sum := x - arr[i]
		for j < k {
			if arr[j]+arr[k] == sum {
				return 1
			}
			if arr[j]+arr[k] < sum {
				j += 1
			} else {
				k -= 1
			}
		}
	}
	return 0
}
func main() {
	array1 := []int{1, 4, 45, 6, 10, 8}
	array2 := []int{1, 2, 4, 3, 6}
	n := 6
	x := 13

	fmt.Println(TripletSumArray(array1, n, x))
	fmt.Println(TripletSumArray(array2, n, x))

}
