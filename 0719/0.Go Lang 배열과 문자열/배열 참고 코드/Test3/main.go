package main

import "fmt"

func main() {
	s := "Hello 월드"
	fmt.Println("바이트 일경우 len(s2) = ", len(s))
	for i := 0; i < len(s); i++ {
		//fmt.Print(s2[i], ", ")
		fmt.Print(string(s[i]), ", ")
	}
	fmt.Println()
	s2 := []rune(s)
	fmt.Println("룬 일경우 len(s2) = ", len(s2))
	for i := 0; i < len(s2); i++ {
		//fmt.Print(s2[i], ", ")
		fmt.Print(string(s2[i]), ", ")
	}
}
