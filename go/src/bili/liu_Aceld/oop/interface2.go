package main

import "fmt"

func myFunc(arg interface{}) {
	fmt.Println("myFunc is called..")
	fmt.Println(arg)
}

type Book struct {
	auth string
}

func main() {
	book := Book{"Go"}
	myFunc(book)

	myFunc(100)
	myFunc("abc")
	myFunc(3.14)
}
