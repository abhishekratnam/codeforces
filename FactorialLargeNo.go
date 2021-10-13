package main

import (
	"fmt"
	"math/big"
)

func Factorial(num int) big.Int {
	// slice := []int{1}
	if num == 0 || num == 1 {
		return big.Int(1)
	}
	a := 1
	for i := 2; i < num+1; i++ {
		a = i * a
	}
	return a
}
func main() {
	fmt.Println("Enter Number")
	var num int
	fmt.Scan(&num)
	fmt.Println(Factorial(num))
}
