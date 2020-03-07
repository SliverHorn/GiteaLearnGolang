/**
 *
 * @Description: 结构体字段的可见性与JSON序列化
 * @Author: SliverHorn
 * @File: main.go
 * @Version: 1.13.8
 * @Time:  10:03 上午
 * @Project: LearnGolang
 * @Package: main
 * @Software: GoLand
 */
package main

import (
	json "encoding/json"
	"fmt"
)

// 如果一个Go语言包中定义的标识符是首字母大写的,那么就是对外可见的.
// 如果一个结构体中的字段名首字母是大写的, 那么该字符就是对外可见的

type student struct {
	ID   int
	Name string
}

// student 构造函数
func newStudent(id int, name string) *student {
	return &student{id, name}
}

type class struct {
	Title   string    `json:"title"`
	Student []student `json:"student_list"`
}

func main() {
	s1 := student{
		ID:   1,
		Name: "",
	}
	fmt.Println(s1)
	// 创建一个班级变量c1
	c1 := class{
		Title:   "Class01",
		Student: make([]student, 0, 10),
	}
	fmt.Println(c1)

	// 往班级c1中添加学生
	for i := 0; i < 10; i++ {
		// 模拟10个学生
		tmpStu := newStudent(i, fmt.Sprintf("Stu%02d", i))
		c1.Student = append(c1.Student, *tmpStu)
	}
	fmt.Printf("%#v\n", c1)

	// JSON序列化:Go语言中的数据 -> JSON格式化的字符串
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("Json marshal failed, err:", err)
		return
	}
	fmt.Printf("%T\n", data)
	fmt.Printf("%s\n", data)

	// JSON反序列化: JSON格式的字符串 --> 结构体
	str := `{"Title":"Class01","Student":[{"ID":0,"Name":"Stu00"},{"ID":1,"Name":"Stu01"},{"ID":2,"Name":"Stu02"},{"ID":3,"Name":"Stu03"},{"ID":4,"Name":"Stu04"},{"ID":5,"Name":"Stu05"},{"ID":6,"Name":"Stu06"},{"ID":7,"Name":"Stu07"},{"ID":8,"Name":"Stu08"},{"ID":9,"Name":"Stu09"}]}`
	c2 := &class{}
	err = json.Unmarshal([]byte(str), c2)
	if err != nil {
		fmt.Println("Json unmarshal failed")
		return
	}
	fmt.Printf("%T\n", c2)
	fmt.Printf("%#v\n", c2.Student)
}
