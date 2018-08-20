package main

import (
	"fmt"
	"math"
)

// 剩余类集合下的乘法操作: x⋅y/p (mod N)
func MontMul() {
}

// Horner’s rule
//
// x⋅y =(y0⋅x⋅100)+(y1⋅x⋅101)+(y2⋅x⋅102)
//	   =(y0⋅x⋅0)+(y1⋅x⋅10)+(y2⋅x⋅100)
//	   =(y0⋅x)+10⋅(y1⋅x+10⋅(y2⋅x⋅+10⋅(0)))
// 最后一步，即是霍纳法则(Horner’s rule)
func Horner_s_Rule() {
}

func main() {
	// var (
	// 	a []uint32
	// 	b []uint32
	// )

	// a = []uint32{uint32(0x32b411c2), uint32(0x2214b3c2)}
	// b = []uint32{uint32(0x32b411c2), uint32(0x2214b3c2)}

	// for i := 0; i < 2; i++ {
	// 	fmt.Printf("A = %T, %d \n", a, a)
	// 	fmt.Printf("B = %T, %d \n", a, a)
	// 	fmt.Printf("C = %+v \n", a[i]*b[i])
	// }

	var (
		x, y, N, p uint32
		b, lb      int
	)

	x, y, N, b, lb = uint32(4294967295), uint32(4184967295), uint32(1394167273), 10, 10
	// x, y, N, b, lb = uint32(8), uint32(1), uint32(7), 10, 1
	p = uint32(math.Pow(float64(b), float64(lb)))
	mont_x, mont_y := (x*p)%N, (y*p)%N

	// Mont
	t := mont_x * mont_y
	m := (t / N) % N
	u := t + (m*p)/N
	mont := u - N

	// mont := (mont_x * mont_y * (1 / p)) % N
	fmt.Println(p, "(p)")
	fmt.Println(mont_x, "(mont_x)")
	fmt.Println(mont_y, "(mont_y)")
	fmt.Println(mont, "(Mount)")
	fmt.Println((x*y)%N, "(Trad)")
	// fmt.Println(((76 * 22) / 100) * 1)

	// r := uint32(0)
	// for i := 0; i <= lb-1; i++ {
	// 	fmt.Println("123")
	// 	r = (r + y_i*x + u*N) / b
	// 	fmt.Println(r)
	// }

	t = x * y
	mont = t - N*(t/N)
	fmt.Println(mont)
}
