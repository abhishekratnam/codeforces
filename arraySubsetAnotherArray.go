package main

import "fmt"

//Optimized approach
func subSetArrayOptimized(arr, subarr []int) string {
	count := 0
	m := make(map[int]bool, len(subarr))
	for _, j := range subarr {
		_, ok := m[j]
		if !ok {
			m[j] = true
		}
	}
	for _, v := range arr {
		if m[v] {
			count += 1
		}
	}
	if count == len(subarr) {
		return "Yes"
	}
	return "No"
}

//Naive approach
func subSetArray(arr, subarr []int) string {
	count := 0
	for _, v := range arr {
		for _, j := range subarr {
			if j == v {
				count += 1
			}
		}
	}
	if count == len(subarr) {
		return "Yes"
	}
	return "No"
}
func main() {
	array1 := []int{10, 5, 2, 23, 19}
	subSet := []int{19, 5, 3}
	fmt.Println(subSetArray(array1, subSet))
	fmt.Println(subSetArrayOptimized(array1, subSet))

}
