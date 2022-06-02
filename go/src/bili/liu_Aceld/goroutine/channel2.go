package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)
	fmt.Println("len(c) = ", len(c), ",cap(c)", cap(c))

	go func() {
		defer fmt.Println("sub goroutine end")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("sub goroutine is running,send = ", i, "len(c) = ", len(c), ",cap(c)=", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("num = ", num)
	}
	fmt.Println("main end")
}
