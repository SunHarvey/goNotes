package client

import (
	"chatserver/protocol"
)

type Client interface {
	Dial(address string) error
	Start()
	Close()
	Send(comand interface{}) error
	SetName(name string) error
	SendMess(message string) error
	InComing() chan protocol.MessCmd
}
