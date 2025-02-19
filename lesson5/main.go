package main

import "fmt"

/**
对应 Go语言基础05 流程控制
*/

func main() {

	// 最常见的形式
	//score := 100
	//if score >= 90 {
	//	fmt.Println("A")
	//} else if score >= 80 {
	//	fmt.Println("B")
	//} else if score >= 70 {
	//	fmt.Println("C")
	//} else {
	//	fmt.Println("D")
	//}

	// 特殊形式
	//if score := 90; score >= 85 {
	//	fmt.Println("优秀")
	//} else if score >= 75 {
	//	fmt.Println("良好")
	//} else {
	//	fmt.Println("需努力")
	//}

	// for 循环
	// 1.常见的 for 循环形式
	//for i := 1; i < 11; i++ {
	//	fmt.Println(i)
	//}

	// 2.省略初始语句
	//i := 0
	//for ; i < 10; i++ {
	//	fmt.Println(i)
	//}

	// 3.省略初始化语句和结束语句
	//i := 0
	//for i < 10 {
	//	fmt.Println(i)
	//	i++ 	// *这里这里的终止循环条件
	//}

	// 4.无限循环，和 while 很像
	//v := 100
	//for {
	//	fmt.Println(v)
	//	v++           // *注意这里的改变条件
	//	if v >= 105 { // *要有能终止循环的判断，否则就是死循环了
	//		break
	//	}
	//}
	//fmt.Println(v)

	// for range(键值循环)
	// Go1.22版本开始支持 for range 整数
	//for i := range 5 {
	//	fmt.Println(i)
	//}

	//for range 2 {
	//	fmt.Println("不放弃，坚持下去！")
	//}

	// switch 循环
	//finger := 0
	//switch finger {
	//case 1:
	//	fmt.Println("小了")
	//case 2:
	//	fmt.Println("接近了")
	//case 3:
	//	fmt.Println("对了")
	//default:
	//	fmt.Print("那我告诉你答案吧，是", finger)
	//}

	//var num = 5
	//switch num {
	//case 1, 3, 5, 7, 9:
	//	fmt.Println("这是个奇数")
	//case 2, 4, 6, 8:
	//	fmt.Println("这是个偶数")
	//default:
	//	fmt.Println("再试着猜一下呢？")
	//}

	//s := "a"
	//switch {
	//case s == "a":
	//	fmt.Println("a")
	//	fallthrough
	//case s == "b":
	//	fmt.Println("b")
	//case s == "c":
	//	fmt.Println("c")
	//default:
	//	fmt.Println("...")
	//}

	//var breakFlag bool
	//for i := 0; i < 10; i++ {
	//	for j := 0; j < 10; j++ {
	//		if j == 2 {
	//			// 设置退出标签
	//			breakFlag = true
	//			break
	//		}
	//		fmt.Printf("%v-%v\n", i, j)
	//	}
	//	// 外层for循环判断
	//	if breakFlag {
	//		break
	//	}
	//}

	//gotoDemo2()

	//breakDemo1()

	continueDemo()

}

func gotoDemo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")
}

func breakDemo1() {
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	fmt.Println("...")
}

func continueDemo() {
forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
}
