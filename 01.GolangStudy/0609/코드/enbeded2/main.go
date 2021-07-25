package main

import (
	"fmt"
	"log"

	"src/github.com/headfirstgo/calendar"
)

func main() {
	event := calendar.Event{}
	err := event.SetTitle("dfsfsfsfsdddddddddddddddddddddddddddddddfsfs")
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetYear(2021)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Title()) //Event로 부터 승격
	fmt.Println(event.Year())  //Data로 부터 승격
}
