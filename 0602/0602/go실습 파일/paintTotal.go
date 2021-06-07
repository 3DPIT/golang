package main

import (
	"fmt"
)

func paintUse(width float64, height float64) float64 {
	area := width * height
	fmt.Printf("총 필요한것은 양은 %.2f\n", area/10.0)
	return area / 10.0
}

func main() {
	total := paintUse(1.5, 3.0)
	total += paintUse(3.5, 5.0)
	fmt.Printf("총 페인트량은 %.2f\n", total)

}
