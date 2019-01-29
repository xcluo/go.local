// 编译执行下面代码会出现什么?
package main

func test(n int) (func(), func()) {
    return func() {
        println(n)
        n += 10
    }, func() {
        println(n)
    }       // 闭包，这里复用一个变量
}

func main() {
    a, b := test(100)
    a()
    b()
}

// 输出： 100, 110
// 闭包引用相同变量
