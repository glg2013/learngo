package main

import (
	"encoding/json"
	"fmt"
)

// tag

// Student 定义一个学生结构体
type Student struct {
	ID     int `json:"id"`
	Gender string
	name   string
}

func main() {
	s1 := Student{
		ID:     1,
		Gender: "男",
		name:   "小王子",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json err:", err)
		return
	}
	fmt.Printf("json str:%s\n", data)
}
