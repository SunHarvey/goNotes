package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)
	ch <- 10
	ch <- 11
	fmt.Println("send successful")
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
