/**
 *
 * @Description: 使用值接收者实现接口和使用指针接受者实现接口的区别
 * @Author: SliverHorn
 * @File: main.go
 * @Version: 1.13.8
 * @Time:  1:22 下午
 * @Project: LearnGolang
 * @Package:
 * @Software: GoLand
 */
package main

import "fmt"

type animal interface {
    mover
    sayer
}
type mover interface {
	move()
}

type sayer interface {
	say()
}

type person struct {
	name string
	age  int8
}

func (p *person) say() {
	fmt.Printf("%s在叫...\n", p.name)
}

//// 使用值接收者实现接口: 类型的值和类型的指针都能在够保存到接口变量中
//func (p person) move() {
//	fmt.Printf("%s在动\n", p.name)
//}

// 使用指针接收者实现接口:只有类型指针能够保存到接口变量中
func (p *person) move() {
	fmt.Printf("%s在动\n", p.name)
}

func main() {
	var m mover // 定义一个mover类型的变量
	//p1 := person{ // person类型的值
	//	name: "SliverHorn2",
	//	age:  17,
	//}
	p2 := &person{ // p2是person类型的指针
		name: "SliverHorn2",
		age:  17,
	}
	//m = p1  // 无法赋值,因为p1是person类型的值没有实现mover接口
	m = p2
	m.move()
	fmt.Println(m)
	var s sayer = p2 // 定义一个sayer类型的变量
	p2.say()
	fmt.Println(s)
	var a animal = p2  // 定义一个animal类型的变量
	a.say()
	a.move()
	fmt.Println(a)
}
