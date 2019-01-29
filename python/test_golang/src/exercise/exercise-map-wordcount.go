/*
实现 WordCount。它应当返回一个映射，其中包含字符串 s 中每个“单词”的个数。函数 wc.Test 会对此函数执行一系列测试用例，并输出成功还是失败。

你会发现 strings.Fields 很有帮助。
*/
package main

import (
	"golang.org/x/tour/wc"
	"strings"
	// "fmt"
)

func WordCount(s string) map[string]int {
	var rest map[string]int = map[string]int{}
	// fmt.Printf("%v, %d \n", strings.Split(s, ""), len(strings.Fields(s)))
	for _, v := range strings.Fields(s) {
		// println(v)
		if _, exist := rest[v]; exist {
			rest[v] += 1
		} else {
			rest[v] = 1
		}
	}
	return rest //map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}

