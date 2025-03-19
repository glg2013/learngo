package main

import "fmt"

type userInfo struct {
	name string
	age  int
}

var graph = make(map[string]map[string]bool)

func main() {
	//fmt.Println("hello world")
	//test(2, 3)
	//
	//var n = 1
	//var p *int = &n
	//fmt.Println(p)
	//
	//var m = new(int)
	//fmt.Println(*m)
	//
	//var usr = userInfo{
	//	name: "hello",
	//}
	//fmt.Println("usr.name-> ", usr.name)
	//// 看看可以修改不
	//usr.name = "no hello"
	//fmt.Println("usr.name2-> ", usr.name)

	//x := "hello!"
	//for i := 0; i < len(x); i++ {
	//	x := x[i]
	//	if x != '!' {
	//		x := x + 'A' - 'a'
	//		fmt.Printf("%c", x)
	//	}
	//}

	//symbol := [...]string{1: "$", 2: "%", 3: "*"}
	//for i, i2 := range symbol {
	//	fmt.Println(i, i2)
	//}

	//ages := make(map[string]int)
	//fmt.Printf("%T\n %[1]v", ages)
	//if ages != nil {
	//	fmt.Println("---------", len(ages))
	//}

	//arr0 := [4]string{"1", "2", "3", "4"}
	//arr := [3]string{"a", "b", "c"}
	//s1 := arr[1:]
	//fmt.Println("s1的长度：", len(s1), " s1的容量：", cap(s1))
	//fmt.Printf("%p\n", &s1)
	//
	////s1 = append(s1, "d")
	//s1 = append(s1, arr0[:]...)
	//fmt.Println("end s1的长度：", len(s1), " s1的容量：", cap(s1))
	//fmt.Printf("%p\n", &s1)

	//res := hasEdge("a", "b")
	//fmt.Println(res)

	//fmt.Println(len(graph))
	//for k, v := range graph {
	//	fmt.Printf("key:%s,value:%v\n", k, v)
	//}

	//ch := make(chan string, 3)
	//fmt.Println(len(ch))

	//m := make(map[int]int, 2)
	//m[1] = 1
	//m[2] = 2
	//m[3] = 3
	//fmt.Println(len(m))
}

func alert(name string) func() string {
	return func() string {
		return name
	}
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func test(x ...int) {
	fmt.Println("如果没有传递参数，可以打印正常吗？")
}
