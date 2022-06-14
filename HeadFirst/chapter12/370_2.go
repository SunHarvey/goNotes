package main

import "fmt"

func freakOut() {
	panic("oh go")
	recover()
}

func main() {
	freakOut()
	fmt.Println("Exiting normally")
}
