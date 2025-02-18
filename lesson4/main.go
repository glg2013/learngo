package main

import "fmt"

/**
对应 Go语言基础04 运算符
*/

func main() {
	// 1. 算术运算符
	// + - * / % ++ --
	a := 5
	b := 2
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	a++
	fmt.Println(a)
	a--
	fmt.Println(a)

	// 2. 关系运算符
	fmt.Println(10 > 2)
	fmt.Println(10 != 2)
	fmt.Println(4 < 5)

	// 3. 逻辑运算符
	fmt.Println(4 > 5 && 10 > 2)
	fmt.Println(!(10 > 5))
	fmt.Println(1 > 5 || 10 > 2)

	// 4. 位运算符
	n := 1              // 0001
	m := 5              // 0101
	fmt.Println(n & m)  // 0001
	fmt.Println(n | m)  // 0101
	fmt.Println(n ^ m)  // 0100
	fmt.Println(n << m) // 00010000
	fmt.Println(n >> m) // 00000001

	// 5. 赋值运算符
	var c int
	c = 10
	c += 10
	fmt.Println(c)
}
