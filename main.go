package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type A struct {
	Id   string
	Name string
}

type B struct {
	Id   string
	Name string
}

func main() {
	a := "aaabbbccc"
	ssh := *(*reflect.StringHeader)(unsafe.Pointer(&a))
	fmt.Printf("%#v\n", ssh)
	b := *(*[]byte)(unsafe.Pointer(&ssh))
	fmt.Printf("%s\n", b)

	s1 := A{Id: "1", Name: "A"}
	s2 := *(*B)(unsafe.Pointer(&s1))
	fmt.Printf("%#v", s2)

	aaa := new([]string)
	fmt.Println(aaa == nil)
}
