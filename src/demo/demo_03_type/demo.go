package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	a := 1
	b := 2
	fmt.Println(a + b)

	s1 := "abc"
	var s2 string = "def"
	fmt.Println(s1, s2)

	// 多行字符串使用反引号，里面的转义字符均失效
	s3 := `line 1
	line2
	line3
	/n /b /r
	`
	fmt.Println(s3)

	var b1 bool = strings.Contains(s1, "ab")
	fmt.Print(b1)

	// go 的字符分两种，一种是uint8，一种是rune。前者是 ascii 码，后者对应 utf-8编码。
	s4 := "测试中文字符"
	for _, r := range s4 { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()

	changeString()

	sqrtDemo()
}

func changeString() {
	s1 := "hello"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))

	s2 := "博客"
	runeS2 := []rune(s2)
	runeS2[0] = '狗'
	fmt.Println(string(runeS2))
}

func sqrtDemo() {
	var a, b = 3, 4
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	var c int = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
