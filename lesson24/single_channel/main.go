package main

import "fmt"

/*
单向通道
*/
func main() {
	ch := Producer()

	res := Consumer(ch)
	fmt.Println("main 。。。 ", res)
}

func Producer() <-chan int {
	ch := make(chan int, 2)
	// 创建一个新的 goroutine 执行发送数据的任务
	go func() {
		for i := 0; i < 5; i++ {
			if i%2 == 0 {
				ch <- i
			}
		}
		close(ch)
	}()

	return ch
}

// Consumer /*
func Consumer(ch <-chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}
