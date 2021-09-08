package main

import "fmt"

func SubArraySmartForKelement(pos []int, k int) bool {
	length := len(pos)
	contains := make(map[int]bool)
	sum := 0
	for i := 0; i < length; i++ {
		contains[sum] = true
		sum += pos[i]
		if contains[sum-k] {
			return true
		}
	}
	return false

}
func SubArraySmart(pos []int) bool {
	length := len(pos)
	contains := make(map[int]bool)
	sum := 0
	for i := 0; i < length; i++ {
		contains[sum] = true
		sum += pos[i]
		if contains[sum] {
			return true
		}
	}

	return false

}
func SubArrayNaive(pos []int) bool {
	length := len(pos)
	for i := 0; i < length; i++ {
		sum := 0
		for j := i; j < length; j++ {
			sum += pos[j]
			if sum == 0 {
				return true
			}
		}
	}
	return false

}
func main() {
	slice := []int{2, 1, 3, 4, -2}
	fmt.Println(SubArrayNaive(slice))
	fmt.Println(SubArraySmart(slice))
}
