package main

import "fmt"

func Socialize1() {
	defer fmt.Println("Goodbye!")
	fmt.Println("Hello!")
	fmt.Println("Nice Weather, eh?")
}

func main1() {
	Socialize1()
}
