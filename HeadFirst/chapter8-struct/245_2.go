package main

import "fmt"

type myStruct struct {
	myField int
}

func main() {
	var value myStruct
	var pointer *myStruct = &value
	pointer.myField = 9
	fmt.Println(pointer.myField)
}
