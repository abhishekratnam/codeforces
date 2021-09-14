package main

import "fmt"

func morethanNdK(arr []int, k int) []int {
	table := make(map[int]int)
	n := len(arr)
	threshold := n / k
	for _, v := range arr {
		_, ok := table[v]
		if ok {
			table[v] += 1
		} else if !ok {
			table[v] = 1
		}
	}
	var out []int
	for key, val := range table {
		if val > threshold {
			out = append(out, key)
		}
	}
	return out
}
func main() {
	arr := []int{3, 1, 2, 2, 1, 2, 3, 3, 3}
	fmt.Println("Enter K value :-")
	var k int
	fmt.Scan(&k)
	fmt.Println(morethanNdK(arr, k))
	arr1 := []int{4, 5, 6, 7, 8, 4, 4}
	k = 3
	fmt.Println(morethanNdK(arr1, k))
}
