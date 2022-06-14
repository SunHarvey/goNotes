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
	recorder := player.(gadget.TapeRecorder)
	fmt.Println(recorder)
}
