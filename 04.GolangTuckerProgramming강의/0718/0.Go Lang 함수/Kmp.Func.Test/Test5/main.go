//1~10의 합 재귀호출  
package main

import (
	"fmt"
)

func main() {
	s := sum(10, 0)
	fmt.Println(s)
}
func sum(x, s int) int {
	if x == 0 {
		return s
	}
	s += x
	return sum(x-1, s)
}
