// 是否可以编译通过？如果通过，输出什么？
package main

import "fmt"

func main() {
    s1 := []int{1, 2, 3}
    s2 := []int{4, 5}
    // s3 := []int{7, 8}
    fmt.Printf("len=%d, cap=%d\n", len(s1), cap(s1))
    s1 = append(s1, s2...)
    fmt.Printf("len=%d, cap=%d\n", len(s1), cap(s1))
    s1 = append(s1, 1, 2, 3)
    fmt.Printf("len=%d, cap=%d\n", len(s1), cap(s1))
    s1 = append(s1, s1...)
    fmt.Printf("len=%d, cap=%d\n", len(s1), cap(s1))
    s1 = append(s1, 1)
    fmt.Printf("len=%d, cap=%d\n", len(s1), cap(s1))
    s1 = append(s1, 1, 2, 3, 4)
    fmt.Printf("len=%d, cap=%d\n", len(s1), cap(s1))
    s1 = append(s1, 1, 2, 3)
    fmt.Printf("len=%d, cap=%d\n", len(s1), cap(s1))
    s1 = append(s1, 1, 2)
    fmt.Printf("len=%d, cap=%d\n", len(s1), cap(s1))
    fmt.Println(s1)
}

// 切片容量不够的时候，之前的cap 翻一倍
