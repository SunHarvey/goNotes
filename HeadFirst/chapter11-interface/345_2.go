package main

import "fmt"

type Whistle string

func AcceptAnything(thing interface{}) {
	fmt.Println(thing)
	thing.Makesound()
}

func main() {
	AcceptAnything(3.1415)
	AcceptAnything(Whistle("Toyco Canary"))
}
