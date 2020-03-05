package main

import "fmt"

func main() {
	var a = new(int)
	fmt.Println(*a)
	*a = 100
	fmt.Println(*a)

	var b = make(map[string]int, 10)
	fmt.Println(b)
	b["SliverHorn"] = 7
	fmt.Println(b)

}
