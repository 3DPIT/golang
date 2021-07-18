package main

import "fmt"

func main() {
	s := 0
	for i := 10; i >= 0; {
		s += i
		i--
	}
	fmt.Println(s)
}
