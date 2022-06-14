package main

import (
	"fmt"
	"github.com/headfirstgo/calendar"
	"log"
)

func main() {
	date := calendar.Date{}
	err := date.SetYear(0)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(4)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
}
