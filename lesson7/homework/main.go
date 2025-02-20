package main

import (
	"fmt"
	"sort"
)

/*
练习题
1.写出下面代码的输出结果
func main() {
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
}

2.请使用内置的sort包对数组var a = [...]int{3,7,8,9,1}进行排序（附加题，自行查资料解答）
*/

func main() {
	// 1.解答
	answer()

	// 2.解答
	//var a = [...]int{3, 7, 8, 9, 1}
	//sortList(a)
}

func answer() {
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
	fmt.Println(len(a), cap(a))
}

func sortList(list [5]int) {
	// [...]int{3, 7, 8, 9, 1}
	sort.Ints(list[:])
	fmt.Println(list)
}
