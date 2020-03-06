/**
 *
 * @Description: 结构体构造函数:构造一个结构体实例的函数
 * @Author: SliverHorn
 * @File: main.go
 * @Version: 1.13.8
 * @Time:  2:30 下午
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

// newPerson person的构造函数
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

func main() {
	person7 := newPerson("sliverHorn7", "广州7", 17)
	fmt.Printf("Type:%T, value:%#v\n", person7, person7)
}
