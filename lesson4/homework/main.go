package main

import "fmt"

func main() {

	/**
	练习题
	1.有一堆数字，如果除了一个数字以外，其他数字都出现了两次，那么如何找到出现一次的数字？
	*/

	// 这个是课后参考别人的位亦或的解题方式，不是很明白
	/**
	任何数和 00 做异或运算，结果仍然是原来的数，即 a ^ 0=aa⊕0=a。
	任何数和其自身做异或运算，结果是 00，即 a ^ a=0a⊕a=0。
	异或运算满足交换律和结合律，即 a ^ b ^ a=b ^ a ^ a=b ^ (a ^ a)=b ^ 0=ba⊕b⊕a=b⊕a⊕a=b⊕(a⊕a)=b⊕0=b。
	*/
	//var numberArray = [5]int{1, 2, 1, 3, 3}
	//var num int = 0
	//for i := 0; i < len(numberArray); i++ {
	//	num ^= numberArray[i]
	//	fmt.Println(num)
	//}
	//fmt.Printf("只出现一次的数是:%d", num)

	//nums := [...]int{1, 1, 2, 2, 3, 3, 4}
	//res := 0
	//for _, v := range nums {
	//	res ^= v
	//	fmt.Println(res)
	//}
	//fmt.Println(res)

	var numberArray = [5]int{1, 2, 1, 3, 3}
	var list = make([]int, len(numberArray))
	for _, i := range numberArray {
		//fmt.Println(i)
		list[i]++
	}
	//fmt.Println(list)
	for _, v := range list {
		if v == 1 {
			fmt.Println(v)
		}
	}

}
