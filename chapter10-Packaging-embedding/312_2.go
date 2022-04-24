package main

import (
	"fmt"
	"github.com/headfirstgo/calendar"
)

func main() {
	event := calendar.Event{}
	event.Title = "An extremely long title is impractical to print"
	fmt.Println(event.Title)
}
