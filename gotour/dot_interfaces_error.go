package main

import (
	"fmt"
	"time"
)

type error interface {
	Error() string
}

type MyError struct {
	When time.Time
	What string
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Error!!!!!!!")
}

// func (e *MyError) Error() string {
// }

func Sqrt(x ErrNegativeSqrt) (float64, error) {
	fmt.Println("x = ", x)
	if x < 0 {
		return 0, error(&MyError{time.Now(), "Less 0"})
	} else {
		return 0, nil
	}
}

func main() {
	fmt.Println(Sqrt(2))
	// fmt.Println(Sqrt(-2))
}
