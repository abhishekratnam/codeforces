package main

import (
	"fmt"
	"math"
)

func merge(arr, temp_ar []int, left, mid, right int) int {
	i := left
	j := mid + 1
	k := left
	inv_count := 0
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			temp_ar[k] = arr[i]
			k += 1
			j += 1
		} else {
			temp_ar[k] = arr[j]
			inv_count += (mid - i + 1)
			k += 1
			j += 1
		}
	}
	for i <= mid {
		temp_ar[k] = arr[i]
		k += 1
		i += 1
	}
	for j <= right {
		temp_ar[k] = arr[j]
		k += 1
		j += 1
	}
	// for k := left; k <= right; k++ {
	// 	arr[k+i] = temp_ar[k]
	// }
	return inv_count
}
func mergesort(arr, temp_ar []int, left, right int) int {
	inv_count := 0
	if left < right {
		mid := math.Floor(float64(left+right) / float64(2))
		inv_count += mergesort(arr, temp_ar, left, int(mid))
		inv_count += mergesort(arr, temp_ar, int(mid)+1, right)
		inv_count += merge(arr, temp_ar, left, int(mid), right)
	}
	return inv_count
}
func mergeSort(arr []int, array_size int) int {
	temp_ar := make([]int, array_size)
	return mergesort(arr, temp_ar, 0, array_size-1)
}
func main() {
	arr := []int{1, 20, 61, 4, 5}
	fmt.Println(mergeSort(arr, len(arr)))
}
