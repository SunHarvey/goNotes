package main

import "fmt"

func myFunc(arg interface{}) {
	fmt.Println("myFunc is called..")
	fmt.Println(arg)

	value, ok := arg.(string)
	if !ok{
		fmt.Println("arg is not string")
	} else {
		fmt.Println("arg is string, value is ", value)
		fmt.Printf("value type is %T\n", value)
	}
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
