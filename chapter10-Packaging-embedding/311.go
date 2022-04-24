package main

import (
	"fmt"
	"github.com/headfirstgo/calendar"
	"log"
)

func main() {
	event := calendar.Event{}
	err := event.SetYear(2022)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetMonth(4)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetDay(24)
	if err != nil {
		log.Fatal(err)
	}

/*
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())
	*/
	fmt.Println(event.Date.Year())
	fmt.Println(event.Date.Month())
	fmt.Println(event.Date.Day())
}
