package main

import "fmt"

type Liters float64
type Gallons float64
type Milliliters float64

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}
func (g Gallons) ToMillilters() Milliliters {
	return Milliliters(g * 3785.41)
}
func main() {
	milk := Gallons(2)
	fmt.Printf("%0.3f gallons equals %0.3f liters\n", milk, milk.ToLiters())
	fmt.Printf("%0.3f gallons equals %0.3f millilters\n", milk, milk.ToMillilters())
}
