/* 
	练习：斐波纳契闭包
	实现一个 fibonacci 函数，它返回一个函数（闭包），该闭包返回一个斐波纳契数列 
		`(0, 1, 1, 2, 3, 5, 7, 9, 11, 13, 15 ...)`。
*/
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	var last_v1 int = 0
	var last_v2 int = 0
	return func(x int) int {
		// fmt.Printf("last_v1 = %d, last_v2 = %d\n", last_v1, last_v2)
		var m int = last_v1 + last_v2

		last_v1, last_v2 = last_v2, x

		return m
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
