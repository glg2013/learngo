package main

import (
	"encoding/json"
	"fmt"
)

// 结构体与 json 序列化

// Student 定义一个学生类型 结构体
type Student struct {
	ID     int
	Gender string
	Name   string
}

// Class 定义一个班级类型 结构体
type Class struct {
	Title    string
	Students []Student
}

func main() {
	c := &Class{
		Title:    "101",
		Students: make([]Student, 0, 50),
	}
	// 循环添加学生
	for i := 1; i <= 10; i++ {
		stu := Student{
			ID:     i,
			Gender: "男",
			Name:   fmt.Sprintf("stu%02d", i),
		}
		// 这里 append 为啥要接收返回值？
		// 因为 append 可能会导致扩容，扩容后返回新的切片，需要接收
		c.Students = append(c.Students, stu)
	}

	// JSON 序列化：结构体 -> JSON 格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json mashal failed", err)
		return
	}

	// 打印一下 JSON 格式的字符串
	fmt.Printf("json:%s\n", data)

	// json 反序列化：JSON 格式的字符串 -> 结构体
	str := `{"Title":"101","Students":[{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"},{"ID":10,"Gender":"男","Name":"stu10"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed", err)
		return
	}
	fmt.Printf("json:%v\n", c1)
}
