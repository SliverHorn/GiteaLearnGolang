/**
 *
 * @Description:
 * @Author: SliverHorn
 * @File: main.go
 * @Version: 1.13.8
 * @Time:  2:19 下午
 * @Project: LearnGolang
 * @Package: main
 * @Software: GoLand
 */
package main

// 当你的代码都放在$GOPATH目录下
// 导入包的路径要从$GOPATH/src后面继续写起

import (
	a "SliverHorn/PackageDemo/calc" // 给导入的包起别名
	"fmt"
)

// Go语言中不允许导入包而不使用!!!
// Go语言中不允许循环引用包!!!

func main() {
	sum := a.Add(1, 3)
	fmt.Println(sum)
	fmt.Println(a.Name)
}

// 多用做一些初始化的操作
func init() {
	fmt.Println()
}
