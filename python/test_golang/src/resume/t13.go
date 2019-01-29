// 下面函数有什么问题？
package main

import "fmt"

// func funcMui(x,y int)(sum int,error){
// func funcMui(x,y int)(sum int, res error){
func funcMui(x,y int)(sum int){
    // return x+y,nil
    // return (x+y), (nil)
	return x + y
}

func main() {
	res := funcMui(1, 2)
	fmt.Println(res)
}

// 函数返回值命名问题：
// 如果有两个返回值，一旦其中一个声明了，另一个也必须声明
// 如果有多个返回值，必须加上括号
// 如果只有一个返回值，并且命名了，也需要加上括号
