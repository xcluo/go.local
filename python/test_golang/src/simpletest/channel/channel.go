package main

import "fmt"

func sum(s []int, c chan int) {
             sum := 0
             for _, v := range s {
                 sum  += v
             }

            c <- sum
}

func main() {
    s := []int{1, 3, 5, 7, 9}
    c := make(chan int)

    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)
    x, y := <-c, <-c

    fmt.Println("len=", len(s)/2)
    fmt.Println("x=", x)
    fmt.Println("y=", y)
}
