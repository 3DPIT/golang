//구조체 임베딩
package main

import (
	"fmt"
)

type Student struct {
	name    string
	age     int
	address Address
}
type Teacher struct {
	name    string
	age     int
	salary  float64
	address Address
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
	s.address.street = "24번가"
	s.address.city = "수원"
	s.address.state = "영통구"
	s.address.postalCode = "123123"
	t.address.city = "수원"
	fmt.Println("학생 :", s, "선생님 :", t)
}
