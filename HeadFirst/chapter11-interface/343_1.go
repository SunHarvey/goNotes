package main

import "fmt"

type CoffeePot string

func (c CoffeePot) String() string {
	return string(c) + " coffee pot"
}

func main() {
	coffeePot := CoffeePot("LuxBrew")
	fmt.Print(coffeePot, "\n")
	fmt.Println(coffeePot.String())
	fmt.Printf("%s", coffeePot)
}
