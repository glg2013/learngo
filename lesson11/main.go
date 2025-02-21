package main

import "fmt"

/*
对应 Go语言基础11 指针
2025年2月21日15:50:52
*/
func main() {
	a := 10
	b := &a
	fmt.Println(a, b, &*b, *b)
	fmt.Printf("%T\n", b)
}
