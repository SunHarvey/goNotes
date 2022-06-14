package main

import "fmt"

func main() {
	ranks := make(map[string]int)
	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3
	fmt.Println(ranks["bronze"])
	fmt.Println(ranks["gold"])

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["Li"] = "Lithium"
	fmt.Println(elements["Li"])
	fmt.Println(elements["H"])

	isPrime := make(map[int]bool)
	isPrime[4] = false
	isPrime[7] = true
	fmt.Println(isPrime[4])
	fmt.Println(isPrime[7])
}
