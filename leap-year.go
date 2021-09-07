package main

import (
	"fmt"
)

func isLeapYear(year int) bool {
	leapFlag := false
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				leapFlag = true
			} else {
				leapFlag = false
			}
		} else {
			leapFlag = true
		}
	} else {
		leapFlag = false
	}
	return leapFlag
}
func main() {
	bool := isLeapYear(2021)
	bool1 := isLeapYear(2022)
	fmt.Println(" 2021 leap year?:", bool)
	fmt.Println(" 2022 leap year?:", bool1)

}
