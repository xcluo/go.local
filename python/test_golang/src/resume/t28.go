// 编译执行下面代码会出现什么?
package main

func test(n int) []func() {
	var funcs []func()

	for i:=0; i<n; i++ {
		x := i
		funcs = append(funcs, func(){
			// println(&i, i)
			println(&x, x)
		})
	}
	return funcs
}

func main() {
	funcs := test(5)

	for _, f := range funcs {
		f()
	}
}

// 闭包延迟求值

// 0x4200000a0 5
// 0x4200000a0 5
// 0x4200000a0 5
// 0x4200000a0 5
// 0x4200000a0 5

// 每次输出一样：for循环会复用局部变量i，每一次放入匿名函数的应用都是想一个变量
