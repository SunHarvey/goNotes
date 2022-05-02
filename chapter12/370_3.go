package main

import "fmt"

func calmDown() {
	recover()
}

func freakOut() {
	defer calmDown()
	panic("oh no")
}

func main() {
	freakOut()
	fmt.Println("Exiting normally")
}
