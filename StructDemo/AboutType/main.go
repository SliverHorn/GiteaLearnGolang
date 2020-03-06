/**
 *
 * @Description: 自定义类型和类型别名示例
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

// 1.自定义类型

// MyInt 基于int类型的自定义类型
type MyInt int

// 2.类型别名

// NewInt int类型别名
type NewInt = int

func main() {
	var a MyInt
	fmt.Printf("Tpye:%T, Value:%v\n", a, a)
	var x NewInt
	fmt.Printf("Tpye:%T, Value:%v\n", x, x)
}
