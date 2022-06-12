package main

import (
	"fmt"
	"time"
)

func hello(c chan int) {
	for i := 0; i < 20; i++ {
		//fmt.Println("hello ", i)
		c <- i
	}
}

func main() {
	ch := make(chan int, 20)
	go hello(ch)

	go func() {
		for {
			i, ok := <-ch
			if !ok {
				break
			} else if i%4 == 0 {
				fmt.Println(i)
			}
		}
		close(ch)
	}()
	time.Sleep(1 * time.Second)
}
