package main

import "fmt"

func main() {
	/**
	1.编写代码打印9*9乘法表。
	*/
	for i := 1; i <= 9; i++ {
		for l := 1; l <= 9; l++ {
			if i > l {
				fmt.Printf("%d x %d = %d ", l, i, i*l)
			} else if i == l {
				fmt.Printf("%d x %d = %d\n", l, i, i*l)
			}
		}
	}
}
