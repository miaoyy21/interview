package main

import "fmt"

func main() {
	s0 := new(string)

	*s0 = "sss"
	fmt.Println(*s0)
}
