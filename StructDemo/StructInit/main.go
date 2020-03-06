/**
 *
 * @Description:结构体的初始化
 * @Author: SliverHorn
 * @File: main.go
 * @Version: 1.13.8
 * @Time:  2:17 下午
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

func frist() {
	fmt.Println("-------1.键值对初始化")
	var person4 = person{
		name: "SliverHorn4",
		city: "广州4",
		age:  17,
	}
	fmt.Printf("person4=%#v\n", person4)
	var person5 = &person{
		name: "SliverHorn5",
		city: "广州5",
		age:  17,
	}
	fmt.Printf("person4=%#v\n", person5)
}

func second() {
	fmt.Println("-------2.值的列表进行初始化")
	var person6 = person{"SliverHorn6", "广州6", 17}
	fmt.Printf("%#v\n", person6)
}

func third() {
	fmt.Println("-------third")

}

func fourth() {
	fmt.Println("-------Fourth")

}

func fifth() {
	fmt.Println("-------Fifth")

}

func main() {
	frist()
	second()
	third()
	fourth()
	fifth()
}
