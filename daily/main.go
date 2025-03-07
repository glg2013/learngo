package main

import "fmt"

type userInfo struct {
	name string
	age  int
}

func main() {
	//fmt.Println("hello world")
	//test(2, 3)
	//
	//var n = 1
	//var p *int = &n
	//fmt.Println(p)
	//
	//var m = new(int)
	//fmt.Println(*m)
	//
	//var usr = userInfo{
	//	name: "hello",
	//}
	//fmt.Println("usr.name-> ", usr.name)
	//// 看看可以修改不
	//usr.name = "no hello"
	//fmt.Println("usr.name2-> ", usr.name)

	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x)
		}
	}
}

func test(x ...int) {
	fmt.Println("如果没有传递参数，可以打印正常吗？")
}
