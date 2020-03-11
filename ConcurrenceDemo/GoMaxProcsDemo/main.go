/**
 *
 * @Description:
 * @Author: SliverHorn
 * @File: main.go
 * @Version: 1.13.8
 * @Time:  2:11 下午
 * @Project: Go
 * @Package: main
 * @Software: GoLand
 */
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("A", i)
	}
	wg.Done()
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("B", i)
	}
	wg.Done()
}

func main() {
	runtime.GOMAXPROCS(2)
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
