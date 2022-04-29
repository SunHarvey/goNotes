package main

import (
	"fmt"
	"github.com/headfirstgo/gadget"
)

type Player interface {
	Play(string)
	Stop()
}

func main() {
	var player Player = gadget.TapePlayer{}
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Player was not a TapeRecorder")
	}
}
