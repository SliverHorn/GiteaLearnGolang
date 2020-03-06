/**
 *
 * @Description: 结构体指针
 * @Author: SliverHorn
 * @File: main.go
 * @Version: 1.13.8
 * @Time: 2020/3/6 1:51 下午
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

func main() {
	var person2 = new(person)
	//(*person2).name = "SliverHorn"
	//(*person2).city = "广州"
	//(*person2).age = 17
	person2.name = "SliverHorn"
	person2.city = "广州"
	person2.age = 17
	fmt.Printf("%#v\n", person2)

	// 去结构体的地址进行实例化
	person3 := &person{}
	fmt.Printf("Type=%T,Value=%#v\n", person3, person3)
	person3.name = "SliverHorn3"
	person3.city = "广州3"
	person3.age = 17
	fmt.Printf("Type=%T,Value=%#v\n", person3, person3)
}
