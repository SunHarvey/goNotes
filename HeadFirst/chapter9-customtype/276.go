package main

import "fmt"

type MyType string

func (m MyType) sayHi() {
	fmt.Println("Hi from", m)
}

func main() {
	value := MyType("a MyType value")
	value.sayHi()
	anotherValue := MyType("another value")
	anotherValue.sayHi()
}
