package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "\t trim check string \n"
	fmt.Println(s)
	fmt.Println(strings.TrimSpace(s))
}
