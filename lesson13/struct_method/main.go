package main

import "fmt"

// 方法和接收者

// Person 定义一个结构体
type Person struct {
	name string
	age  int
}

type MyInt int

// 结构体的匿名字段
type question struct {
	string
	int
}

func main() {
	// 调用方法
	p := NewPerson("李大毛", 18)
	fmt.Printf("p: %#v\n", p)
	p.Dream()
	Dream(*p)

	// 调用给自定义类型添加的方法
	var m MyInt
	m.Instruction()

	// 使用结构体的匿名字段
	q := question{
		"小王子",
		int(m),
	}
	fmt.Printf("q: %#v\n", q) // q: main.question{string:"小王子", int:0}
}

// NewPerson 定义一个构造函数
func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// Dream 给 Person 结构体添加一个方法
func (p *Person) Dream() {
	p.name = "李二毛"
	fmt.Printf("%s 的Go语言学习之路一定会坚持到底的！", p.name)
}

// Dream 定义一个和 Dream 方法相同作用的函数
func Dream(p Person) {
	fmt.Printf("%#v\n", p)
	//p.name = "李三毛"
	fmt.Printf("%s 的Go语言学习之路一定会坚持到底的！", p.name)
}

// Instruction 给我们自定义的类型添加方法
func (m MyInt) Instruction() {
	fmt.Println("我是自定义的一个基于 int 类型的类型")
}
