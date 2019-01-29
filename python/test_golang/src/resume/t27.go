// 编译执行下面代码会出现什么?
package main

import (
    "errors"
    "fmt"
)

var ErrDidNotWork = errors.New("did not work")

func DoTheThing(reallyDoIt bool) (err error) {
	var bb error
    if reallyDoIt {
        result, err := tryTheThing()		// err 是局部的
		fmt.Println(result, "|", err)
        if err != nil || result != "it worked" {
			fmt.Println("In err = ", err)
            // err = ErrDidNotWork
            bb = ErrDidNotWork
        }
    }
	// fmt.Println(result, "|", err)
    return bb
}

func tryTheThing() (string,error)  {
    return "",ErrDidNotWork
}

func main() {
    fmt.Println(DoTheThing(true))
    // fmt.Println(DoTheThing(false))
}


// if 里面的变量 err 是有作用域的，所以输出一直是两个 nil，并没有输出 errors
