package main

import "fmt"

// 事先定义好的类型约束类型

type Value interface {
	int | float64
}

func min[T Value](a, b T) T {
	if a <= b {
		return a
	}
	return b
}

func main() {
	//fmt.Println(reverse([]int{1, 2, 3, 4}))
	//fmt.Println(reverse2([]interface{}{1, 2, 3, 4}))
	ret := min(1, 0.1)
	fmt.Println(ret)
}

func reverse(s []int) []int {
	l := len(s)
	r := make([]int, l)

	for i, e := range s {
		r[l-i-1] = e
	}
	return r
}

func reverse2(s []interface{}) []interface{} {
	l := len(s)
	r := make([]interface{}, l)

	for i, e := range s {
		r[l-i-1] = e
	}
	return r
}

func reverseWithGenerics[T any](s []T) []T {
	l := len(s)
	r := make([]T, l)

	for i, e := range s {
		r[l-i-1] = e
	}
	return r
}

func reve[T any](s []T) []T {
	l := len(s)
	r := make([]T, l)

	for i, e := range s {
		r[l-i-1] = e
	}
	return r
}
