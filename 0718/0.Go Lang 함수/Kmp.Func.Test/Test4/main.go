//재귀 동작확인
package main

import "fmt"

func main() {
	f1(10)
}

func f1(x int) {
	if x == 0 {
		return
	}
	fmt.Printf("f1(%d) before call f1(%d)\n", x, x-1)
	f1(x - 1)
	fmt.Printf("f1(%d) after call f1(%d)\n", x, x-1)

}
