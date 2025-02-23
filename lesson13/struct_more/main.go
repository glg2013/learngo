package main

import "fmt"

// 嵌套结构体

// Address 地址结构体
type Address struct {
	Province string
	City     string
}

// User 用户结构体
type User struct {
	Name    string
	Gender  string
	Address Address // 嵌套了 地址 Address 结构体
}

func main() {

	// 嵌套结构体
	user1 := User{
		Name:   "小王子",
		Gender: "男",
		Address: Address{
			Province: "山东",
			City:     "威海",
		},
	}
	fmt.Printf("user1: %#v\n", user1)
	fmt.Println(user1)
}
