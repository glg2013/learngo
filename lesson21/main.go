package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
对应 Go语言基础21 文件操作 36:09  (博客的地址 https://www.liwenzhou.com/posts/Go/file/)
*/
func main() {
	//// 只读方式打开当前目录下的 main.go 文件
	//file, err := os.Open("./main.go1")
	//if err != nil {
	//	fmt.Println("打开文件发生了错误： ", err)
	//	return
	//}
	//
	//// 关闭文件
	//file.Close()

	//readFile()
	//readAllFile()
	readFileByReadFile()

	//file, err := os.OpenFile("./readme2.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	//if err != nil {
	//	fmt.Println("open file failed, err:", err)
	//	return
	//}
	//defer file.Close()
	//str := "hello 沙河"
	//file.Write([]byte(str))       //写入字节切片数据
	//file.WriteString("hello 小王子") //直接写入字符串数据
}

func opFileInit() *os.File {
	// 只读方式打开当前目录下的 main.go 文件
	file, err := os.Open("./readme.txt")
	if err != nil {
		fmt.Println("打开 readme.txt 文件发生了错误： ", err)
		return nil
	}

	return file
}

// 只读方式读取文件--只能读取一行
func readFile() {
	file := opFileInit()

	// 使用 Read 方式读取数据
	var tmp = make([]byte, 128)
	n, err := file.Read(tmp)
	if err == io.EOF {
		fmt.Println("readme.txt 文件读取完成")
		return
	}

	if err != nil {
		fmt.Println("读取 readme.txt 文件发生了错误： ", err)
		return
	}

	fmt.Printf("读取了 %d 字节数据\n", n)
	fmt.Println(tmp)
	fmt.Println(string(tmp))

	// 关闭文件
	defer file.Close()
}

// 只读方式读取文件--读取全部
func readAllFile() {
	file := opFileInit()

	// 循环读取文件
	var content []byte
	// 使用 Read 方式读取数据
	var tmp = make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("readme.txt 文件读取完成")
			// 读取完成要退出循环
			break
		}

		if err != nil {
			fmt.Println("读取 readme.txt 文件发生了错误： ", err)
			// 有错误停止继续执行
			return
		}
		// 追加的方式汇总读取到的数据
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))

	// 关闭文件
	defer file.Close()
}

// bufio 读取文件
func readFileByBufio() {
	file := opFileInit()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件失败了。。。：", err)
	}

	// 关闭文件
	defer file.Close()
}

func readFileByReadFile() {
	fmt.Println("使用 readfile 方式")
	content, err := os.ReadFile("./readme.txt")
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return
	}
	fmt.Println(string(content))
}

func writeFile() {
	str := "hello 沙河"
	err := os.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

func writeFileByBufio() {
	file, err := os.OpenFile("xx.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello沙河\n") //将数据先写入缓存
	}
	writer.Flush() //将缓存中的内容写入文件
}
