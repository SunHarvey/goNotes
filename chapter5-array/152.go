package main

import "fmt"

func main() {
	var primes [5]int
	primes[0] = 2
	fmt.Println(primes[0])
	fmt.Println(primes[2])
	fmt.Println(primes[4])

	var notes [7]string
	notes[0] = "do"
	fmt.Println(notes[3])
	fmt.Println(notes[6])
	fmt.Println(notes[0])

	var counters [3]int
	counters[0]++
	counters[0]++
	counters[2]++
	fmt.Println(counters[0])
	fmt.Println(counters[1])
	fmt.Println(counters[2])
}
