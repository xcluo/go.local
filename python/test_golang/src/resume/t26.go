//编译执行下面代码会出现什么?
package main

import "fmt"

type T1 struct {
}
func (t T1) m1(){
    fmt.Println("T1.m1")
}

type T2 = T1		// 完全相等了
type MyStruct struct {
    T1
    T2
}
func main() {
    my:=MyStruct{}

    my.m1()
    // my.T2.m1()	// 改为这个
}


//结果不限于方法，字段也也一样；也不限于type alias，type defintion也是一样的，只要有重复的方法、字段，就会有这种提示，因为不知道该选择哪个。
