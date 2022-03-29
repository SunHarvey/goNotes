package main

import "fmt"

func main() {
	counts := make(map[string]int)
	counts["Amber Graham"]++
	counts["Brian Martin"]++
	counts["Amber Graham"]++
	fmt.Println(counts)
}
