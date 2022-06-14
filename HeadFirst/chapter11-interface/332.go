package main

import "fmt"

type Switch string

func (s *Switch) toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
	fmt.Println(*s)
}

type Toggleable interface {
	toggle()
}

func main() {
	s := Switch("Off")
	var t Toggleable = &s
	t.toggle()
	t.toggle()
}
