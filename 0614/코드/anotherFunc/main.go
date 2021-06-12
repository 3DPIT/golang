package main

import (
	"fmt"
)

func sayHi() {
	fmt.Println("Hi")
}

func sayBye() {
	fmt.Println("Bye")
}

func twice(theFunction func()) {
	theFunction()
	theFunction()
}

func main() {
	twice(sayHi)
	twice(sayBye)
	// http.HandleFunc("/hello",sayHi)
	// err := http.ListenAndServe("localhost:8080",nil)
	// log.Fatal(err)
}
