package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go hello()
	fmt.Println("已经添加了子进程。。。。", time.Now())
	// 测试发现，wg.Wait() 阻塞了所在的进程 5 秒
	wg.Wait()
	fmt.Println("end", time.Now())
}

func hello() {
	defer wg.Done()
	// 模拟子进程执行了5秒
	time.Sleep(time.Second * 5)
}
