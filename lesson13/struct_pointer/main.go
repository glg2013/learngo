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
	//fmt.Println(p2)
	//fmt.Printf("%T\n", p2) // *main.person
	////fmt.Printf("p2=%#v\n", p2) // p2=&main.person{name:"", city:"", age:0}
	//var p3 = &person{}
	//fmt.Println(p3)
	//
	//var p4 = person{}
	//fmt.Println(p4)

	//// 2.取结构体的地址实例化
	//p3 := &person{}
	//fmt.Printf("%T\n", p3)     // *main.person
	//fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
	//p3.name = "小王子"
	//p3.age = 18
	//p3.city = "北京"
	//fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"小王子", city:"北京", age:18}

	//// 3.结构体初始化
	//var p4 person
	//fmt.Printf("%#v\n", p4) //没有初始化，main.person{name:"", city:"", age:0}
	//
	//// 使用键值对初始化
	//p5 := person{
	//	name: "小王子",
	//	city: "北京",
	//	age:  18,
	//}
	//fmt.Printf("%#v\n", p5) //初始化 main.person{name:"小王子", city:"北京", age:18}
	//
	//// 对结构体指针进行键值对初始化
	//p6 := &person{
	//	name: "小王子",
	//	city: "北京",
	//	age:  18,
	//}
	//fmt.Printf("%#v\n", p6)
	//
	//// 使用零值初始化
	//p7 := &person{
	//	name: "小王子",
	//	// 这里城市没有初始化，就会用对应类型的零值，string 就是 ""
	//	age: 18,
	//}
	//fmt.Println("=====================")
	//fmt.Printf("%#v\n", p7)
	//fmt.Println(p7.name)
	//
	//// 使用值的列表初始化
	//p8 := &person{
	//	"小王子",
	//	"北京",
	//	18,
	//}
	//fmt.Printf("%#v\n", p8)
	//
	//// 空结构体
	//var v struct{}
	//fmt.Println(unsafe.Sizeof(v))
}
