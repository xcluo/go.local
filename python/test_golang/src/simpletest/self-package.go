package main

import "fmt"

func main() {
    println("test")

    type MyInt1 int
    type MyInt2 = int

	var i int = 1
    var myint1 MyInt1 = MyInt1(i)
    var myint2 MyInt2 = i

    println(myint1)
    println(myint2)

	var vm string
	fmt.Printf("%f" , 1.2)
	println(vm)
}


// func main()  {
//     type MyInt1 int
//     type MyInt2 = int
//     var i int =9
//     var i1 MyInt1 = i
//     var i2 MyInt2 = i
//     fmt.Println(i1,i2)
// }
