package main

import "fmt"

func heapify(arr []int, n, i int) {
	largest := i
	l := 2 * i
	r := 2*i + 1
	if l < n && arr[l] > arr[largest] {
		largest = l
	}
	if r < n && arr[r] > arr[largest] {
		largest = r
	}
	if largest != i {
		arr[largest], arr[i] = arr[i], arr[largest]
		heapify(arr, n, largest)
	}
}
func heapsort(arr []int, n int) {
	for i := n / 2; i >= 0; i-- {
		heapify(arr, n, i)
	}
}
func main() {
	arr := []int{12, 11, 13, 5, 6, 7}
	n := len(arr)
	var result []int
	heapsort(arr, n)
	for i := 0; i < n; i++ {
		result = append(result, arr[i])
	}

	fmt.Println("Sorted arry is: \n", result)
	fmt.Println("Max is: ", result[len(arr)-1])
	fmt.Println("Min is: ", result[0])
}
