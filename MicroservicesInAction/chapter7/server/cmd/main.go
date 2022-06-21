package main

import (
	"chatserver/server"
)

func main() {
	var s server.Server
	s = server.NewServer()
	s.Listen(":333")
	s.Start()
}
