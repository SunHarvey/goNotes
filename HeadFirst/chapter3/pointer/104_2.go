package main

import "fmt"

func main() {
	var myInt int
	var myIntPointer *int
	myIntPointer = &myInt
	fmt.Println(myIntPointer)

	var myFloat float64
	var myFloatPointer *float64
	myFloatPointer = &myFloat
	fmt.Println(myFloatPointer)

	var myBool bool
	myBoolPointer := &myBool
	fmt.Println(myBoolPointer)
}
