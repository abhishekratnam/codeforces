package main

import (
	"fmt"
	"math"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func maxprofit(arr []int, k int) []int {

	var profits []int
	if len(arr) == 0 {
		return profits
	}
	evenProfits := []int{}
	oddProfits := []int{}
	currentProfits := []int{}
	previousProfits := []int{}
	for i := 1; i < k; i++ {
		maxThusFar := math.MinInt32
		if i%2 == 1 {
			currentProfits = oddProfits
			previousProfits = evenProfits
		} else {
			currentProfits = evenProfits
			previousProfits = oddProfits
		}
		for i := 0; i < len(arr); i++ {
			maxThusFar = max(maxThusFar, previousProfits[i-1]-arr[i-1])
			currentProfits[i] = max(currentProfits[i-1], maxThusFar+arr[i])
		}
	}
	return profits
}
func main() {
	//Wrong
	slice := []int{10, 22, 5, 75, 65, 80} //4
	k := 2
	fmt.Println(slice)
	fmt.Println(maxprofit(slice, k))

}
