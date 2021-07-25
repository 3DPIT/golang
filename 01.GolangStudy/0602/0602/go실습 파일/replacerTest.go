package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# r#ocks!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}
