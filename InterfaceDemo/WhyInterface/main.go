/**
 *
 * @Description:
 * @Author: SliverHorn
 * @File: main.go
 * @Version: 1.13.8
 * @Time:  5:46 下午
 * @Project: LearnGolang
 * @Package:
 * @Software: GoLand
 */
package main

import "fmt"

type dog struct {
}

func (d dog) say() {
	fmt.Println("汪汪汪~")
}

type cat struct {
}

func (c cat) say() {
	fmt.Println("喵喵喵~")
}

type person struct {
	name string
}

func (p person) say() {
	fmt.Println("啊啊啊~")
}

// 定义一个类型, 一个抽象的类型,只要实现了say()这个方法的类型都可以称为sayer类型
type sayer interface {
	say()
}

func do(arg sayer) {
	arg.say() // 不管进来的是什么,都要执行它的say
}
func main() {
	c1 := cat{}
	do(c1)
	d1 := dog{}
	do(d1)
	p1 := person{"SliverHorn"}
	do(p1)
	var s sayer
	//c2 := cat{}
	//s = c2
	p2 := person{"SliverHorn"}
	s = p2
	fmt.Println(s)
}
