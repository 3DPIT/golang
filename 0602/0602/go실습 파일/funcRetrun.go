package main

import "fmt"

func double(number float64) int {
	return int(number * 2)
}
func stringReturn(score float64) string {
	if score >= 60 {
		return "passing"
	}
	return "failing"
}

func main() {
	dozen := double(2.5)
	fmt.Println(dozen)
	fmt.Println(double(1.5))
	fmt.Println(stringReturn(61))
}
