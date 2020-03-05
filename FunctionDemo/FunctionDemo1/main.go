package main

import "fmt"

// 定义一个不需要参数也没有返回值的函数:sayHello
func sayHello() {
	fmt.Println("Hello Golang")
}

// 定义一个接收string类型的name参数
func sayHi(name string) {
	fmt.Printf("Hello %s\n", name)
}

// 定义接收多个参数的函数并且有一个返回值
func intSum1(a int, b int) int {
	result := a + b
	return result
}

func intSum2(a int, b int) (result int) {
	result = a + b
	return
}

// 函数可变参数, 在参数名后面加... 表示可变参数
// 可变参数在函数体中是切片类型,
func intSum3(a ...int) int {
	fmt.Println(a)
	fmt.Printf("%#v\n", a)
	result := 0
	for _, arg := range a {
		result += arg
	}
	return result
}

// 固定参数和可变参数同时出现时, 可变参数要放在最后
// go语言的函数中没有默认参数
func intSum4(a int, b ...int) int {
	result := a
	for _, arg := range b {
		result += arg
	}
	return result
}

// Go语言中函数参数类型简写
func intSum5(a, b int) (result int) {
	result = a + b
	return
}

// 定义具有多个返回值的函数, 多返回值也支持类型简写
func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

// 函数进阶之变量作用域
//var num = 10

// 定义函数
func testGlobal() {
	num := 100
	name := "SliverHorn"
	// 可以在函数中使用变量
	// 1. 现在自己函数中查找,找到了就用自己函数中的
	// 2. 函数中找不到就往外层找全局变量
	//fmt.Println("全局变量", num)
	fmt.Println("变量num", num)
	fmt.Println(name)
}

func main() {
	// 调用函数
	sayHello()
	sayHi("Python")
	sum1 := intSum1(1, 4)
	fmt.Println(sum1)
	sum2 := intSum2(1, 4)
	fmt.Println(sum2)
	result3 := intSum3(1, 2, 3, 4, 5, 6)
	fmt.Println(result3)
	result4 := intSum4(1, 2, 3, 4, 5, 6)
	fmt.Println(result4)
	result5 := intSum5(1, 2)
	fmt.Println(result5)
	sum, sub := calc(result3, result4)
	fmt.Println(sum, sub)
	testGlobal()
	// 外层不能访问到内层函数的内部变量
	//fmt.Println(name)

	// 变量i此时只在for循环的语句块中生效
	for i := 0; i < 7; i++ {
		fmt.Println(i)
	}
	//fmt.Println(i)  // 外层访问不到内部for语句中的变量
	abc := testGlobal
	fmt.Printf("%T\n", abc)

}
