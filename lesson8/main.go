package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/*
对应 Go语言基础08 map
*/

func main() {

	// 一、map 的基本使用
	// 使用 make() 方法
	//scoreMap := make(map[string]int, 8)
	//scoreMap["张三"] = 90
	//scoreMap["小李"] = 100
	//fmt.Println(scoreMap)
	//fmt.Println(scoreMap["小李"])
	//fmt.Printf("%T\n", scoreMap)

	// 声明的时候填充元素
	//scoreMap := map[string]string{
	//	"username": "沙和尚",
	//	"password": "123456",
	//}
	//fmt.Println(scoreMap)

	// 二、判断某个键是否存在
	//scoreMap := make(map[string]int, 8)
	//scoreMap["张三"] = 90
	//scoreMap["小李"] = 100
	//
	//v, ok := scoreMap["王二麻子"]
	//if ok {
	//	fmt.Println("这里有一个人叫做：", v)
	//} else {
	//	fmt.Println("很遗憾这里没有你找的那个人哦！")
	//}

	// 三、map 的遍历
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小李"] = 100

	// 遍历 key => value
	//for key, value := range scoreMap {
	//	fmt.Println(key, value)
	//}

	// 只遍历 key
	//for key := range scoreMap {
	//	fmt.Println(key)
	//}

	//for _, value := range scoreMap {
	//	fmt.Println(value)
	//}

	//for key, _ := range scoreMap {
	//	fmt.Println(key)
	//}

	// 删除指定的键值对
	//scoreMap["王二麻子"] = 58
	//fmt.Println(scoreMap, "增加新人之后")
	//delete(scoreMap, "王二麻子")
	//fmt.Println(scoreMap, "移除新人之后")

	//var mapSlice = make([]map[string]string, 3)
	//for index, value := range mapSlice {
	//	fmt.Printf("index:%d value:%v\n", index, value)
	//}
	//fmt.Println("after init")
	//// 对切片中的map元素进行初始化
	//mapSlice[0] = make(map[string]string, 3)
	//mapSlice[0]["name"] = "小王子"
	//mapSlice[0]["password"] = "123456"
	//mapSlice[0]["address"] = "沙河"
	//fmt.Printf("mapSlice:%p\n", mapSlice[0])
	//mapSlice[0]["iphone"] = "1367890"
	//fmt.Printf("mapSlice:%p\n", mapSlice[0])
	//fmt.Println(mapSlice)
	//fmt.Println(len(mapSlice[0]), cap(mapSlice))

	//var mapList = make(map[int]int, 2)
	//mapList[0] = 1
	//mapList[1] = 2
	//fmt.Printf("mapList:%p\n", mapList)
	//mapList[2] = 3
	//fmt.Printf("mapList:%p\n", mapList)
	//fmt.Println(mapList)

}

func sortKey() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
