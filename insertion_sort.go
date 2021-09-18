package main

import "fmt"

func Swap(a *int, b *int) {
	*a, *b = *b, *a
}
func Insertion_sort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			Swap(&arr[j], &arr[j-1])
		}
	}
	return arr
}

func main() {
	arr := []int{1, 3, 2, 5, 8}
	fmt.Println(Insertion_sort(arr)) //Works perfect

}
