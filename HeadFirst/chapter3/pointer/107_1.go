package main

import "fmt"

func main() {
	var myFloatPointer *float64 = createPointer()
	fmt.Println(*myFloatPointer)
}

func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat
}
