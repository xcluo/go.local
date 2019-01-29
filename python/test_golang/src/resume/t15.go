// 是否可以编译通过？如果通过，输出什么？
package main

import "fmt"

func main() {
	list := new([]int)
	// list := make([]int, 0)	// new 类似于，返回指针
	fmt.Println(list)

	list = append(list, 1)
	// *list = append(*list, 1)		// 所以，需要使用 * 取内容

	fmt.Println(*list)
}

// [1]
// list := make([]int, 0)
