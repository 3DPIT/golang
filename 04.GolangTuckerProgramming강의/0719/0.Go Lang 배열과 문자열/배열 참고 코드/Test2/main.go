package main

import "fmt"

func main() {
	s := "Hello World"
	for i := 0; i < len(s); i++ {
		//fmt.Print(s[i], ", ")
		fmt.Print(string(s[i]), ", ")
	}
}
