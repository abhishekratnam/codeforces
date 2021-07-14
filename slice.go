package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var size int
	k := make([]int, size)
	for {
		fmt.Println("ENter INtegers:-")
		var no int
		fmt.Scanln(&no)
		k = append(k, no)
		sort.Slice(k, func(i, j int) bool {
			return k[i] < k[j]
		})
		fmt.Println(k)
		r := bufio.NewReader(os.Stdin)
		text, ret := r.ReadString('\n')
		if ret == nil {
			text = strings.TrimSpace(text)
			if text == "X" {
				break
			}
		}

	}
}
