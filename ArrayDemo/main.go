package main

import "fmt"

// 数组相关内容

func frist() {
	fmt.Println("-------数组定义")
	var a [3]int
	var b [4]int
	// a = b
	fmt.Println(a)
	fmt.Println(b)

}

func second() {
	fmt.Println("-------数据的初始化")

	// 1. 定义时使用初始值列表的方式初始化
	var cityArray = [4]string{"北京", "上海", "广州", "深圳"}
	fmt.Println(cityArray)

	// 2. 编译器推导数组的长度
	var boolArray = [...]bool{true, false, false, true, false}
	fmt.Println(boolArray)

	// 3. 使用索引值方式初始化
	var langArray = [...]string{1: "Golang", 3: "Python", 7: "C++"}
	fmt.Println(langArray)
}

func third() {
	fmt.Println("-------数组的遍历")
	var CityArray = [4]string{"北京", "上海", "广州", "深圳"}

	// 1. for循环遍历
	for i := 0; i < len(CityArray); i++ {
		fmt.Println(CityArray[i])
	}

	// 2. for range遍历
	for _, value := range CityArray {
		fmt.Println(value)
	}
}

func fourth() {
	fmt.Println("-------多维数组")
	// 多维数组, 第一层可以使用编译器推导数组的长度,其他层不可以
	cityArray := [4][2]string{
		{"北京", "西安"},
		{"上海", "杭州"},
		{"重庆", "成都"},
		{"广州", "深圳"},
	}
	fmt.Println(cityArray)
	// 打印重庆
	fmt.Println(cityArray[2][0])

	for _, v1 := range cityArray {
		fmt.Printf("v1=%s\n", v1)
		for _, v2 := range v1 {
			fmt.Printf("v2=%s\n", v2)
		}
	}
}

func fifth() {
	fmt.Println("-------数组是值类型")
	x := [3]int{1, 2, 3}
	fmt.Println(x)
	f1(x)
	fmt.Println(x)

	y := [3][2]int{
		{1,2},
		{3,4},
		{5,6},
	}
	fmt.Println(y)
	f2(y)
	fmt.Println(y)
	z := y
	z[0][0] = 1000
	fmt.Println(y)
}

func f1(a [3]int) {
	a[0] = 100
}

func f2(a [3][2]int) {
	a[0][0] = 100
}

func main() {
	frist()
	second()
	third()
	fourth()
	fifth()
}
