package main

import "fmt"

type Liters float64
type Milliliters float64
type Gallons float64

func (l Liters) ToMillilters() Milliliters {
	return Milliliters(l * 1000)
}
func (m Milliliters) ToLiters() Liters {
	return Liters(m / 1000)
}

func main() {
	l := Liters(3)
	fmt.Printf("%0.3f liters is %0.1f milliliters\n", l, l.ToMillilters())
	ml := Milliliters(500)
	fmt.Printf("%0.3f Milliliters is %0.1f Liters\n", ml, ml.ToLiters())

}
