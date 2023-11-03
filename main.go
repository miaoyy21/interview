package main

import "fmt"

const (
	i = 7 + iota
	j
	k
)

func main() {
	fmt.Println(i, j, k)
}
