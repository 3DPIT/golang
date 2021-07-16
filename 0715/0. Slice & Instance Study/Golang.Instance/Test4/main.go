//인스턴스 경우
package main

import "fmt"

type Student struct {
	name  string
	age   int
	grade int
}

func (t *Student) SetName(newName string) { ////주소 타입 assgin(=)

	t.name = newName
}
func main() {
	a := Student{"aaa", 20, 10}
	a.SetName("bbb") //여기서a가 instance 즉 주체가된다. 생명을 가짐
	fmt.Println(a)
}
