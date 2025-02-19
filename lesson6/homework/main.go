package main

import "fmt"

/*
练习题
1.求数组[1, 3, 5, 7, 8]所有元素的和
2.找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)
*/

func main() {
	// 1.解
	//var numArray = [5]int{1, 3, 5, 7, 8}
	//fmt.Printf("numArray:%T\n", numArray)
	//sum(numArray)

	// 2.解
	var numArray = [5]int{1, 3, 5, 7, 8}
	getIndex(numArray, 8)
}

func sum(a [5]int) {
	var total = 0
	for _, num := range a {
		total += num
	}
	fmt.Println("数组 numArray 所有元素总和是：", total)
}

func getIndex(numList [5]int, num int) {
	for i1, v1 := range numList {
		// 这样会出现重复，(0,3) (3,0)
		//for i2, v2 := range numList {
		//	if v1+v2 == num {
		//		fmt.Println("找到了一对和为", num, " 的数据,索引是:", "(", i1, ",", i2, ")")
		//	}
		//}

		// 下面这种就过滤掉了重复
		// 每次内部的 for 循环只遍历比外层更靠后的所有数据
		// i2 := i1 + 1 就把外层已经遍历过的，内层不再去遍历了
		for i2 := i1 + 1; i2 < len(numList); i2++ {
			if v1+numList[i2] == num {
				fmt.Println("找到了一对和为", num, " 的数据,索引是:", "(", i1, ",", i2, ")")
			}
		}
	}
}
