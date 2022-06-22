package main

import (
	"chatserver/client"
	"chatserver/gui"
	"flag"
	//"github.com/gpmgo/gopm/modules/log"
	"log"
)

func main() {
	address := flag.String("server", "127.0.0.1:3333", "address of server")
	flag.Parse()
	client := client.NewClient()
	err := client.Dial(*address)

	if err != nil {
		log.Fatal("Error when connect to server", err)
	}

	defer client.Close()

	go client.Start()
	go StartUi(client)
}
