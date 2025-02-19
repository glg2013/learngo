package main

import "fmt"

/**
对应 Go语言基础06 数组
*/

func main() {
	// 1.定义
	//var a [3]int
	//var b [4]int
	//fmt.Println(a, b)

	//// 这样不可以赋值，长度一致，但是存储的数据类型不一致
	//var c [3]string = [3]string{"1", "2", "3"}
	//var d [3]int = [3]int{4, 5, 6}
	//d = c
	//fmt.Println(d)

	//// 这样可以赋值，长度和存储的数据类型一致
	//var c [3]int = [3]int{1, 2, 3}
	//var d [3]int = [3]int{4, 5, 6}
	//d = c
	//fmt.Println(d)

	// 2.数组的初始化
	// 方法一，初始化数组时可以使用初始化列表来设置数组元素的值。
	var testArray [3]int        // 数组会初始化为 int 类型的零值
	var numArray = [3]int{1, 2} // 使用指定的初始化值完成初始化    // 没有指定的值会被 0 初始化
	var cityArray = [3]string{"北京", "上海", "广州"}
	fmt.Println(testArray, numArray, cityArray)
	fmt.Println(len(testArray), len(cityArray), len(numArray))

	// 方法二，让编译器根据初始值的个数自行推断数组的长度
	//var testArray [3]int             // 数组会初始化为 int 类型的零值
	//var numArray = [...]int{1, 2, 3} // 使用指定的初始化值完成初始化
	//var cityArray = [...]string{"北京", "上海", "广州"}
	//fmt.Println(testArray, numArray, cityArray)
	//fmt.Println(len(testArray), len(cityArray), len(numArray))

	// 方法三，使用指定索引值的方式来初始化数组
	//a := [...]int{1: 1, 3: 3, 4: 5}
	//fmt.Println(a)
	//fmt.Printf("type of a:%T\n", a)

	// 3.数组的遍历
	//var a = [...]string{"北京", "上海", "广州"}
	//// 方法1：for循环遍历
	//for i := 0; i < len(a); i++ {
	//	fmt.Println(a[i])
	//}

	// 方法2：for range 方式
	//for index, value := range a {
	//	fmt.Println(index, value)
	//}

	// 4.多维数组
	//a := [3][2]string{
	//	{"北京", "上海"},
	//	{"广州", "深圳"},
	//	{"成都", "重庆"},
	//}
	//fmt.Println(a)
	//fmt.Println(a[2][1])

	// 5.多维数组
	//a := [3][2]string{
	//	{"北京", "上海"},
	//	{"广州", "深圳"},
	//	{"成都", "重庆"},
	//}
	//// 使用 2 次 for 循环
	//for _, v1 := range a {
	//	for _, city := range v1 {
	//		fmt.Printf("%s\t", city)
	//	}
	//	fmt.Println()
	//}

	////支持的写法
	//a := [...][2]string{
	//	{"北京", "上海"},
	//	{"广州", "深圳"},
	//	{"成都", "重庆"},
	//}
	//////不支持多维数组的内层使用...
	//b := [3][...]string{
	//	{"北京", "上海"},
	//	{"广州", "深圳"},
	//	{"成都", "重庆"},
	//}

	// 比较数组是否相等
	//a := [...][2]string{
	//	{"北京", "上海"},
	//	{"广州", "深圳"},
	//	{"成都", "重庆"},
	//}
	//b := [...][2]string{
	//	{"北京", "上海"},
	//	{"广州", "深圳"},
	//	{"成都", "重庆"},
	//}
	//if a == b {
	//	fmt.Println("a == b")
	//}

}
