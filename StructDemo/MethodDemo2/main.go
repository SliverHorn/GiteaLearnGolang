/**
 *
 * @Description:为任意类型添加方法
 * @Author: SliverHorn
 * @File: main.go
 * @Version: 1.13.8
 * @Time:  3:34 下午
 * @Project: LearnGolang
 * @Package:main
 * @Software: GoLand
 */
package main

import "fmt"

type myInt int

func (m myInt) sayhi() {
	fmt.Println("hi")
}

func main() {
	var m1 myInt = 100
	m1.sayhi()
}
