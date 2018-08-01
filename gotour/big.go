package main

import (
	"fmt"
	"math"
)

func main() {
	sum, sp := 0, 0
	for i := 0; i < 10; i++ {
		sum += 1
	}

	j := 0
	for ; j < 10; j++ {
		sp += j + j
	}
	print(sum, sp)
	print("a", "b", "C")

	print("\n")

	if v, p := 30.0, 20.0; v > p {
		print(true)
		print(math.Pow(v, p))
	} else {
		print(false)
		print(math.Pow(v, p) * 10)
	}
	fmt.Print("\nabcvd")

	switch mac := "ios9"; mac {
	case "ios10":
		fmt.Println("\nOS X")
	case "ios9":
		fmt.Println("\nOS Y")
	}

	for k := 0; k < 10; k++ {
		defer fmt.Println("k = ", k)
	}
	fmt.Println("Done")
}
