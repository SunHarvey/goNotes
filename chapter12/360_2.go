package main

import "fmt"

func count(start int, end int) {
	fmt.Println(start)
	if start < end {
		count(start+1, end)
	}
}

func main() {
	count(1, 3)
}
