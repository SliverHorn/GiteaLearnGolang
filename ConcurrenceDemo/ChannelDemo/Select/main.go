/**
 *
 * @Description:
 * @Author: SliverHorn
 * @File: main.go
 * @Version: 1.13.8
 * @Time:  1:33 下午
 * @Project: Go
 * @Package: Select
 * @Software: GoLand
 */
package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		default:
			fmt.Println("啥都没干")
		}
	}
}
