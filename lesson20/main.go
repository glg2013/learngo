package main

import (
	"fmt"
	"time"
)

/*
对应 Go语言基础20 标准库time介绍
2025年2月23日21:19:30
这一课主要是学习time包的使用
*/

/*
练习题:
1.获取当前时间，格式化输出为2017/06/19 20:30:05`格式。
2.编写程序统计一段代码的执行耗时时间，单位精确到微秒。
*/

func main() {
	// 检查练习题1
	//getNowTime()

	// 检查练习题2
	getTime()
}

// 练习题1
func getNowTime() {
	fmt.Println("getNowTime", time.Now().Format("2006/01/02 15:04:05"))
}

// 练习题2
func getTime() {
	start := time.Now().UnixMilli()
	// 执行一段 for 循环
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
	end := time.Now().UnixMilli()
	//fmt.Println("for 循环的耗时为：", (end-start)/1000)
	fmt.Println("for 循环的耗时毫秒数为：", end-start)
}
