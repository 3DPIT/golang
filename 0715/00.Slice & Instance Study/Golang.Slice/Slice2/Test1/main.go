//슬라이스 자르기
package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[4:8]
	fmt.Println(b)
	b = a[4:] //시작위치만 지정
	fmt.Println(b)
	b = a[:4]
	fmt.Println(b)

}
