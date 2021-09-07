package main

import "fmt"

func Swapper(arr1, arr2 []int) ([]int, []int) {
	temp := make([]int, len(arr2))
	for i := 0; i < len(arr1); i++ {
		temp[i] = arr1[i]
		arr1[i] = arr2[i]
		arr2[i] = temp[i]
	}
	return arr1, arr2
}
func main() {
	array1 := []int{1, 3, 2, 6}
	array2 := []int{3, 2, 6, 7}
	if len(array1) == len(array2) {
		fmt.Println("Before Swapping")
		fmt.Println(array1, array2)
		fmt.Println("After Swapping")
		fmt.Println(Swapper(array1, array2))
	}
}
