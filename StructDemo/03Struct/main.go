/**
 *
 * @Description: 嵌套结构体的字段冲突
 * @Author: SliverHorn
 * @File: main.go
 * @Version: 1.13.8
 * @Time:  5:37 下午
 * @Project: LearnGolang
 * @Package: main
 * @Software: GoLand
 */
package main

import "fmt"

// Address 地址结构体
type Address struct {
	Province   string
	City       string
	UpdateTime string
}

type Email struct {
	Addr       string
	UpdateTime string
}

// Person 人结构体
type Person struct {
	Name    string
	Gender  string
	Age     int
	Address // 嵌套别外一个结构体
	Email   // 嵌套别外一个结构体
}

func main() {
	person2 := Person{
		Name:   "SliverHorn",
		Gender: "男",
		Age:    17,
		Address: Address{
			Province:   "广东",
			City:       "广州",
			UpdateTime: "2020-3-6",
		},
		Email: Email{
			Addr:       "sliver_horn@qq.com",
			UpdateTime: "2020-3-6",
		},
	}
	//fmt.Println(person2.UpdateTime) 嵌套结构体中包含多个同名的UpdateTime字段
	fmt.Println(person2.Address.UpdateTime)
	fmt.Println(person2.Email.UpdateTime)
}
