package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	   go func() {
	       defer fmt.Println("A.defer")
	       func() {
	           defer fmt.Println("B.defer")
	           fmt.Println("B")
	       }()
	       fmt.Println("A")
	   }()
	*/
	go func(a int, b int) bool {
		fmt.Println(a, b)
		return true
	}(10, 20)

	for {
		time.Sleep(1 * time.Second)
	}
}
