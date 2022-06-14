package main

import "fmt"

func main() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
//	index := 1
//	fmt.Println(index, notes[index])
//	index = 3
//	fmt.Println(index, notes[index])

	for i := 0; i <= 7; i++ {
		fmt.Println(i, notes[i])
	}
}
