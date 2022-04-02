package main

import "fmt"

func main() {
	var subscriber1 struct {
		name   string
		rate   float64
		active bool
	}

	subscriber1.name = "Aman Singh"
	fmt.Println("Name:", subscriber1.name)
	var subscriber2 struct {
		name   string
		rate   float64
		active bool
	}
	subscriber2.name = "Beth Ryan"
	fmt.Println("Name:", subscriber2.name)
}
