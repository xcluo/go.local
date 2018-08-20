package main

import (
	"fmt"
	// "math"
	// "math/rand"
)

const (
	PP = 3.33333
	CC = "abcde"
)

var c, python, java, ruby, golang bool

func add(x, y int) int {
	return x * y
}

func swap(x, y int) (m1, m2 string) {
	m1, m2 = "a", "b"
	return m1, m2
}

func main() {
	// var i, j int
	// k := 0
	// var m = 0
	// fmt.Println(i, j, k, m)

	var m, j, k = 1, 2, 3
	// var m, j, k int = 1, 2, 3
	fmt.Println(m, j, k)
	// start new
	// fmt.Println(rand.Intn(10))
	// fmt.Println(math.Pi)
	// fmt.Println(add(2, 3))
	// fmt.Print(swap(2, 3))

	v := 42.
	fmt.Printf("%T\n", v)
	fmt.Printf("%T\n", PP)
	fmt.Printf("%T\n", CC)
}
