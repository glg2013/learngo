package main

import (
	"fmt"
	"os"
)

// 学员管理系统

func showMenu() {
	// 1. 显示菜单
	fmt.Println("欢迎光临学员管理系统")
	fmt.Println("请选择要执行的操作:")
	fmt.Println("1. 添加学员")
	fmt.Println("2. 更新学员")
	fmt.Println("3. 显示所有学员")
	fmt.Println("4. 删除学员")
	fmt.Println("5. 退出系统")
}

// 定义一个方法获取用户输入的选项
func getInput() *Student {
	fmt.Print("请输入学员的学号:")
	var sId int
	fmt.Scanf("%d\n", &sId)
	fmt.Print("请输入学员的姓名:")
	var sName string
	fmt.Scanf("%s\n", &sName)
	stu := NewStudent(sId, sName)
	return stu
}

func main() {
	var stuMar = NewStudentMgr()
	for {
		// 1.加载菜单
		showMenu()
		// 2.获取用户的选择
		var choice int
		fmt.Scanf("%d\n", &choice)
		switch choice {
		case 1:
			fmt.Println("选择了 添加学员")
			stu := getInput()
			// 调用添加学员的方法
			stuMar.AddStudent(stu)
		case 2:
			fmt.Println("选择了 更新学员")
			stu := getInput()
			// 调用添加学员的方法
			stuMar.EditStudent(stu)
		case 3:
			fmt.Println("选择了 显示所有学员")
			stuMar.ShowStudent()
		case 4:
			fmt.Println("选择了 删除学员")
			stu := getInput()
			// 调用添加学员的方法
			stuMar.DeleteStudent(stu.ID)
		default:
			// 退出系统
			os.Exit(0)
		}
	}
}
