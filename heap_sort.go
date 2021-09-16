package main

import "fmt"

func Maxheapify(arr []int, heap_size, i int) {
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
		Maxheapify(arr, len(arr), largest)
	}
}
func Minheapify(arr []int, heap_size, i int) {
	smallest := i
	left := 2 * i
	right := 2*i + 1
	if left < heap_size && arr[left] < arr[smallest] {
		smallest = left
	}
	if right < heap_size && arr[right] < arr[smallest] {
		smallest = right
	}
	if smallest != i {
		arr[i], arr[smallest] = arr[smallest], arr[i]
		Minheapify(arr, len(arr), smallest)
	}
}
func BuildMinHeap(arr []int) ([]int, int) {
	length := len(arr)
	for i := length / 2; i >= 0; i-- {
		Minheapify(arr, length, i)
	}
	return arr, length
}

func BuildMaxHeap(arr []int) ([]int, int) {
	length := len(arr)
	for i := length / 2; i >= 0; i-- {
		Maxheapify(arr, length, i)
	}
	return arr, length
}

func Maxheap_sort(arr []int) []int {
	ar, length := BuildMaxHeap(arr)
	for i := length - 1; i >= 0; i-- {
		ar[0], ar[i] = ar[i], ar[0]
		Maxheapify(ar[:length], length, i)
	}
	return ar
}
func Minheap_sort(arr []int) []int {
	ar, length := BuildMinHeap(arr)
	for i := length - 1; i >= 0; i-- {
		ar[0], ar[i] = ar[i], ar[0]
		Minheapify(ar[:length], length, i)
	}
	return ar
}

func main() {
	arr := []int{1, 3, 2, 5, 8}
	fmt.Println(Maxheap_sort(arr)) //Works perfect
	fmt.Println(Minheap_sort(arr))

}
