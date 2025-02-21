package main

import (
	"fmt"
	"strings"
)

/*
对应 Go语言基础10 匿名函数和闭包
*/

func main() {

	//// 1.匿名函数
	//// 将匿名函数保存到函数
	//add := func(x, y int) {
	//	fmt.Println(x + y)
	//}
	//add(1, 2)
	//
	//// 自执行函数：匿名函数定义完成加上 () 直接执行
	//func(x, y int) {
	//	fmt.Println(x - y)
	//}(10, 5)

	//// 2.闭包使用
	//var f = adder()
	//fmt.Println(f(10))
	//fmt.Println(f(20))
	//fmt.Println(f(30))

	//// 3.闭包进阶1使用
	//f1 := adder2(20)
	//fmt.Println(f1(40))
	//fmt.Println(f1(50))

	//// 4.闭包进阶2使用
	//jpgFunc := makeSuffixFunc(".jpg")
	//txtFunc := makeSuffixFunc(".txt")
	//fmt.Println(jpgFunc("test")) //test.jpg
	//fmt.Println(txtFunc("test")) //test.txt

	//// 5.闭包进阶3
	//f1, f2 := calc(10)
	//fmt.Println(f1(1), f2(2)) //11 9
	//fmt.Println(f1(3), f2(4)) //12 8
	//fmt.Println(f1(5), f2(6)) //13 7

	//// 6.defer 执行的时机和顺序
	//fmt.Println("start")
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//defer fmt.Println(3)
	//fmt.Println("end")

	//// 7.panic / recover
	//// 7.1 没有使用 recover 的会报错退出
	//funcA()
	//funcB()
	//funcC()

	// 7.2 使用 panic 来处理
	funcA()
	funcB2()
	funcC()
}

// 2.闭包函数
func adder() func(int) int {
	var x int
	fmt.Println("看看闭包引用的外部变量的值变化1: ", x)
	return func(y int) int {
		fmt.Println("看看闭包引用的外部变量的值变化2: ", x)
		x += y
		fmt.Println("看看闭包引用的外部变量的值变化3: ", x)
		return x
	}
}

// 3.闭包进阶示例1
func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

// 4.闭包进阶2
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 5.闭包进阶3
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

// 7.panic / recover
func funcA() {
	fmt.Println("func A")
}

func funcB() {
	panic("panic in B")
}

func funcB2() {
	// 必须在 defer 中定义才生效
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("有错出现在b2函数中。。", err)
		}
	}()
	// 这里触发一下，来测试效果
	panic("panic in B2")
}

func funcC() {
	fmt.Println("func C")
}
