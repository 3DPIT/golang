package main

import "fmt"

type Number int

func (n *Number) Double() {
	*n *= 2
}

func main() {
	number := Number(4)
	fmt.Println("원래 숫자:", number)
	number.Double()
	fmt.Println("두배된 함수 들어간 숫자:", number)
}
