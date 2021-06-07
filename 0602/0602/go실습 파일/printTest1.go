package main

import "fmt"

func paint(width float64, height float64) {
	area := width * height
	fmt.Printf("%.2f liters needed\n", area/10.0)
}

func main() {
	paint(4.2, 3.0)
	paint(1.54, 3.0)
}
