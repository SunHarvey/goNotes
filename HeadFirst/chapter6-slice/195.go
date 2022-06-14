package main

import "fmt"

func main() {
	fmt.Println(1)
	fmt.Println(1, 2, 3, 4, 5)
	letters := []string{"a"}
	letters = append(letters, "b")
	letters = append(letters, "c", "d", "e", "f", "g")
	fmt.Println(letters)
}
