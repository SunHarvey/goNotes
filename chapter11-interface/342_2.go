package main

import "fmt"
/*
type Stringer interface {
	String() string
}
*/
type CoffeePot string

func (c CoffeePot) String() string {
	return string(c) + " coffee pot"
}

func main() {
	coffeePot := CoffeePot("LuxBrew")
	fmt.Println(coffeePot.String())
}
