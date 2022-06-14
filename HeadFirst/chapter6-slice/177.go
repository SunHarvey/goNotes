package main

import "fmt"

func main() {
	notes := make([]int, 7)
	primes := make([]int, 5)
	fmt.Println(len(primes))
	fmt.Println(len(notes))

	letters := []string{"a", "b", "c"}
	for i := 0; i < len(letters); i++ {
		fmt.Println(letters[i])
	}
	for _, letter := range letters {
		fmt.Println(letter)
	}
}
