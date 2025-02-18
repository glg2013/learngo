package main

/**
  对应 Go语言基础02 变量与常量
  变量
*/

// 0.全局变量
var name0 = "笔袋"
var age0 = 100

func main() {

	// 0.打印全局变量
	println("0.", name0, age0)

	// 1.标准声明格式
	var name string
	var age int
	var isOk bool
	println("1.", name, age, isOk)

	// 2.批量声明变量
	var (
		a string
		b int
		c bool
		d float32
	)
	println("2.", a, b, c, d)

	// 3.声明变量并指定初始值
	var name1 string = "小王子"
	var age1 int = 18
	println("3.", name1, age1)

	// 4.一次性声明并初始化多个变量
	var name2, age2 = "大风天", 28
	println("4.", name2, age2)

	// 5.类型推导
	var name3 = "十字军"
	var age3 = 99
	println("5.", name3, age3)

	// 6.简短声明
	m := 10
	println("6.", m)
}
