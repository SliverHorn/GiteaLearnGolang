/**
 *
 * @Description: go routine demo1
 * @Author: SliverHorn
 * @File: main.go
 * @Version: 1.13.8
 * @Time:  1:12 下午
 * @Project: Go
 * @Package: main
 * @Software: GoLand
 */
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("Hello", i)
	wg.Done() // 通知wg计数器减一
}

func main() { // 开启一个主goroutine去执行main函数
	for i := 0; i < 1000; i++ {
		wg.Add(1)   // 计数牌加一
		go hello(i) // 开启了一个goroutine去执行hello这个函数
	}
	fmt.Println("Hello Main")
	//time.Sleep(time.Second)
	wg.Wait() // 阻塞, 等所有小弟都干完活之后才收兵
}
