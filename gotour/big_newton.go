package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	plus := 0.000001
	counter := 100000 * 10000

	diff := x
	for i := 0; i < counter; i++ {
		if tmp := x - z*z; tmp > 0.0 && tmp < diff {
			diff = tmp
		} else {
			fmt.Printf("\nBreak(%d): z = %f, tmp = %f, diff = %f", i, z, tmp, diff)
			break
		}
		z += plus
	}
	return z
}

func Newton(x float64) float64 {
	z := 1.0
	z1 := 0.0

	for i := 0; i < 100; i++ {
		z, z1 = z1, -(z*z+x)/(2*z)

		if z*z != z1*z1 {
			z = z1
		} else {
			fmt.Printf("Break(%d), \tz=%f, \tz1=%f \tx-z*z=%f, \tz!=z2(%t)", i, z, z1, x-z*z, z != z1)
			break
		}
	}
	return z1
}

func main() {
	// x := 9238.0
	x := 9223372036854775807.0
	x = math.Sqrt(x)
	fmt.Printf("\nx: %f \n", x)
	x = math.Sqrt(x)
	fmt.Printf("\nx: %f \n", x)
	x = math.Sqrt(x)
	fmt.Printf("\nx: %f \n", x)
	x = math.Sqrt(x)
	fmt.Printf("\nx: %f \n", x)
	x = math.Sqrt(x)
	fmt.Printf("\nx: %f \n", x)
	x = math.Sqrt(x)
	fmt.Printf("\nx: %f \n", x)
	x = math.Sqrt(x)
	fmt.Printf("\nx: %f \n", x)
	// fmt.Printf("\nInteractor Results: %f \n", Sqrt(x))
	fmt.Printf("\nMath Results: %f \n\n", math.Sqrt(x))
	// fmt.Printf("\nNewton Result: %f\n", Newton(x))
}
