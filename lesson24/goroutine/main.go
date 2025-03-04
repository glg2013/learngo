package main

import (
	"fmt"
	"sync"
)

/*
对应 Go语言基础24 并发编程之goroutine 38:25
2025年2月26日11:20:52
*/

// 声明全局等待组变量
//var wg sync.WaitGroup

func main() {
	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1)
		go hello(i, &wg)
	}
	wg.Wait()
}

func hello(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("hello：", i)
}

// 练习题1
// 请写出下面程序的执行结果
func exam() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}
