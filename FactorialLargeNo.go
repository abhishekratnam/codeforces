package main

import "fmt"

func Factorial(num int) int {
	// slice := []int{1}
	if num == 0 || num == 1 {
		return 1
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
