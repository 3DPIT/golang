package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var p Person
	p1 := Person{"홍길동", 15}
	p2 := Person{name: "길동홍", age: 21}
	p3 := Person{name: "동길홍"}
	p4 := Person{}
	fmt.Println(p, p1, p2, p3, p4)
	p.name = "Smith"
	p.age = 25
	fmt.Println(p)
}
