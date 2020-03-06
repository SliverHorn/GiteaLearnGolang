/**
 *
 * @Description: 结构体的继承
 * @Author: SliverHorn
 * @File: main.go
 * @Version: 1.13.8
 * @Time:  6:06 下午
 * @Project: LearnGolang
 * @Package: main
 * @Software: GoLand
 */
package main

import "fmt"

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动\n", a.name)
}

type Dog struct {
	Feet    int8
	*Animal // 匿名嵌套, 而且嵌套的是一个结构体指针
}

func (d *Dog) wang() {
	fmt.Printf("%s会旺旺\n", d.name)
}

func main() {
	d1 := &Dog{
		Feet:   4,
		Animal: &Animal{name: "李华"},
	}
	d1.move()
	d1.wang()

}
