package main

import (
	"fmt"
	//"golang.org/x/tour/wc"
	"strings"
)

func main() {
	fmt.Println("abcd")
	s := "aaa bbb ccc"
	arr := strings.Split(s, " ")
	fmt.Println(arr)
}
