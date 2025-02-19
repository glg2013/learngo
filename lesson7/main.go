package main

import "fmt"

/*
对应 Go语言基础07 切片
*/

func main() {
	// 声明切片类型
	//var a []string              //声明一个字符串切片
	//var b []int                 //声明一个整型切片并初始化
	//var c = []bool{false, true} //声明一个布尔切片并初始化
	//fmt.Println(a)              //[]
	//fmt.Println(b)              //[]
	//fmt.Println(c)              //[false true]
	//fmt.Println(a == nil)       //true
	//fmt.Println(b == nil)       //false
	//fmt.Println(c == nil)       //false
	//fmt.Println(c == d)          // 切片是引用类型，不支持直接比较，只能和nil比较

	//var list = [...]int{1, 2, 3, 4, 5}
	//var s = list[:]
	//fmt.Printf("%T\n", list)
	//fmt.Printf("%T\n", s)
	//fmt.Println(s)
	//var n = list[:]
	//var n = s[:]
	//
	// 基于同一个底层数组，如果切片改变了元素，那么数组也会改变
	//s[0] = 100
	//fmt.Println(list, s, n)
	//
	//s = append(s, 7, 8, 9)
	//fmt.Println(s, n)

	// 如果切片的起始位置变化，看看底层数组是否变化, 答案是底层数组不变
	// 基于同一个底层数组，某个切片的起始位置变化，其他切片并没有变化
	//s := list[1:4]
	////n := s[4:4] // 再切片的时候，如果起始索引和结束索引一样，那么切片的长度为0，左闭右合，导致取不到值
	//n := s[3:4] // 这里的关键点是：切片的索引是基于底层数组的，而不是基于切片本身的。(这是deepseek求的答案)
	//fmt.Println(s, len(s), cap(s))
	//fmt.Println(n, len(n), cap(n))
	//
	//// append() 函数可以为切片动态添加元素
	//var s1 []int
	//s1 = append(s1, 1, 2, 3)
	//
	//s2 := []int{5, 6, 7}
	//s1 = append(s1, s2...)
	//fmt.Println(s1)

	// copy()复制切片
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := make([]int, 5, 5)
	c[0] = 100
	fmt.Println(c) //[1 2 3 4 5]
	copy(c, a)     //使用copy()函数将切片a中的元素复制到切片c
	//fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(c) //[1 2 3 4 5]
	//c[0] = 1000
	//fmt.Println(a) //[1 2 3 4 5]
	//fmt.Println(c) //[1000 2 3 4 5]
}
