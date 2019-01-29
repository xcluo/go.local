package main

import "time"

func gen(i int, ch chan struct{}) {
	for range ch {
		println("i = ", i)
		ch <- struct{}{}
	}
}

func main() {
	ch := make(chan struct{})
	go gen(1, ch)
	go gen(2, ch)
	go gen(3, ch)
	go gen(4, ch)
	ch <- struct{}{}

    time.Sleep(time.Millisecond * 1)	//阻塞
}
