package main

import "fmt"

/*
对应 Go语言基础12自定义类型和类型别名
*/

// NewInt 自定义类型定义
type NewInt int

// MyInt 类型别名
type MyInt = int

func main() {
	var a NewInt
	var b MyInt

	fmt.Printf("type of a: %T\n", a)
	fmt.Printf("type of b: %T\n", b)

	//type of a: main.NewInt
	//type of b: int
}
