package main

import (
    "fmt"
    "net"
)

type Server struct {
    Ip string
    Port int
}

func NewServer(ip string, port int) *Server {
    server := &Server{
        Ip: ip,
        Port: port,
    }
    return server
}


func (s *Server) Handler(conn net.Conn) {
    fmt.Println("connect build success")
    
}


func (s *Server) Start() {
    // socket listen
    listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
    if err !=nil {
        fmt.Println("net.Listen err:", err)
        return
    }
    // close listen socket
    defer listener.Close()

    for {
        // accept
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("listener accept error:", err)
            continue
        }
    
        //do handler
                
        go s.Handler(conn)

        
    }
}



