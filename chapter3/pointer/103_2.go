package main

import "fmt"

func main() {
	var myInt int
	fmt.Println(&myInt)

	var myFloat float64
	fmt.Println(&myFloat)

	var myBool bool
	fmt.Println(&myBool)
}
