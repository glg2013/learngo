package main

import (
	"fmt"
)

func main() {

	/**
	练习题：
	1.编写代码分别定义一个整型、浮点型、布尔型、字符串型变量，使用fmt.Printf()搭配%T分别打印出上述变量的值和类型
	2.编写代码统计出字符串"hello沙河小王子"中汉字的数量
	*/

	//// 1.解
	//var a int = 100
	//var b float32 = 100.2
	//var c bool = true
	//var d string = "天天天蓝"
	//fmt.Printf("a = %v %T\nb = %v %T\nc = %v %T\nd = %v %T\n", a, a, b, b, c, c, d, d)

	// 2.解
	var message string = "hello沙河小王子"
	var cn int = 0
	for _, msg := range message {
		//fmt.Printf("%c %T %v\n", msg, msg, msg)

		// 这个是 deepseek 的答案
		//if unicode.Is(unicode.Han, msg) {
		//	cn++
		//}

		w := fmt.Sprintf("%c", msg)
		if len(w) >= 3 {
			cn++
		}

	}
	fmt.Println(cn)
}
