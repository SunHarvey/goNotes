package main

import (
	"fmt"
	"github.com/headfirstgo/calendar"
)

func main() {
	date := calendar.Date{}
	date.year = 2022
	date.month = 14
	date.day = 50
	fmt.Println(date)
	date = calendar.Date{year: 0, month: 0, day: -2}
	fmt.Println(date)

}
