package main

import (
	"fmt"
	//	"time"
)

func main() {
	fmt.Println(getNumber())
}

func getNumber() int {
	var i int
	go func() {
		i = 5
	}()
	//	time.Sleep(1 * time.Second)
	return i
}
