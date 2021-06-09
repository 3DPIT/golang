package main

import "fmt"

type Gallons float64

func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

type Liters float64

func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}

type Milliliters float64

func String() {
	fmt.Println(13)
}

func main() {
	fmt.Println(Gallons(12.12345678))
	fmt.Println(Liters(12.12345678))
	//fmt.Println(Milliliters(12.12345678))
}
