package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) PrintName() {
	fmt.Print(p.name)
}
func main() {
	var p Person
	p.name = "홍길동"
	p.PrintName()
}
