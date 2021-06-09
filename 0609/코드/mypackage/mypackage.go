package mypackage //아래 타입은 별도의 패키지에 속함

import "fmt"

type MyType struct {
	EmbeddedType //임베딩
}
type EmbeddedType string //임베딩할 타입을 선언

func (e EmbeddedType) ExportedMethod() { // 이 메서드는 MyType으로 승격
	fmt.Println("Hi from ExportedMethod on EmbeddedType")
}
func (e EmbeddedType) unexportedMethod() { //이 메서드는 승격 X

}
