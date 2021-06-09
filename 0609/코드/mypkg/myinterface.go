package mypkg

import "fmt"

type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameters(float64)
	MethodWithReturnValue() string
}
type MyType int

func (m MyType) MethodWithoutParameters() {
	fmt.Println("MethodWithoutParameters called")
}
func (m MyType) MethodWithParameters(f float64) { //매게변수를 받는 경우
	fmt.Println("MethodWithoutParameters called with", f)
}
func (m MyType) MethodWithReturnValue() string { //반환값 받는 경우
	return "Hi method return "
}
func (my MyType) MethodNotInInterface() { // 인터페이스와 무관한 것
	fmt.Println("MethodNotInInterface called")
}
