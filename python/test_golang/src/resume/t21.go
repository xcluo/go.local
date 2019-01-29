// 编译执行下面代码会出现什么?
package main

var(		// 问题3
    size := 1024		// 问题 1+2
    // max_size = size * 2   // 问题1 + 2
)

func main()  {
    println(size,max_size)
}

// 1, 定义变量同时显式初始化
// 2, 不能提供数据类型
// 3, 只能在函数内部使用
