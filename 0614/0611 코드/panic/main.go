package main

import "fmt"

func calmDown() {
	p := recover()
	err, ok := p.(error)
	if ok {
		fmt.Println(err.Error())
	} else {
		panic(p)
	}
}
func a() int {
	a1 := 1
	a1++
	return a1
}

func main() {
	defer calmDown()
	//a()
	err := fmt.Errorf("여기 에러다")
	panic(err)
}
