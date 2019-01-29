// 是否可以编译通过？如果通过，输出什么？
package main

import "fmt"

const (
	x = iota
	y
	z = "zz"
	k
	p = iota
)

func main()  {
	fmt.Println(x,y,z,k,p)
}

// 0 1 zz zz 4
// iota 枚举
