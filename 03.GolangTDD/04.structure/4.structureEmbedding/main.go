//구조체 임베딩
package main

import (
	"fmt"
)

type Student struct {
	name string
	age  int
	Address
}
type Teacher struct {
	name   string
	age    int
	salary float64
	Address
}
type Address struct {
	street     string
	city       string
	state      string
	postalCode string
}

func main() {
	s := Student{name: "홍길동", age: 20}
	t := Teacher{name: "선생님", age: 40}
	s.street = "24번가"
	s.city = "수원"
	s.state = "영통구"
	s.postalCode = "123123"
	t.city = "수원"
	fmt.Println("학생 :", s, "선생님 :", t)
}
