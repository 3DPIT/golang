package main

import "fmt"

func main() {
	one()
}
func one() {
	defer fmt.Println("deferred in one()")
	two()
}
func two() {
	defer fmt.Println("defferred in two()")
	panic("지연된 호출은 크래시 발생하기전 실행 됨.")
}
