package main

import "fmt"

func main() {
	a, b := func1(2, 3)
	fmt.Println(a, b)
	func2(a, b)
}

//함수에서 두개 이상의 인자를 한번에 리턴할 수 있다.
func func1(x, y int) (int, int) {
	return y, x
}

//함수는 함수를 호출할 수 있다.
func func2(x, y int) {
	func3(x, y)
}
func func3(x, y int) {
	fmt.Println("func2", x, y)
}
