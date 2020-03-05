package main

import "fmt"

// 指针
func frist() {
	a := 10 // int类型
	fmt.Printf("%v %p\n", a, &a)
	b := &a // *int类型(int指针)
	fmt.Printf("%v %p\n", b, b)
	// 变量b是一个int类型的指针(*int), 它存储的是变量a的内存地址
	c := *b
	fmt.Println(c)
}

func modify1(x int) {
	x = 100
}

func modify2(y *int) {
	*y = 100
}

func main() {
	frist()
	a := 1
	modify1(a)
	fmt.Println(a)
	modify2(&a)
	fmt.Println(a)
}
