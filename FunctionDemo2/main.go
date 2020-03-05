package main

import (
	"fmt"
)

// 匿名函数和闭包

// 定义一个函数返回一个函数
func a() func() {
	return func() {
		fmt.Println("我是匿名函数")
	}
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}
func main() {
	func() {
		fmt.Println("定义并执行匿名函数")
	}()
	s1 := func() { fmt.Println("匿名函数执行方式2") }
	s1()
	a()()
	r1 := calc(100, 200, add)
	fmt.Println(r1)
	r2 := calc(100, 200, sub)
	fmt.Println(r2)
}
