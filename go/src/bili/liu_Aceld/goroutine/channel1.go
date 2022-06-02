package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine stop")
		fmt.Println("goroutine running...")

		c <- 6
	}()

	num := <-c
	fmt.Println("num = ", num)
	fmt.Println("main goroutine end...")
}
