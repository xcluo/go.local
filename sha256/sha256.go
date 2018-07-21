package main

import (
	"fmt"
	"io"
	"os"
)

// import "encoding/binary"

func main() {
	args := os.Args
	if args == nil || len(args) < 2 || len(args) > 2 {
		fmt.Println("You must assign one filename to hash!")
		return
	}

	fp, err := os.Open("t2.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer fp.Close()
	// // make a buffer to keep chunks that are read
	// buf := make([]byte, 512)
	// for {
	// 	// read a chunk
	// 	n, err := fi.Read(buf)
	// 	if err != nil && err != io.EOF {
	// 		panic(err)
	// 	}
	// 	if n == 0 {
	// 		break
	// 	}
	// 	fmt.Println("len=", n)
	// }
	CountStrings(fp)
	CountBinarys(fp)
}

func CountStrings(fp *os.File) {
	fmt.Println("Count Strings ------------")

	fmt.Println()
}

func CountBinarys(fp *os.File) {
	fmt.Println("Count Binarys ------------")
	data := make([]byte, 64)
	count := 0
	// data := []byte{}
	for i := 0; ; i++ {
		data = data[:cap(data)]
		n, err := fp.Read(data)
		if err != nil {
			if err == io.EOF {
				fmt.Println("\n\ni= ", i)
				break
			}
			fmt.Println(err)
			return
		}

		data = data[:n]
		// fmt.Printf("% X", data)
		fmt.Printf("n[%d]=%d, ", i, n)
		count += n
	}
	fmt.Println("Binary[byte] Count: ", count)
	fmt.Println()
}
