/*
	接口，类型变换，类型断言
*/
package main

import "fmt"

type NodeManager interface {
	startNode() error
}

type NodeManagerStruct struct{}

func (n *NodeManagerStruct) startNode() error {
	fmt.Println("func")
	return nil
}

func main() {
	fmt.Println("Way-1")
	var m NodeManager
	m = &NodeManagerStruct{} // 必须赋予地址给接口
	m.startNode()

	fmt.Println("Way-2")
	// m := &NodeManagerStruct{}.(NodeManager)		# 类型断言不支持：这里是显式转化
	n := NodeManager(&NodeManagerStruct{})
	n.startNode()
}
