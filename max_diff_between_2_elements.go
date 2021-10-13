package main

import "fmt"

func maxdiffbet2elements(arr []int) int {
	max_diff_sofar := arr[1] - arr[0]
	min_element := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i]-min_element > max_diff_sofar {
			max_diff_sofar = arr[i] - min_element
		}
		if arr[i] < min_element {
			min_element = arr[i]
		}
	}
	return max_diff_sofar
}
func main() {
	//Wrong
	slice := []int{10, 22, 5, 75, 65, 80} //4
	fmt.Println(slice)
	fmt.Println(maxdiffbet2elements(slice))

}
