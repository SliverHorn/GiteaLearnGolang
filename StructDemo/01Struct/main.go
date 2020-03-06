/**
 *
 * @Description: 结构体的匿名字段
 * @Author: SliverHorn
 * @File: main.go
 * @Version: 1.13.8
 * @Time:  4:12 下午
 * @Project: LearnGolang
 * @Package: main
 * @Software: GoLand
 */
package main

import "fmt"

type Person struct {
	string
	int
}

func main() {
	person1 := Person{
		"SliverHorn1",
		18,
	}
	fmt.Println(person1)
	fmt.Println(person1.string, person1.int)
}
