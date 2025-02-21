package main

/*
对应 Go语言基础 13结构体定义与实例化
指针
2025年2月21日17:44:49
*/

type person struct {
	name string
	city string
	age  int8
}

func main() {

	//p4 := &person{
	//	name: "小李",
	//	city: "上海",
	//	age:  18,
	//}
	//fmt.Printf("type of p4: %#v\n", p4) // type of p4: &main.person{name:"小李", city:"上海", age:18}

	//// 1.创建指针类型的结构体
	//var p2 = new(person)
	//fmt.Printf("%T\n", p2) // *main.person
	//fmt.Printf("p2=%#v\n", p2)

	// 2.去结构体的地址实例化
	// 到家继续
}
