//값 주소 전달 테스트
package main

import "fmt"

type Student struct {
	name  string
	age   int
	grade int
}

// func SetName(t Student, newName string) {////값 타입 assgin(=)

// 	t.name = newName
// }
func SetName(t *Student, newName string) { ////주소 타입 assgin(=)

	t.name = newName
}
func main() {
	a := Student{"aaa", 20, 10}
	SetName(&a, "bbb")
	fmt.Println(a)
}
