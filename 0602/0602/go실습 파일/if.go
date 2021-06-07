package main

import (
	"fmt"
)

func main() {
	if true && false {
		fmt.Println("0을 반환 어짜피 여기 못들어옴")
	}
	if true || false {
		fmt.Println("1을 반환 그래서 여기 들어와서 출력")
	}

}
