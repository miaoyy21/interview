package main

import "fmt"

func fn() (a int) {
	defer func() {
		a = 2
		fmt.Println("9999")
	}()

	a = 1
	return
}

func main() {
	fmt.Println(fn())
}
