package main

import "fmt"

func twoInts(first int, second int) {
	fmt.Println(first, second)
}

func main() {
	twoInts(1)
	twoInts(1, 2, 3)
}
