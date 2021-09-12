package main

import (
	"fmt"
	"sort"
)

func heapify(arr []int, heap_size, i int) {
	largest := i
	left := 2 * i
	right := 2*i + 1
	if left < heap_size && arr[left] > arr[largest] {
		largest = left
	}
	if right < heap_size && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, heap_size, largest)
	}
}
func heapsort(arr []int, heap_size int) {
	for i := heap_size / 2; i >= 0; i-- {
		heapify(arr, heap_size, i)
	}
}
func removeDuplicateValues(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
func findLongestComseqSubseq(arr []int, size int) int {
	count := 1
	// Sort the array
	// heapsort(arr, cap(arr))
	sort.Slice(arr, func(a, b int) bool {
		return arr[a] < arr[b]
	})
	fmt.Println(arr)
	arr = removeDuplicateValues(arr)
	fmt.Println(arr)
	element := make(map[int]int)
	for i := 0; i < len(arr)-1; i++ {
		element[arr[i]] = arr[i]
		element[arr[i+1]] = arr[i+1]
		if element[arr[i+1]]-element[arr[i]] == 1 {
			count += 1
		} else if element[arr[i+1]]-element[arr[i]] > 1 {
			break
		}
	}
	return count
}
func main() {
	//Wrong
	slice := []int{10, 20, 21, 22, 23} //4
	// slice := []int{1, 9, 3, 10, 4, 20, 2} //5
	//Right
	// slice := []int{2, 6, 1, 9, 4, 5, 3} //6
	// slice := []int{6, 6, 2, 3, 1, 4, 1, 5, 6, 2, 8, 7, 4, 2, 1, 3, 4, 5, 9, 6} //9
	// slice := []int{8, 8, 9, 9, 3, 4} //2
	fmt.Println(slice)
	fmt.Println(findLongestComseqSubseq(slice, len(slice)))
}
