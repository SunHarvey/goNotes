package main

import "fmt"

type Whistle string 

func AcceptAnything(thing interface{}) {
	fmt.Println(thing)
}

func main() {
	AcceptAnything(3.1415)
	AcceptAnything(Whistle("Toyco Canary"))
}

