package main

import "fmt"

// Defaulter 定义了一个接口，有一个初始化结构体的方法
type Defaulter interface {
	InitDefault()
}

type Foo struct {
	Msg   string
	Value int
}

func (f *Foo) InitDefault() {
	f.Msg = "Foo 默认值"
	f.Value = 10
}

type Bar struct {
	Msg   string
	Value float64
}

func (b *Bar) InitDefault() {
	b.Msg = "Bar 默认值"
	b.Value = 5.5
}

type defaulterPtr[T any] interface {
	*T
	Defaulter
}

func Default[T Defaulter, P defaulterPtr[T]]() T {
	var v T
	P.InitDefault(&v)
	return v
}

func main() {

	//f := &Foo{}
	//f.InitDefault()
	//fmt.Println(f)

	foo := Default[Foo]()
	bar := Default[Bar]()

	fmt.Println("foo", foo)
	fmt.Println("bar", bar)

}
