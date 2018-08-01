/*
 * By function closures
 */
package main

import (
	"fmt"
)

// return: (0, 1, 1, 2, 3, 5, ...)
func fibonacci() func(int) int {
	prev, next := 0, 0
	return func(x int) int {
		switch {
		case x == 0:
		case x == 1:
			next = 1
		case x > 1:
			prev, next = next, prev+next
		}
		return next
	}
}
func main() {
	f := fibonacci()
	x := 10
	for i := 0; i < x; i++ {
		fmt.Println(f(i))
	}

}
