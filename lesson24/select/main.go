package main

import "fmt"

// 并发之 select

func main() {
	ch0 := make(chan int, 1)
	ch0 <- 1
	close(ch0)
	v := <-ch0
	v2 := <-ch0
	fmt.Println("获取的值是： ", v, v2)

	//ch := make(chan int, 1)
	//for i := 0; i < 10; i++ {
	//	select {
	//	case ch <- i:
	//		fmt.Println("当前存入的值是： ", i)
	//	case n := <-ch:
	//		fmt.Println("取到值了，是： ", n)
	//	default:
	//		fmt.Println("什么都没有做哦~")
	//	}
	//}

	/*
		$ go run main.go
		当前存入的值是：  0
		取到值了，是：  0
		当前存入的值是：  2
		取到值了，是：  2
		当前存入的值是：  4
		取到值了，是：  4
		当前存入的值是：  6
		取到值了，是：  6
		当前存入的值是：  8
		取到值了，是：  8
	*/
	// 这个输出结果是因为，缓冲通道只有 1 个容量，所以先存入的是 0，下次 for 循环到 1 的时候，0 还没有被取出去，所以
}
