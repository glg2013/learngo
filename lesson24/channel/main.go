package main

import (
	"fmt"
	"sync"
)

/*
对应 Go语言基础25 并发编程之channel 53:19
2025年2月26日11:20:44
*/

func main() {
	// 无缓冲通道调用
	// 因为没有添加 sync.WaitGroup 会导致 main 退出，recv 没有接收到
	//send()

	// 有缓冲通道调用
	// 同样因为没有添加 sync.WaitGroup 会导致 main 退出，recv 没有接收到
	send2()
}

/*
无缓冲通道 - 发送
*/
func send() {
	var wg sync.WaitGroup
	wg.Add(1)

	var ch chan int = make(chan int)
	go func() {
		// 创建一个 goroutine 从通道接收值
		recv(ch)
		wg.Done()
	}()
	ch <- 10
	fmt.Println("发送成功！")
	wg.Wait()
}

/*
无缓冲通道 - 接收
*/
func recv(ch chan int) {
	ret := <-ch
	fmt.Println("接收成功： ", ret)
}

/*
有缓冲通道 - 发送
*/
func send2() {
	var ch chan int = make(chan int, 2)
	//go recv2(ch)
	ch <- 100
	ch <- 200
	fmt.Println("发送成功~")
	close(ch)
	//recv2(ch)
	recv3(ch)
}

/*
有缓冲通道 - 接收
通道内的值获取完后还会继续取，如果通道关闭，就会取到类型的零值
*/
func recv2(c chan int) {
	for {
		ret, err := <-c
		if !err {
			fmt.Println("有错误发生了~", err, ret)
			break
		}
		fmt.Println("接收成功：", ret)
	}
}

/*
通道内所有值获取完毕后会自动退出
*/
func recv3(ch chan int) {
	for v := range ch {
		fmt.Println("接收到的值是：", v)
	}
}
