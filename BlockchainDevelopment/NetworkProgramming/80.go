package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"unsafe"
)

type ChatMsg struct {
	From, To, Msg string
}

type ClientMsg struct {
	To      string  `json:"to"`
	Msg     string  `json:"msg"`
	Datalen uintptr `json:"datalen"`
}

var chan_msgcenter chan ChatMsg
var mapName2CliAddr map[string]string
var mapCliaddr2Clients map[string]net.Conn

func main() {
	mapCliaddr2Clients = make(map[string]net.Conn)
	mapName2CliAddr = make(map[string]string)
	chan_msgcenter = make(chan ChatMsg)

	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Panic("Ftiled to Listen", err)
	}

	defer listener.Close()

	go msg_center()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to Accept", err)
			break
		}
		go handle_conn(conn)
	}
}

func logout(conn net.Conn, from string) {
	defer conn.Close()
	delete(mapCliaddr2Clients, from)
	msg := ChatMsg{from, "all", from + "->logout"}
	chan_msgcenter <- msg
}

func handle_conn(conn net.Conn) {
	from := conn.RemoteAddr().String()
	mapCliaddr2Clients[from] = conn
	msg := ChatMsg{from, "all", from + "->login"}
	chan_msgcenter <- msg

	defer logout(conn, from)
	buf := make([]byte, 256)
	for {
		n, err := conn.Read(buf)
		fmt.Println(n, err)
		if err != nil {
			fmt.Println("Failed to Read", err, from)
			break
		}
		if n > 0 {
			var climsg ClientMsg
			err = json.Unmarshal(buf[:n], &climsg)
			if err != nil {
				fmt.Println("Failed to Unmarshal", err, string(buf[:n]))
				continue
			}
			if climsg.Datalen != unsafe.Sizeof(climsg) {
				fmt.Println("Msg format err:", climsg)
				continue
			}

			chatmsg := ChatMsg{from, "all", climsg.Msg}
			fmt.Println(climsg)
			switch climsg.To {
			case "all":
			case "set":
				mapName2CliAddr[climsg.Msg] = from
				chatmsg.Msg = from + " set name=" + climsg.Msg + " sucess"
				chatmsg.From = "server"
			default:
				chatmsg.To = climsg.To
			}
			chan_msgcenter <- chatmsg
		}
	}
}

func msg_center() {
	for {
		msg := <-chan_msgcenter
		go send_msg(msg)
	}
}

func send_msg(msg ChatMsg) {
	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Failed to Marshal", err, msg)
		return
	}
	if msg.To == "all" {
		for _, v := range mapCliaddr2Clients {
			if msg.From != v.RemoteAddr().String() {
				v.Write(data)
			}
		}
	} else {
		from, ok := mapName2CliAddr[msg.To]
		if !ok {
			fmt.Println("User not exists", msg.To)
			return
		}
		conn, ok := mapCliaddr2Clients[from]
		if !ok {
			fmt.Println("client not exists", from, msg.To)
			return
		}
		conn.Write(data)
	}
}
