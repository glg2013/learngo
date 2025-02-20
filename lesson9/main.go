package main

import (
	"errors"
	"fmt"
)

/*
对应 Go语言基础09 函数
*/

// 6.全局变量
var num int64 = 10

func main() {
	//// 3.函数的调用
	//sayHello()
	//ret := initSum(10, 20)
	//fmt.Println(ret)

	//// 可变参数调用
	//ret := initSum3(10, 20, 30)
	//fmt.Println(ret)

	//// 测试一个突然想到的前面课程提到的 for range 可以遍历一个整数了，利用这个来计算一个数之内的所有数的和
	//total := getSumNum(100)
	//fmt.Println(total)

	//// 多返回值的函数调用
	//v1, v2 := calc(10, 20)
	//fmt.Println(v1, v2)

	//// 返回值补充的调用
	//v := someFunc("")
	//fmt.Println(v)

	//// 全局变量的访问
	//testGlobalVar()

	//// 局部变量在函数外是无法访问的
	//testLocalVar()
	//fmt.Println(x)	// 会报错的 .\main.go:39:14: undefined: x

	//// 8.函数类型变量
	//type calculation func(int, int) int
	//var c calculation               // 声明一个calculation类型的变量c
	//c = add                         // 把add赋值给c
	//fmt.Printf("type of c:%T\n", c) // type of c:main.calculation
	//fmt.Println(c(1, 2))            // 像调用add一样调用c
	//
	//f := add                        // 将函数add赋值给变量f
	//fmt.Printf("type of f:%T\n", f) // type of f:func(int, int) int
	//fmt.Println(f(10, 20))          // 像调用add一样调用f

	//// 9.高阶函数
	//// 9.1 函数作为参数
	//ret2 := calc3(10, 20, add)
	//fmt.Println(ret2)

	//// 9.2 函数作为返回值
	//op, err := do("")
	//fmt.Printf("op:%T\n,err:%T\n, err%v\n", op, err, err)
	//if err != nil {
	//	fmt.Println("桥豆麻袋。。。", err.Error())
	//} else {
	//	result := op(10, 15)
	//	fmt.Println(result)
	//}

	//// 10 匿名函数
	//// 10.1将匿名函数保存到变量
	////opAdd := func(x, y int) {
	////	fmt.Println("匿名函数被调用啦  ", x+y)
	////}
	////opAdd(10, 15)
	//
	//// 10.2 自执行函数，匿名函数定义完成加()直接执行
	//func(x, y int) {
	//	fmt.Println(x + y)
	//}(25, 30)

	// 七米老师教程里提到的 闭包其实并不复杂，只要牢记闭包=函数+引用环境
	// deepseek 里获取的解释：闭包使得函数可以“记住”其创建时的环境，即使这个环境已经超出了其原始作用域。
	// 个人理解就是它会一直记得创建时的环境（那些变量，即使那些变量值发生了改变，比如这里的 x，它一直被记得，所以多次调用，会发现这个值一直在变）
	var f = adder()
	//fmt.Printf("%T\n", f)
	fmt.Println(f(10))
	fmt.Println(f(5))

}

// 1.基本定义
func initSum(x int, y int) int {
	return x + y
}

// 省略类型
func initSum2(x, y int) int {
	return x + y
}

// 2.函数参数和返回值都是可选的
func sayHello() {
	fmt.Println("hello world")
}

// 4.可变参数
func initSum3(x ...int) int {
	fmt.Println(x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

// 自测的，统计给的一个数字包含的所有数字的总和
func getSumNum(maxNum int) int {
	var total = 0
	for num := range maxNum {
		total += num + 1
	}
	return total
}

// 5.多个返回值
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// 5.1 返回值命名(返回值命名后，return 那边可以省略)
func calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

// 5.2 返回值补充
// 当我们的一个函数返回值类型为slice时，nil可以看做是一个有效的slice，没必要显示返回一个长度为0的切片。
func someFunc(x string) []int {
	if x == "" {
		fmt.Println("找到你啦，小不点")
		return nil
	}
	fmt.Println("没用的，不会执行")
	return []int{}
}

// 6.全局变量的使用
func testGlobalVar() {
	fmt.Printf("num=%d\n", num) // 函数中可以访问全局变量 num
}

// 7.局部变量的使用
func testLocalVar() {
	// 定义一个函数局部变量 x, 仅在该函数内部生效
	var x int64 = 100
	fmt.Printf("x=%d\n", x)
}

// 8.函数类型变量需要用的2个函数
func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}

// 9.高阶函数
// 9.1 函数可作为参数
func calc3(x, y int, op func(int, int) int) int {
	return op(x, y)
}

// 9.2 函数可以作为返回值
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符！")
		return nil, err
	}
}

// 11.闭包函数
func adder() func(int) int {
	var x int
	fmt.Println("0每次调用的时候，看看x的值是多少: ", x)
	return func(y int) int {
		fmt.Println("1每次调用的时候，看看x的值是多少: ", x)
		x += y
		fmt.Println("2每次调用的时候，看看x的值是多少: ", x)
		return x
	}
}
