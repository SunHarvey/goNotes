package main

import (
	"fmt"
	"sync"
	"time"
)

var c chan string
var w sync.WaitGroup

func reader() {
	msg := <-c
	fmt.Println("I am reader", msg)
}

func main() {
	c = make(chan string)
	w.Add(1)
	go reader()
	fmt.Println("Begin sleep")
	time.Sleep(time.Second * 3)
	c <- "hello"
	time.Sleep(time.Second * 1)
}
