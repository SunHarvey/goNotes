package main

import (
	"fmt"
	"github.com/headfirstgo/calendar"
)

func main() {
	event := calendar.Event{}
	event.Title = "Mom's birthday"
	fmt.Println(event.Title)
}
