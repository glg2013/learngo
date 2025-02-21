package main

import "fmt"

/*
对应 Go语言基础09 函数 - defer 语句
defer 的几个重要知识点
1. 先被defer的语句最后被执行，最后被defer的语句，最先被执行
2. defer注册要延迟执行的函数时该函数所有的参数都需要确定其值
（defer 注册的函数，
		如遇到 注册的是匿名立即执行函数，并且传递了参数进去，那么它会当时立即捕获那个值
		如遇到 注册的函数中，有将函数作为参数传递，那么那个被当成参数的函数也会捕获当时的关联变量
		这些 defer 都不会等到执行的时候才去捕获变量
其他的情况都是 等执行 defer 的时候才去捕获相关变量
）

*/

//func main() {
//	fmt.Println("start")
//	//defer fmt.Println(1)
//	//defer fmt.Println(2)
//	//defer fmt.Println(3)
//	defer test()
//	fmt.Println("end")
//}
//
//func test() {
//	fmt.Println("我们不一样-----")
//}

/*
*
f1() 返回 5
1.首先局部变量 x 声明初始话为 5
2.然后 先去执行 return 中的 返回值赋值，也就是 return x，这里是获取 x 当前值的副本的值 5，给到返回值
3.执行 defer 注册的匿名立即执行函数，虽然这里 x++，但是，因为在 return 赋值的时候，int 类型是值的拷贝，所以这里修改 x 也不会影响
4.最终执行 return 的时候，返回值就是 5
*/
func f1() int {
	x := 5
	defer func() {
		x++
		fmt.Println("此时x值是： ", x)
	}()
	fmt.Println("打印一下x的值： ", x)
	return x
}

/*
f2() 返回6
1.这里定义了一个命名返回值 x 类型是 int，并初始化为 x = 0
2.然后去执行了 return 的赋值操作，此时 x = 5
3.最后在函数 return 之前执行 defer 的立即执行匿名函数，这里将 命名返回值修改了
4.执行最终的 return 的时候，匿名返回值 x 的值被更新为了6
*/
func f2() (x int) {
	defer func() {
		fmt.Println("f2 虽然我是延迟执行，但是看看我捕获到的 x 的值是多少呢？ ", x)
		x++
	}()
	return 5
}

/*
f3() 返回5
1.定义并初始化了一个明明返回值 y 此时 y = 0
2.return 这里进行返回值赋值，将当前局部变量 x 的值的副本赋值给了 y
3.执行 defer 注册的匿名函数，此时虽然 x++ 改变了 x 的值，但是最终返回值在之前已经获取到了副本，所以这里就没影响了
4.执行最终的 return ，命名返回值 y 获取的就是最初的 x:= 5 值的那个副本
*/
func f3() (y int) {
	x := 5
	defer func() {
		fmt.Println("f3 虽然我是延迟执行，但是看看我捕获到的 x 的值是多少呢？ ", x)
		x++
	}()
	fmt.Println("看看是我先执行，还是defer先执行的")
	return x + 1
}

/*
f4() 返回5
1.定义并初始化了一个命名返回值 x 此时 x = 0
2.defer 注册了一个匿名函数，并立即传入了当前的 x 值（此时 x 的值为 0）
3.执行 return 返回值赋值操作， 将 return 语句这里的 5 赋值给了 命名返回值变量的 x 此时 x = 5
4.执行 defer 立即执行匿名函数，2的时候已经传递了参数，而且 x++ 改的是匿名函数中的局部变量，不是命名返回值的那个 x
最终 return 返回的是 5
*/
func f4() (x int) {
	defer func(x int) {
		fmt.Println("f4 虽然我是延迟执行，但是看看我捕获到的 x 的值是多少呢？ ", x)
		x++
		fmt.Println("f4 虽然我是延迟执行，但是看看我捕获到的 x 的值是多少呢---？ ", x)
	}(x)
	fmt.Println("f4 。。。。。。")
	return 5
}

func main() {
	// 这里练习 defer经典案例
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())

	// 这里是练习 defer面试题
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
