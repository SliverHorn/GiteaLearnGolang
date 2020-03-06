/**
 *
 * @Description: 方法的定义实例
 * @Author: SliverHorn
 * @File: main.go
 * @Version: 1.13.8
 * @Time:  2:46 下午
 * @Project: LearnGolang
 * @Package: main
 * @Software: GoLand
 */
package main

import "fmt"

// Person 是一个结构体
type Person struct {
	name string
	age  int8
}

// NewPerson 是Person的构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// Dream 是为Person类型定义方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好GO语言\n", p.name)
}

// SetAge 是一个修改年龄的方法
// 指针接收者指的就是接收者的类型是指针
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

// SetAge2 是一个使用值接收者修改年龄的方法
func (p Person) SetAge2(newAge int8) {
	p.age = newAge
}

func main() {
	person8 := NewPerson("SliverHorn8", 17)
	(*person8).Dream()
	person8.Dream()
	person8.SetAge(int8(20))
	fmt.Println(person8)
	person8.SetAge2(int8(30))
	fmt.Println(person8)

}
