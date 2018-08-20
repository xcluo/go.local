package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

var m map[string]Vertex

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func main() {
	// fmt.Println("Start")
	// var p *int
	// fmt.Println(p)

	// i := 40
	// p = &i
	// fmt.Println(p)
	// fmt.Println(&p)
	// fmt.Println(*p)

	// v := Vertex{2, 3}
	// fmt.Printf("v Type = %T \n", *p)
	// fmt.Println(v, v.X, v.Y, v.X*v.Y)
	// fmt.Println(v1, p, v2, v3)

	// var a []int = []int{1, 2, 3, 4}
	// b := [6]int{2, 3, 4, 5, 6, 7}
	// fmt.Println(a, b[:0])
	// var c []int = []int{}
	// fmt.Println(c, c == nil, len(c), cap(c))

	// fmt.Println(len(b), cap(b))
	// fmt.Println(len(b[4:4]), cap(b[4:4]))

	// e := [][]int{
	// 	[]int{1, 2, 3},
	// 	[]int{4, 5, 6},
	// 	[]int{7, 8, 9},
	// 	[]int{},
	// }
	// e = append(e, []int{1, 2})
	// fmt.Println(e, len(e), cap(e))
	// for i := 0; i < len(e); i++ {
	// 	fmt.Println("i = ", cap(e))
	// 	fmt.Println(e[i])
	// 	for j := 0; j < len(e[i]); j++ {
	// 		fmt.Println(e[i][j])
	// 	}
	// }
	// for i, v := range e {
	// 	fmt.Printf("2**%d = %d\n", i, v)
	// }

	m = make(map[string]Vertex)
	m["aaa"] = Vertex{1, 2}
	m["bbb"] = Vertex{3, 4}
	fmt.Println(m)

	//var m1 map[int]string
	m1 := make(map[int]string)
	m1[1] = "aaaa"
	m1[2] = "bbbb"
	fmt.Println(m1)

	elem, ok := m1[2]
	fmt.Println(elem)
	fmt.Println(ok)
}
