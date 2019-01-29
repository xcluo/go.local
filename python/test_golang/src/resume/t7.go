// 请写出以下输入内容
package main

import "fmt"

func main() {
    s := make([]int, 5)		// 这里是有默认值的
	// s := make([]int, 0)  // 这里就没有默认值了
    s = append(s, 1, 2, 3)
    fmt.Println(s)
}
// [0, 0, 0, 0, 0, 1, 2, 3]
