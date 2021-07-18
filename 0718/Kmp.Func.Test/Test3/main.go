package main

import "fmt"

func main() {
	f1(10)
}

func f1(x int) {
	fmt.Println(x)
	if x == 0 {
		return
	}
	f1(x - 1)
}
