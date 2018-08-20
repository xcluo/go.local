package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("EEEE=", recover())
	}()

	if err := "A"; err != "B" {
		panic("AAAAA")
	}

	if err := "B"; err != "A" {
		panic("BBBB")
	}

	print("1111")
}
