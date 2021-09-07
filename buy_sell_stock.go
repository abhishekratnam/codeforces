package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	max := 0
	sum := 0
	for i := 0; i < len(prices)-1; i++ {
		sum += prices[i+1] - prices[i]
		if sum > max {
			max = sum
		} else if sum < 0 {
			sum = 0
		}
	}
	return max
}

func main() {
	arr := []int{2, 4, 1}
	fmt.Println(maxProfit(arr))
}
