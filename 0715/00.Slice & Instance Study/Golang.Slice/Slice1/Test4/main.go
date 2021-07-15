//같은 메모리 슬라이스일때 주의
package main

import "fmt"

func main() {
	a := make([]int, 2, 4)
	b := append(a, 3)
	fmt.Printf("%p %p\n", a, b)

	fmt.Println(a)
	fmt.Println(b)

	b[0] = 1
	b[1] = 2
	fmt.Println(a)
	fmt.Println(b)
}
