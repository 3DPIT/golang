package main

import "fmt"

func recurses() {
	fmt.Println("oh, no, I'm stuck")
	recurses()
}

func main() {
	recurses()
}
