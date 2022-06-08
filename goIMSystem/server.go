package main

import (
    "fmt"
    "net"
    "sync"
)

type Server struct {
    Ip string
    Port int
    
    //online user list
    OnlineMap map[string]*User
    mapLock sync.RWMutex

    // message boardcast channel
    Message chan string 
}

func NewServer(ip string, port int) *Server {
    server := &Server{
        Ip: ip,
        Port: port,
        OnlineMap: make(map[string]*User),
        Message: make(chan string),
    }
    return server
}

func (s *Server) ListenMessager() {
    for {
        msg := <- s.Message
        s.mapLock.Lock()
        for _, cli := range s.OnlineMap {
            cli.C <- msg
        }
        s.mapLock.Unlock()
    }
}

func (s *Server) BroadCast(user *User,msg string) {
    sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
    
    s.Message <- sendMsg 
}


func (s *Server) Handler(conn net.Conn) {
    //fmt.Println("connect build success")
    user := NewUser(conn)
    
    // user online ,add user to onlineMap
    s.mapLock.Lock()
    s.OnlineMap[user.Name] = user
    s.mapLock.Unlock()
    
    // broadcast user online message
    s.BroadCast(user, "I am online") 
    
    //
    select {}
    
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


    // start  listener message goroutine
    go s.ListenMessager()

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



