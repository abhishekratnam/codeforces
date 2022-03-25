package main

import (
	"fmt"
	"sort"
)

func sorted(ar []int32) []int32 {
	for i := 0; i < len(ar); i++ {
		swap := false
		for j := i + 1; j < len(ar); j++ {
			if ar[i] > ar[j] {
				swap = true
				ar[j], ar[i] = ar[i], ar[j]

			}
			if !swap {
				break
			}
		}

	}
	return ar
}
func migratoryBirds(ar []int32) int32 {
	// Write your code here
	pairs := map[int]int{}
	sort.SliceStable(ar, func(a, b int) bool {
		return ar[a] < ar[b]
	})
	for _, s := range ar {
		if _, ok := pairs[int(s)]; ok {
			pairs[int(s)] += 1
		} else {
			pairs[int(s)] = 1
		}
	}

	i := 0
	keys := make([]int, len(pairs)+1)
	for k := range pairs {
		keys[i] = k
		i++
	}
	sort.Ints(keys)

	b := 0
	fre_sighted_bird := int32(0)
	for i := 0; i < len(keys); i += 1 {
		if pairs[keys[i]] > b {
			fre_sighted_bird = int32(keys[i])
			b = pairs[keys[i]]

		}
	}
	return fre_sighted_bird

}

func main() {
	l := []int32{1, 1, 3, 1, 2, 1, 3, 3, 3, 3}
	fmt.Println(migratoryBirds(l))
}
