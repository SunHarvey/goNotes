package main

import "fmt"

func recurses() {
	fmt.Println("Oh,no, i'm stuck!")
	recurses()
}

func main() {
	recurses()
}
