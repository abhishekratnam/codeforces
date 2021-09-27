package main

import "fmt"

func Merge(left, right []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			final = append(final, left[i])
			i += 1
		} else {
			final = append(final, right[j])
			j += 1
		}
	}
	for ; i < len(left); i++ {
		final = append(final, left[i])
	}
	for ; j < len(right); j++ {
		final = append(final, right[j])
	}
	return final
}
func MergeSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	middle := length / 2
	left := MergeSort(arr[:middle])
	right := MergeSort(arr[middle:])
	return Merge(left, right)
}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Distribution(arr []int, N, M int) int {
	arr = MergeSort(arr)

	const MaxUint = ^uint(0)
	const MinUint = 0
	MaxInt := int(MaxUint >> 1)

	ans := -MaxInt - 1
	for i := 0; i+M-1 < N; i++ {
		k := arr[i+M-1] - arr[i]
		if k < ans {
			ans = k
		}
	}

	return ans
}
func main() {
	chocolates := []int{3, 4, 1, 9, 56, 7, 9, 12}
	var N int
	var M int
	fmt.Println(Distribution(chocolates, N, M))
}
