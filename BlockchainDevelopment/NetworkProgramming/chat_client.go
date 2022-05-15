package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"unsafe"
)

type ClientMsg struct {
	To      string  `json:"to"`
	Msg     string  `json:"msg"`
	Datalen uintptr `json:"datalen"`
}

func Help() {
	fmt.Println("1. set:your name")
	fmt.Println("2. all:your msg -- broadcadt")
	fmt.Println("1. anyone:your msg -- private msg")
}

func handle_conn(conn net.Conn) {
	buf := make([]byte, 256)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Panic("Failed to Read", err)
		}

		fmt.Println(string(buf[:n]))
		fmt.Printf("pdj's chat >")
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		log.Panic("Failed to Dial", err)
	}

	defer conn.Close()
	go handle_conn(conn)
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("welcome to pdj's pub chat\n")

	Help()
	for {
		fmt.Printf("pdj's chat>")
		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Panic("Failed to ReadString", err)
		}
		msg = strings.Trim(msg, "\r\n")
		if msg == "quit" {
			fmt.Println("bye bye ")
			break
		}

		if msg == "help" {
			Help()
			continue
		}
		msgs := strings.Split(msg, ":")
		if len(msgs) == 2 {
			var climsg ClientMsg
			climsg.To = msgs[0]
			climsg.Msg = msgs[1]
			climsg.Datalen = unsafe.Sizeof(climsg)
			data, err := json.Marshal(climsg)
			if err != nil {
				fmt.Println("Failed to Marshal", err, climsg)
				continue
			}

			_, err = conn.Write(data)
			if err != nil {
				fmt.Println("Failed to write", err, climsg)
				break
			}
		}
	}
}
