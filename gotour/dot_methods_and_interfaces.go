package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64
type MyInt int

func (v Vertex) Sqrt() float64 {
	return math.Sqrt(v.X*v.Y + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Sqrt() float64 { return 1.2 }

func (v Vertex) Pow() float64 {
	return math.Pow(v.X, v.Y)
}

func (mi MyInt) Adder() int {
	fmt.Println("mi = ", mi)
	return int(mi) * 10
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Sqrt())
	fmt.Println(v.Pow())
	fmt.Println(Sqrt())

	mf, mi := MyFloat(1.1), MyInt(3)
	fmt.Println(mf, mi, mi.Adder())

	v.Scale(10)
	fmt.Println(v)
}
