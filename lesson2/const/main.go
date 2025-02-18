package main

import "fmt"

/**
对应 Go语言基础02 变量与常量
常量
*/

// 常量
const pi = 3.14
const e = 2.7

const (
	x = 100
	y
)

const (
	n1 = iota
	n2
	n3
	n4
	n5
)

const (
	m1 = iota
	_
	m2
	m3
)

const (
	u1 = iota
	u2 = iota
	u3 = 100
	u4
	u5 = iota
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
)

const (
	a99, b99 = iota + 1, iota + 2
	c99, d99
	e99, f99
)

func main() {
	// 打印常量
	fmt.Println(pi, e, x, y)

	fmt.Println(n1, n2, n3, n4, n5)

	fmt.Println(m1, m2, m3)

	fmt.Println(u1, u2, u3, u4, u5)

	fmt.Println(a99, b99, c99, d99, e99, f99)
}
