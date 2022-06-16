package main

import "fmt"

func main() {
	var v1, v2 bool
	v1 = true
	v3, v4 := false, true

	fmt.Printf("v1: %t\n", v1)
	fmt.Printf("v2: %t\n", v2)
	fmt.Printf("v3: %t\n", v3)
	fmt.Printf("v4: %t\n", v4)
}
