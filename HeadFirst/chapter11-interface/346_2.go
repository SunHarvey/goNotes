package main

import "fmt"

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

func AcceptWhistle(whistle Whistle) {
	fmt.Println(whistle)
	whistle.MakeSound()
}

func main() {
	//AcceptWhistle(3.1415)
	AcceptWhistle(Whistle("Toyco Canary"))
}
