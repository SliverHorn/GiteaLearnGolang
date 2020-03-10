package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

// map(映射)

func frist() {
	fmt.Println("-------声明map类型")
	var a map[string]int // 光声明map类型, 但是没有初始化, a就是初始化nil
	fmt.Println(a == nil)

}

func second() {
	fmt.Println("-------map的初始化")
	a := make(map[string]int, 8)
	fmt.Println(a == nil)
}

func third() {
	fmt.Println("-------map中添加键值对")
	a := make(map[string]int, 8)
	a["1"] = 1
	a["2"] = 2
	fmt.Println(a)
	fmt.Printf("a:%#v\n", a)
	fmt.Printf("Tpye=%T\n", a)
}

func fourth() {
	fmt.Println("-------声明map的同时完成初始化")
	b := map[int]bool{
		1: true,
		2: false,
	}
	fmt.Printf("b:%#v\n", b)
	fmt.Printf("Type:%T\n", b)

}

func fifth() {
	fmt.Println("-------map没有初始化不能赋值")
	var c map[int]int
	//c[100] = 200 // c这个map没有初始化不能直接操作,编译不会出错,但是运行会报错
	fmt.Println(c)

}

func seven() {
	fmt.Println("-------判断某个键是否存在")
	scoreMap := make(map[string]int, 20)
	scoreMap["李华"] = 100
	scoreMap["小明"] = 99
	value, ok := scoreMap["小名"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("无数据")
	}
}

func eight() {
	fmt.Println("-------map的遍历")
	//for range
	// map是无序的, 键值对和添加的顺序无关
	scoreMap := make(map[string]int, 20)
	scoreMap["李华"] = 100
	scoreMap["小明"] = 99
	scoreMap["小华"] = 97
	for k, v := range scoreMap {
		fmt.Println(k, "=", v)
	}
	// 只遍历map中的key
	for k := range scoreMap {
		fmt.Println("key=", k)
	}
	// 只遍历map中的value
	for _, v := range scoreMap {
		fmt.Println("key=", v)
	}
}

func ninth() {
	fmt.Println("-------使用delete()函数删除键值对")
	scoreMap := make(map[string]int, 20)
	scoreMap["李华"] = 100
	scoreMap["小明"] = 99
	scoreMap["小华"] = 97
	delete(scoreMap, "小华")
	fmt.Println(scoreMap)
}

func tenth() {
	fmt.Println("-------按照某个固定顺序遍历")
	scoreMap := make(map[string]int, 100)

	// 添加50个键值对
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("Student%02d", i)
		value := rand.Intn(100) // 0~99的随机数
		scoreMap[key] = value
	}
	//fmt.Println(scoreMap)

	// 按照key从小到大的顺序去遍历scroeMap
	// 1. 先取出所有key存放到切片中
	keys := make([]string, 0, 100)
	for k := range scoreMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("key=%v,value=%v\n", key, scoreMap[key])
		//fmt.Println(key, scoreMap[key])
	}
	fmt.Println(scoreMap)
}

func eleventh() {
	// 元素类型为map的切片
	mapSlice := make([]map[string]int, 8, 8) // 只是完成了切片的初始化
	// [nil nil nil nil nil nil nil nil]
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println(mapSlice[0] == nil)
	// 还需要完成内部map元素的初始化
	mapSlice[0] = make(map[string]int, 8) // 完成map的初始化
	mapSlice[0]["main"] = 100
	fmt.Println(mapSlice)
}

func twelfth() {
	fmt.Println("-------值为切片的map")
	sliceMap := make(map[string][]string, 3) // 只完成了map的初始化
	v, ok := sliceMap["中国"]
	if ok {
		fmt.Println(v)
	} else {
		// sliceMap中没有中国这个键
		sliceMap["中国"] = make([]string, 8, 8) // 完成了对切片的初始化
		sliceMap["中国"][0] = "100"
		sliceMap["中国"][1] = "200"
		sliceMap["中国"][2] = "300"
	}
	for k, v := range sliceMap {
		fmt.Println(k, v)
	}
}

func thirteenth() {
	// 统计一个字符串中每个单词出现的次数
	// "how do you do"中每个单词出现的次数
	// 0. 定义一个map[string]int
	s := "how do you do"
	wordCount := make(map[string]int, 10)
	// 1. 字符串中都有哪些单词
	words := strings.Split(s, " ")
	// 2.遍历单词做统计
	for _, word := range words {
		value, ok := wordCount[word]
		if ok {
			// map中有一个单词的统计记录
			wordCount[word] = value + 1
		} else {
			// map中没有这个单词的统计记录
			wordCount[word] = 1
		}
	}
	for k, v := range wordCount {
		fmt.Println(k, v)
	}
}

func main() {
	frist()
	second()
	third()
	fourth()
	fifth()
	seven()
	eight()
	ninth()
	tenth()
	eleventh()
	twelfth()
	thirteenth()
}
