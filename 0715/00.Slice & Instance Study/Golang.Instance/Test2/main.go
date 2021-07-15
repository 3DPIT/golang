//주소 타입 assgin(=)
package main

import "fmt"

type Student struct {
	name  string
	age   int
	grade int
}

func main() {
	a := Student{"aaa", 20, 10}
	var b *Student
	b = &a
	b.age = 30

	fmt.Println(a)
	fmt.Println(b)
}
