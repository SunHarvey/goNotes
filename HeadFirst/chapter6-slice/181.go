package main

import "fmt"

func main() {
	underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	slice5 := underlyingArray[1:]
	fmt.Println(slice5)
}
