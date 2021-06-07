package main

import (
	"fmt"
)

func main() {
	var number int = 1
	var aString string = "aaa"
	var cString string = "ccc"
	var lan = append([]string{}, "english")
	fmt.Println(number, aString, "on", cString, lan)
}
