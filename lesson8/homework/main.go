package main

import (
	"fmt"
	"strings"
)

/*
练习题
1.写一个程序，统计一个字符串中每个单词出现的次数。比如：“how do you do"中how=1 do=2 you=1。
2.观察下面代码，写出最终的打印结果。
func main() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}
*/

func main() {
	// 1.练习题解答
	//getTimes()

	// 2.练习题解答
	test()
}

/*
练习题1
*/
func getTimes() {
	// 1.写一个程序，统计一个字符串中每个单词出现的次数。比如：“how do you do"中how=1 do=2 you=1
	//var msg = "how do you do"
	var msg = "how old are you ? how are you ? how do you do"
	var list = strings.Split(msg, " ")
	var num = len(list)
	var data = make(map[string]int, num)

	for _, word1 := range list {
		_, ok := data[word1]
		if ok {
			data[word1]++
		} else {
			data[word1] = 1
		}
	}
	fmt.Println(data)
}

func test() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}
