package main

import "fmt"

type Number interface {
	int | float64
}

func head[T any](slice []T) (*T, bool) {
	if len(slice) > 0 {
		return &slice[0], true
	}
	return nil, false
}

func min[T Number](a, b T) T {
	if a <= b {
		return a
	}
	return b
}

func main() {
	slice := []int{1, 2, 3}
	a, ok := head(slice)
	fmt.Println(a, ok)

	slice2 := []string{""}
	b, _ := head(slice2)
	fmt.Println(b)

	f := head[int]
	fmt.Println(f(slice))
}
