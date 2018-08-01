package main

import (
	"fmt"
	"image"
)

type Image struct{}

func main() {
	m := image.NewRGBA(image.Rect(100, 100, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())

	n := Image{}
}
