package main

import (
	"fmt"
	"unsafe"
)

// 定义一个结构体
type person struct {
	name string
	city string
	age  int
}

func main() {
	// 调用构造函数
	p1 := newPerson("张三", "北京", 18)
	fmt.Printf("p1:%#v \n", p1)
	fmt.Println(p1, unsafe.Sizeof(p1))
}

// 定义一个构造函数
// 注意：Go语言的结构体没有构造函数，一般可以使用NewXXX来初始化结构体
func newPerson(name, city string, age int) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}
