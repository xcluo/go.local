package main

import "fmt"
import "time"

type error interface {
	Error() string
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

	err := &MyError{time.Now(), "it didn't work"}
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success!")
	}
}
