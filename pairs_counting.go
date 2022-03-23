package main

import (
	"fmt"
)

/*
 * Complete the 'sockMerchant' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY ar
 */

func sockMerchant(n int32, ar []int32) int32 {
	// Write your code here
	pairs := map[int32]int32{}
	for _, s := range ar {
		if _, ok := pairs[s]; ok {
			pairs[s] += 1
		} else {
			pairs[s] = 1
		}
	}

	pair := int32(0)
	for key, val := range pairs {
		fmt.Println(key, val)
		for i := int32(1); i < val; i += 2 {
			pair += 1
		}
	}
	return pair

}

func main() {
	l := []int32{1, 1, 3, 1, 2, 1, 3, 3, 3, 3}
	fmt.Println(sockMerchant(10, l))
}
