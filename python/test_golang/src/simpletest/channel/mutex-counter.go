/*
   互斥测试
*/
// Mutual exclusion
package main

import "fmt"
import "sync"
// import "time"

type SafeCounter struct {
    v map[string]int    // A map of counter
    mux sync.Mutex      // Mutual exclusion
}

func (c *SafeCounter) Inc(key string){
    // c.mux.Lock()
    c.v[key]++
    // c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int{
    // c.mux.Lock()
    // defer c.mux.Unlock()
    return c.v[key]
}

func main() {
    c := SafeCounter{v: make(map[string]int)}

    for i := 0; i < 1000; i++ {
        c.Inc("key1")
    }

    // time.Sleep(time.Millisecond * 2000)
    fmt.Println("Res = ", c.Value("key1"))
}
