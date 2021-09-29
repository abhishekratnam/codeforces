package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(acceleration, initialVelocity, initialDisplacement float64) func(float64) float64 {
	fn := func(time float64) float64 {
		return initialVelocity*time + 0.5*acceleration*math.Pow(time, 2) + initialDisplacement
	}
	return fn
}
func main() {
	fn := GenDisplaceFn(10, 2, 1)
	fmt.Println(fn(3))
	fmt.Println(fn(5))

}
