package main

import (
	"fmt"
)

func intersectionCommonElements(arr1, arr2 []int) (inter []int) {
	check := make(map[int]bool)
	for _, val := range arr1 {
		check[val] = true
	}
	for _, val := range arr2 {
		if check[val] {
			inter = append(inter, val)
		}
	}
	inter = removeDuplicates(inter)
	return inter
}
func removeDuplicates(arr []int) (sets []int) {
	checklist := make(map[int]bool)
	for _, val := range arr {
		if !checklist[val] {
			checklist[val] = true
			sets = append(sets, val)
		}
	}
	return sets
}
func main() {
	arr1 := []int{1, 5, 10, 20, 40, 80}
	arr2 := []int{6, 7, 20, 80, 100}
	arr3 := []int{3, 4, 15, 20, 30, 70, 80, 120}
	arrtemp := intersectionCommonElements(arr1, arr2)
	fmt.Println(intersectionCommonElements(arrtemp, arr3))
}
