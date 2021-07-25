package main

import "fmt"

var a = "a"

func main() {
	a = "a"
	b := "b"
	if true {
		c := "c"
		if true {
			d := "d"
			fmt.Println(a)
			fmt.Println(b)
			fmt.Println(c)
			fmt.Println(d)
		}
		fmt.Println(a)
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
	}
	fmt.Println(a)
	fmt.Println(b)
}
