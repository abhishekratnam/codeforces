package main

import "fmt"

func isTrue(pos []int) bool {
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
	slice := []int{2, 1, -3, 4, 2}
	fmt.Println(isTrue(slice))
}
