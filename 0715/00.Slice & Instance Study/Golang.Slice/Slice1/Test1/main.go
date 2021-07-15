//선언과 length, capacity 확인
package main

import "fmt"

func main() {
	//var a []int
	//a := []int{1, 2, 3, 4, 5}
	a := make([]int, 0, 8)
	fmt.Printf("len(a) = %d\n", len(a))
	fmt.Printf("cap(b) = %d\n", cap(a))
}
