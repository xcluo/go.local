package main

import (
	"fmt"
	//"os"
)

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	fmt.Println("触发异常")
}

func main() {
	defer_call()
}
