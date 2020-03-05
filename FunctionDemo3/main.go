package main

import (
	"fmt"
	"strings"
)

// 匿名函数与闭包
// 定义一个函数它的返回值是一个函数
// 把函数作为返回值
func a(name string) func() {
	return func() {
		fmt.Println("Hello ", name)
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	s := a("SliverHorn") // s此时就是一个闭包
	s()                  // 相当于执行了a函数内部的匿名函数

	r1 := makeSuffixFunc(".txt")
	result1 := r1("SliverHorn")
	fmt.Println(result1)
	r2 := makeSuffixFunc(".avi")
	result2 := r2("SliverHorn")
	fmt.Println(result2)

	// 闭包 = 函数 + 外层变量的引用
	x, y := calc(100)
	result3 := x(200) // base = 100 + 200
	fmt.Println(result3)
	result4 := y(200) // base = 300 - 200
	fmt.Println(result4)

}
