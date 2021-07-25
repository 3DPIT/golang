package main

import (
	"fmt"
	"log"

	"src/github.com/headfirstgo/calendar"
)

func main() {
	event := calendar.Event{}
	err := event.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(10)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(27)
	if err != nil {
		log.Fatal(27)
	}

	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())

	fmt.Println(event.Data.Year())
	fmt.Println(event.Data.Month())
	fmt.Println(event.Data.Day())

}
