package main

import "fmt"

type Date struct {
	Year  int
	Month int
	Day   int
}

func (d *Date) SetYear(year int) {
	d.Year = year
}

func (d *Date) SetMonth(month int) {
	d.Month = month
}

func (d *Date) SetDay(day int) {
	d.Day = day
}

func main() {
	date := Date{}
	date.SetYear(0)
	date.SetMonth(14)
	date.SetDay(38)
	fmt.Println(date)
}
