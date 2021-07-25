package main

import "fmt"

func sayHi() {
	fmt.Println("Hi")
}
func main() {
	var myFunction func()
	myFunction = sayHi
	myFunction()
}
