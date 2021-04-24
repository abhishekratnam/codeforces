package main

import "fmt"

func main() {
	fmt.Println("Enter your array")
	var result []int
	var i int
	fmt.Scan(&i)
	arr := make([]int, i)
	for j := 0; j < i; j++ {
		fmt.Scan(&arr[j])
	}
	for j := len(arr) - 1; j >= 0; j-- {
		result = append(result, arr[j])
	}
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}

}
