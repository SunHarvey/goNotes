package main

import (
	"fmt"
	"github.com/headfirstgo/calendar"
	"log"
)

func main() {
	event := calendar.Event{}
	err := event.SetTitle("An extremely long title is impractical to print")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Title)
}
