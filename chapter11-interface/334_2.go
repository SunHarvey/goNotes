package main

import "github.com/headfirstgo/gadget"

type Player interface {
	Play(string)
	Stop()
}

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recorder := gadget.TapeRecorder(player)
	recorder.Record()
}

func main() {
	TryOut(gadget.TapeRecorder{})
}
