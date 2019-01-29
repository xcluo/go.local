// 下面函数有什么问题？
package main

const cl  = 100
var bl    = 123

func main()  {
    println(&bl,bl)

    println(&cl,cl)		// 这里问题
    // println(cl)

	println("aaaa")
}


// 常量在预处理阶段直接展开到代码中，作为指令数据直接使用，没有像变量一个单独划分数据内存
