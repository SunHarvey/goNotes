package main

import "fmt"

func main() {
	underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	slice1 := underlyingArray[0:3]
	fmt.Println(slice1)

	i, j := 1, 4
	slice2 := underlyingArray[i:j]
	fmt.Println(slice2)

}
