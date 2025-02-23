package main

import "fmt"

// StudentMgr 学生管理

// Student 定义一个学生结构体
type Student struct {
	ID    int
	Name  string
	Class string
	Age   int
	Score int
}

// NewStudent 学生信息构造函数
func NewStudent(id int, name string) *Student {
	return &Student{
		ID:    id,
		Name:  name,
		Class: "1班",
		Age:   18,
		Score: 60,
	}
}

// StudentMgr 定义一个学生管理结构体
type StudentMgr struct {
	allStudent []*Student
}

// NewStudentMgr 定义一个 StudentMgr 结构体的构造函数
func NewStudentMgr() *StudentMgr {
	return &StudentMgr{
		allStudent: make([]*Student, 0, 50),
	}
}

// AddStudent StudentMgr 定义一个方法添加学生
func (s *StudentMgr) AddStudent(stu *Student) {
	s.allStudent = append(s.allStudent, stu)
	fmt.Println("添加成功")
}

// ShowStudent StudentMgr 定义一个方法显示所有学生
func (s *StudentMgr) ShowStudent() {
	for _, v := range s.allStudent {
		fmt.Printf("学号: %d, 姓名: %s, 班级: %s, 年龄: %d, 成绩: %d\n", v.ID, v.Name, v.Class, v.Age, v.Score)
	}
}

// EditStudent StudentMgr 定义一个方法修改学生信息
func (s *StudentMgr) EditStudent(stu *Student) {
	for i, v := range s.allStudent {
		if v.ID == stu.ID {
			s.allStudent[i] = stu
			fmt.Println("修改成功")
			return
		}
	}
	fmt.Println("修改失败, 学号不存在")
}

// DeleteStudent StudentMgr 定义一个方法删除学生信息
func (s *StudentMgr) DeleteStudent(id int) {
	for i, v := range s.allStudent {
		if v.ID == id {
			s.allStudent = append(s.allStudent[:i], s.allStudent[i+1:]...)
			fmt.Println("删除成功")
			return
		}
	}
}
