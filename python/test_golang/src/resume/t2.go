package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for a, stu := range stus {
		fmt.Println(a)
		fmt.Println(stu)
		fmt.Println(stu.Name)
		m[stu.Name] = &stu
		// fmt.Printf("%p\n", stu) 
		fmt.Println("\n")
	}

	// for i := 0; i < len(stus); i++ {
	// 	fmt.Println(stus[i].Name)
	// }
    println(m)
    println(&m)
}

func main() {
	pase_student()
}
