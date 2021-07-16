package main

import (
	"fmt"
)

type StructA struct {
}

func (a *StructA) AAA(x int) int {
	return x * x
}

func main() {
	var c InterfaceA
	c = &StructA{}
	fmt.Println(c.BBB(3))
}
