package main

import (
	"fmt"
	"time"
)

func greeting(myChannel1 chan string) {
	time.Sleep(5 * time.Second)
	myChannel1 <- "hi"
}

func main() {
	myChannel2 := make(chan string)
	//go greeting(myChannel2)
	//fmt.Println(<-myChannel2)
	go greeting(myChannel2)
	abc := <-myChannel2
	fmt.Println(abc)
}
