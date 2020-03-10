/**
 *
 * @Description: 空接口
 * @Author: SliverHorn
 * @File: main.go
 * @Version: 1.13.8
 * @Time:  3:37 下午
 * @Project: Go
 * @Package: main
 * @Software: GoLand
 */
package main

// 接口中没有定义任何需要实现的方法时,该接口就是一个空接口
// 任意类型都实现了空接口 --> 空接口变量可存储任意值!

import "fmt"

// 空接口一般不需要提前定义
//type xxx interface{}

// 空接口的应用
// 1. 空接口类型作为函数的参数
// 2. 空接口作为map的key或value
// 3.

func main() {
	var x interface{}
	x = 100
	fmt.Println(x)
	x = "Golang"
	fmt.Println(x)
	x = false
	fmt.Println(x)
	var m = make(map[interface{}]interface{}, 16)
	m["name"] = "SliverHorn"
	m["age"] = 17
	m["hobby"] = []string{"Code", "Running"}
	fmt.Println(m)
	ret, ok := x.(string) // 类型断言, 猜的类型不对时,ok=false,ret=string类型的零值
	if !ok {
		fmt.Println("不是字符串类型")
	} else {
		fmt.Println("是字符串了类型", ret)
	}

	// 使用switch进行类型断言
	switch value := x.(type) {
	case string:
		fmt.Printf("是字符串类型, value=%v\n", value)
	case bool:
		fmt.Printf("是布尔类型, value=%v\n", value)
	case int:
		fmt.Printf("是int类型, value=%v\n", value)
	default:
		fmt.Printf("猜不到!")
	}

}
