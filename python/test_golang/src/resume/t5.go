// 下面代码会触发异常吗？请详细说明
package main

import "fmt"
import "runtime"

func main() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {			// select 随机，有可能触发
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)	// 这里会触发异常
	}
}
