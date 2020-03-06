/**
 *
 * @Description: 定义结构体
 * @Author: SliverHorn
 * @File: main.go
 * @Version: 1.13.8
 * @Time: 2020/3/6 11:25 上午
 * @Project: LearnGolang
 * @Package:main
 * @Software: GoLand
 */
package main

import "fmt"

type person struct {
	name, city string
	age        int8
}

// 匿名结构体
func frist() {
	var user struct {
		Name string
		Age  int
	}
	user.Name = "SliverHorn"
	user.Age = 17
	fmt.Printf("%#v\n", user)
}

func main() {
	var person1 person
	person1.name = "SliverHorn"
	person1.city = "广州"
	person1.age = 17
	fmt.Printf("person1=%#v\n", person1)
	fmt.Println(person1.name)
	fmt.Println(person1.city)
	fmt.Println(person1.age)

	frist()
}
