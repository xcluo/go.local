package main

import (
    "fmt"
    "time"
)

func main() {
    tick := time.Tick(time.Millisecond * 500)
    boom := time.After(time.Millisecond * 1000)
    for {
        select {
            case x := <- tick:
                fmt.Println("tick = ", x)
            case y := <- boom:
                fmt.Println("boom = ", y)
                return
            default:
                fmt.Println("    .")
                time.Sleep(50 * time.Millisecond)
            //return
        }
    }
    fmt.Println("done")
}
