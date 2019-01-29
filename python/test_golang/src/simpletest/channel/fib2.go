package main

import "fmt"
import "time"

func fib(c, quit chan int) {
    x, y := 0, 1
    for {
        select {
            case c <- x:
                time.Sleep(time.Millisecond*3000)
                x, y = y, x +y
                fmt.Println("computing")
            case <- quit:
                fmt.Println("quit")
                fmt.Println("res=", x)
                return 
        }
    }
}

func main() {
    c := make(chan int)
    quit := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            //time.Sleep(time.Millisecond*1000)
            fmt.Println("c = ", <-c)
        }
        quit <- 0
    }()
    fib(c, quit)
}
