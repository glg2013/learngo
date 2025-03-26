package main

import "fmt"

func main() {
	//printMsg(1, "2", "hello")

	//s := ""
	//res := p(s)
	//fmt.Println(res == nil)

	//var p *int
	//p = new(int)
	////fmt.Println(*p)
	////fmt.Printf("%T\n %p\n %v", *p, p, *p)
	//if p == nil {
	//	fmt.Println("是 nil")
	//}

	//var s *string
	//s = new(string)
	//str := "hello"
	//s = &str
	//fmt.Println(*s)

	//n := 10
	//defer func(n int) {
	//	fmt.Println(n)
	//}(n)
	//
	//n = 20
	
}

func printMsg(msg ...any) {
	fmt.Println(len(msg))
	fmt.Println(msg...)
}

func p(msg string) []string {
	if msg == "" {
		return nil
	}
	return []string{msg}
}
