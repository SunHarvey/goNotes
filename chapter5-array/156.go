package main

import "fmt"

func main() {

	//	primes := [5]int{2, 3, 5, 7, 11}
	//	fmt.Println(len(primes))
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(len(notes))

	for i := 0; i <= len(notes); i++ {
		fmt.Println(i, notes[i])
	}
}
