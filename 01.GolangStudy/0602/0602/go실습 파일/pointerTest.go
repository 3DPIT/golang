package main

import (
	"fmt"
)

func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat
}
func main() {
	var myInt int
	var myIntPointer *int
	myIntPointer = &myInt
	fmt.Println(myIntPointer)
	var myFloatPointer *float64 = createPointer()
	fmt.Println(*myFloatPointer)
}
