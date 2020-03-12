/**
 *
 * @Description: // channel
 * @Author: SliverHorn
 * @File: main.go
 * @Version: 1.13.8
 * @Time:  3:20 下午
 * @Project: Go
 * @Package: ChannelDemo1
 * @Software: GoLand
 */
package main

import "fmt"

func main() {
	// var ch1 chan int // 引用类型, 需要初始化之后才能使用
	//ch2 := make(chan int, 1) // 带缓冲区通道

	//ch1 := make(chan int)    // 无缓冲区, 又称同步通道
	ch1 := make(chan int, 1) // 带缓冲区通道
	ch1 <- 17                // 发送值
	x := <-ch1
	fmt.Println(x)
	ch1len := len(ch1)
	fmt.Println(ch1)
	// len():取通道中元素的数量
	// len():拿到通道的容量
	fmt.Printf("Value=%v, Cap=%v", len(ch1), cap(ch1))
	fmt.Println(ch1len)
	close(ch1)
}
