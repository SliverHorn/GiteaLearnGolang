/**
 *
 * @Description: 嵌套结构体
 * @Author: SliverHorn
 * @File: main.go
 * @Version: 1.13.8
 * @Time:  4:38 下午
 * @Project: LearnGolang
 * @Package: main
 * @Software: GoLand
 */
package main

import "fmt"

// Address 地址结构体
type Address struct {
	Province string
	City     string
}

// Person 人结构体
type Person struct {
	Name    string
	Gender  string
	Age     int
	Address // 嵌套别外一个结构体
}

func main() {
	person1 := Person{
		Name:   "SliverHorn",
		Gender: "男",
		Age:    17,
		Address: Address{
			Province: "广东",
			City:     "广州",
		},
	}
	fmt.Printf("%#v\n", person1)
	fmt.Println(person1.Name, person1.Gender, person1.Age)
	fmt.Println(person1.Address.City, person1.Address.Province) // 通过嵌套的匿名结构体字段访问其内部的字段
	fmt.Println(person1.City, person1.Province)                 // 直接访问匿名结构体中的字段
}
