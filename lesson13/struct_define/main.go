package main

import "fmt"

/*
对应 Go语言基础 13结构体定义与实例化
2025年2月21日17:44:49
*/

// 1.结构体的定义
type person struct {
	name string
	city string
	age  int
}

func main() {
	//var a, b int
	//fmt.Println(a, b)

	//// 1.结构体的实例化
	//var p person
	//p.name = "小李"
	//p.city = "上海"
	//p.age = 18
	//fmt.Printf("type of p: %T\n %v\n", p, p)

	// 2.匿名结构体
	// 标准：这种是标准的匿名构造体,特别注意格式，写在同一行的话是 分号 ; 分割
	//var user1 struct{name string; city string; age int}
	// 简化：这种是同类型简化，写在同一行得话要有 分号 ; 分割
	//var user2 struct{name, city string; age int}
	// 简化2：这种是不写在同一行，就可以省略分割的 分号，换行就可以了
	var user3 struct {
		name, city string
		age        int
	}
	user3.name = "小李"
	user3.city = "北京"
	user3.age = 18
	fmt.Printf("user3: %#v\n", user3)
}
