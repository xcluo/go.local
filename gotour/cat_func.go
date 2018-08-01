package main

import (
	"fmt"
	"math"
)

// function as parameters
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// function closures
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	hypot := func(x, y float64) float64 {
		return x + y
	}

	// fmt.Println(hypot(3, 4))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	// fmt.Println("Function closures: \n")
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(adder()(i), adder()(-2*8))
	// }

	// pos, neg := adder(), adder()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(pos(i), neg(-2*8))
	// }
	bm := adder()
	fmt.Println(bm(2))
	fmt.Println(bm(2))
	fmt.Println(bm(2))
}
