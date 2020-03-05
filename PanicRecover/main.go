package main

import "fmt"

// panoc和recover
func a() {
	fmt.Println("fun a")
}

func b() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("func b defter")
		}
	}()
	panic("panic in b")
}

func c() {
	fmt.Println("fun c")
}

func main() {
	a()
	b()
	c()

}
