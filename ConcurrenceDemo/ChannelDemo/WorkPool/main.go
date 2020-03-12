/**
 *
 * @Description: work_pool
 * @Author: SliverHorn
 * @File: main.go
 * @Version: 1.13.8
 * @Time:  10:54 上午
 * @Project: Go
 * @Package: WorkPool
 * @Software: GoLand
 */
package main

import "fmt"

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, job)
		results <- job * 2
		fmt.Printf("worker:%d end job:%d\n", id, job)
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 开启3个goroutine
	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}

	// 发送5个任务
	for i := 0; i < 5; i++ {
		jobs <- i
	}

	close(jobs)
	// 输出结果
	for k := 0; k < 5; k++ {
		ret := <-results
		fmt.Println(ret)
	}

}
