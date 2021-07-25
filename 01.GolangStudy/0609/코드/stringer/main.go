package main

import "fmt"

type CoffeePot string

func (c CoffeePot) String() string {
	return string(c) + " coffe pot"
}
func main() {
	CoffeePot := CoffeePot("LuxBrew")
	fmt.Print(CoffeePot.String(), "\n")
	fmt.Println(CoffeePot.String())
	fmt.Printf("%s", CoffeePot.String())

}
