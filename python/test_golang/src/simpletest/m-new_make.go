package main

import "reflect"

func main(){
    var p *[]int = new([]int)
    var v []int = make([]int, 0)

    println(reflect.DeepEqual(*p, &v))

    println("a = ", *p)
    *p = append([]int{}, 1)
    println("a = ", *p)
    println("b = ", p)
    v = append(v, 1)
    println("c = ", v)

    // 指针
    i := 42
    m := &i
    println(*m)

    // new
    var q1 *int
    println("q1 = ", q1)
    q1 = new(int)
    println("q1 = ", q1)
    *q1 = 22
    println("q1 = ", q1)
    println("q1 = ", *q1)
    // println("q1 = ", *q1)
}
