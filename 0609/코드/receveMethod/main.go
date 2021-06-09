package main

import "fmt"

type MyType string

func (m MyType) MethodWithParameters(number int, flag bool) {
	fmt.Println(m)
	fmt.Println(number)
	fmt.Println(flag)
}
func main() {
	value := MyType("myType value")
	value.MethodWithParameters(4, true)

}
