package main

import "fmt"
import "github.com/iotaledger/giota"

func main() {
	fmt.Println("vim-go")

	api := giota.NewAPI("https://nodes.testnet.iota.org:443", nil)
	fmt.Println("AAA = ", api)

	a, b := api.GetNodeInfo()
	fmt.Println(a, b)

}
