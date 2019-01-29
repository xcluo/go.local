package main

import "fmt"

type Student struct{
    id int
    name string
}

func main(){
    var s_1 *Student = new(Student)
    s_1.id = 100
    s_1.name = "cat"
    var s_2 Student = Student{id:1,name:"tom"}
    fmt.Println(s_1,s_2)
}
