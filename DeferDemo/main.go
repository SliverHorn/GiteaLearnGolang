package main

import "fmt"

// defer:延迟执行
func frist() {
	fmt.Println("Start.......")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("End.......")
}


func main() {
	frist()
}
