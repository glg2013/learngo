package main

import "fmt"

func main() {
	//n := 123_456_789
	//fmt.Println(n)

	for i := range 10 {
		fmt.Println(i)
	}

	var testA [3]string
	fmt.Print("%T\n", testA)
	fmt.Println(len(testA))
}
