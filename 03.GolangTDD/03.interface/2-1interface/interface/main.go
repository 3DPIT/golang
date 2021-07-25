package main

import (
	"fmt"
	"strconv"
)

type StructA struct {
	val string
}

func (s *StructA) String() string {
	return "Val: " + s.val
}

type Printable interface {
	String() string
}

func Println(p Printable) {
	fmt.Println(p.String())
}

type StructB struct {
	val int
}

func (s *StructB) String() string {
	return "StructB:" + strconv.Itoa(s.val)
}

func main() {
	a := &StructA{val: "AAA"}
	fmt.Println(a)

	b := &StructB{val: 190}
	fmt.Println(b)
}
