package main

import (
	"fmt"
	"math"
)

/**
对应 Go语言基础03 基本数据类型
数据类型
*/

func main() {

	// 十进制
	var n int = 10
	fmt.Printf("%b\n", n)
	fmt.Printf("%d\n", n)

	// 八进制
	var m int = 075
	fmt.Printf("%d\n", m)
	fmt.Printf("%o\n", m)

	// 十六进制
	var f int = 0xff
	fmt.Println(f)
	fmt.Printf("%x\n", f)

	// uint8
	var age uint8 = 255
	fmt.Println(age)

	// 浮点数
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	// 布尔值
	var a bool
	fmt.Println(a)
	a = true
	fmt.Println(a)

	// 字符串
	var s1 string = "hello beijing"
	var s2 string = "你好"
	fmt.Println(s1, "\n", s2)

	// 字符串转义
	fmt.Println("c:\\code\\go.exe")

	// 多行字符串
	var s3 = `
这里的内容是不会被转义的，全部都是原样输出的\n\t
    `
	fmt.Println(s3)
}
