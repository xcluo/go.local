// test defer
package main

import "fmt"

func deferRet(x,y int) (z int){
  defer func() {z += 100}()
  z = x + y
  return z + 50 // 执行顺序 z = z+50 -> (call defer)z = z+100 -> ret  
}

func f1() (result int) {
    defer func() {
        result++
    }()
    return 0
}

func f2() (r int) {
    t := 5
    defer func() {
        t = t + 5
        fmt.Println("f2")
    }()
    return t
}

func f3() (r int) {
    defer func(r int) {
        r = r + 5
    }(r)
    return 1
}

func main() {
 fmt.Println(f1())
 fmt.Println(f2())
 fmt.Println(f3())
}
