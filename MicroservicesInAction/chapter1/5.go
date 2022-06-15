package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var i, j, k int
	fmt.Printf("i:%d,j:%d,k:%d", i, j, k)
	fmt.Println()
	var a, b, c = 1, "s", true
	fmt.Printf("a:%d,b:%s,c:%t", a, b, c)
	fmt.Println()

	f := rand.Float64() * 100
	fmt.Printf("f:%g", f)
}
