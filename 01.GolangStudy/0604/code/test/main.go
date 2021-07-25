package main

import "fmt"

func severalStrings(strings ...string) {
	fmt.Println(strings)
}

func main() {
	severalStrings("a")
	severalStrings("a", "b", "c")
	severalStrings()
}
