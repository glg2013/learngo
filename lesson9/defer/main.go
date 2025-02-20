package main

import "fmt"

/*
对应 Go语言基础09 函数 - defer 语句
*/

//func main() {
//	fmt.Println("start")
//	defer fmt.Println(1)
//	defer fmt.Println(2)
//	defer fmt.Println(3)
//	fmt.Println("end")
//}

func f1() int {
	x := 5
	defer func() {
		x++
		fmt.Println("此时x值是： ", x)
	}()
	fmt.Println("打印一下x的值： ", x)
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
func main() {
	fmt.Println(f1())
	//fmt.Println(f2())
	//fmt.Println(f3())
	//fmt.Println(f4())
}
