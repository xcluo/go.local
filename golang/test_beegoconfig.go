package main

import "fmt"
import "github.com/astaxie/beego/config"

func main() {
	iniconf, err := config.NewConfig("ini", "testini.ini")
	if err != nil {
		fmt.Println("AAA = ", err)
	} else {
		fmt.Println("BBB = ", iniconf)
	}

	fmt.Printf("\nCCC = %v %T, \n", iniconf.String("Bb"), iniconf.String("Bb"))

	fmt.Println("Set = ", iniconf.Set("bb", "cccc"))
	fmt.Println("Set = ", iniconf.Set("cc", "c2v1"))
	fmt.Println("DDD = ", iniconf.SaveConfigFile("./testini.ini"))
}
