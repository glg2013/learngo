package main

import "fmt"

// 结构体的继承

type Animal struct {
	name string
}

func (a *Animal) jump() {
	fmt.Printf("接收者的内存地址 %p\n", a) // 这里的 a 是 Animal 类型的副本，不是指针，所以这里的 a 的内存地址和 cat.Animal 的内存地址是不同的
	fmt.Printf("%s 能 跳\n", a.name)
}

type Cat struct {
	eyes int
	//*Animal		// 直接嵌入 Animal 类型的指针
	//Animal		// 直接嵌入 Animal 结构体
	Animal Animal // 当成字段来嵌入 Animal 结构体
}

func (c *Cat) ability() {
	fmt.Printf("cat接收者的内存地址 %p\n", c)
	fmt.Printf("%s 的特殊本领是 爬树！\n", c.Animal.name)
}

func main() {
	// 结构体的继承 使用案例
	cat := Cat{
		eyes: 2,
		Animal: Animal{
			name: "小黄",
		},
	}
	fmt.Printf("cat.内存地址 %p\n", &cat)
	cat.Animal.jump() // 明明 jump 方法的接收者是 Animal 类型的指针，为什么这里可以直接调用？，因为 go 会自动转换，底层会自动转换为 (&cat).jump()
	cat.ability()
}
