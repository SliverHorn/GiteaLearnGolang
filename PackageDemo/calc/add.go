/**
 *
 * @Description:
 * @Author: SliverHorn
 * @File: add.go
 * @Version: 1.13.8
 * @Time:  2:18 下午
 * @Project: LearnGolang
 * @Package: calc
 * @Software: GoLand
 */
package calc

import (
	"SliverHorn/PackageDemo/snow"
	"fmt"
)

//  标识符首字符大写表示对外可见
// 通常不对外的标识符都是要首字母小写

// Name 定义一个测试的全局变量
var Name = "SliverHorn"

// Add 是一个计算两个int类型数据和的函数
func Add(x, y int) int {
	return x + y
}

// init函数在包导入的时候自动执行
// init函数没有参数也没有返回值
func init() {
	fmt.Println("cala init()")
	fmt.Println(Name)
	snow.Snow()
}
