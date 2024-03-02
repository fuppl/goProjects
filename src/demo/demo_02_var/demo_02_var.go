package main

import (
	"fmt"
)

// 变量赋值
func main() {
	var a int = 1
	var b int = 2
	var c int = a + b
	fmt.Println("c = ", c) // c =  3

	var d, e = 2, 3
	var f int = d + e
	fmt.Println(f) // 5

	g := 4
	h := 5
	i := g + h
	fmt.Println(i) // 9

	// 使用括号可以批量赋值，如果赋值时不写，默认和上一个值相同
	const (
		n1 = 100
		n2
		n3
		n4
	)
	fmt.Println(n1, " ", n2, " ", n3, " ", n4) // 100   100   100   100

	// iota 是常用的计数器，用于枚举等，第一次出现表示 0，后面的值累加
	const (
		m1 = iota
		m2
		m3
		m4
	)
	fmt.Println(m1, m2, m3, m4) // 0 1 2 3

}
