package main

import "fmt"

type Mytype string

func (m Mytype) sayHi() {
	fmt.Println("Hi from", m)
}
func main() {
	value := Mytype("a MyType value")
	value.sayHi()
	anotherValue := Mytype("another value")
	anotherValue.sayHi()
}
