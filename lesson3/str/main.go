package main

import (
	"fmt"
	"strings"
)

/**
对应 Go语言基础03 基本数据类型
数据类型 - 字符串操作
单引号表示字符 - 阿斯克码
双引号表示字符串 - 普通字符串
*/

func main() {

	// 常见字符串操作
	var s string = "hello"
	var s2 string = "hello小宇宙"
	fmt.Println(len(s), len(s2))

	// 拼接字符串
	fmt.Println(s + s2)
	var s3 string = fmt.Sprintf("%s - %s", s, s2)
	fmt.Println(s3)

	// 字符串的分割
	var s4 string = "how do you do"
	fmt.Println(strings.Split(s4, " "))
	fmt.Printf("%T\n", strings.Split(s4, " "))

	// 判断是否包含
	fmt.Println(strings.Contains(s4, "do"))

	// 判断前缀
	fmt.Println(strings.HasPrefix(s4, "how"))
	// 判断后缀
	fmt.Println(strings.HasSuffix(s4, "how"))

	// 判断子串的位置
	fmt.Println(strings.Index(s4, "do"))
	// 最后子串出现的位置
	fmt.Println(strings.LastIndex(s4, "do"))
	// join 连接
	// 这个是使用了 var 的形式，定义并赋值
	var s5 []string = []string{"how", "do", "you", "do"}
	fmt.Println(s5)
	fmt.Println(strings.Join(s5, "+"))

	// 这个是使用了 简短声明的形式
	s6 := []string{"here", "you", "are"}
	fmt.Println(s6)
	fmt.Println(strings.Join(s6, "@"))

	// 类型别名 byte rune
	var c1 byte = 'c'
	var c2 rune = 'c'
	fmt.Println(c1, c2)
	fmt.Printf("c1 %T  c2 %T\n", c1, c2)

	// 遍历查询字符串的每个字符
	// 普通的遍历，如果是有多字节编码的字符串，就会出现乱码
	s7 := "老子打的就是精锐！"
	s7Len := len(s7)
	for i := 0; i < s7Len; i++ {
		fmt.Printf("%c\n", s7[i])
	}

	// 处理中英文混合的
	var s8 string = "hello小宇宙"
	for _, i2 := range s8 {
		// 单个字符就是 阿斯克码的 int 类型了
		fmt.Printf("%T\n, %c\n", i2, i2)
	}

	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	fmt.Printf("%T\n", byteS1)
	fmt.Println(byteS1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s9 := "白萝卜"
	runeS2 := []rune(s9)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))

}
