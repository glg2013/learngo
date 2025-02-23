package main

import (
	"fmt"
	"learngo/lesson18/homework/calc"
)

func main() {
	var a, b int
	a = 10
	b = 2
	fmt.Println("加法：", calc.Add(a, b))
	fmt.Println("减法：", calc.Sub(a, b))
	fmt.Println("乘法：", calc.Mul(a, b))
	fmt.Println("除法：", calc.Div(a, b))
}
