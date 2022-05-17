package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 5 {
		fmt.Println("./client tag remoteIP remotePort port")
		return
	}
	port, _ := strconv.Atoi(os.Args[4])
	tag := os.Args[1]
	remoteIP := os.Args[2]
	remotePort, _ := strconv.Atoi(os.Args[3])
	localAddr := net.UDPAddr{Port: port}

	conn, err := net.DialUDP("udp", &localAddr, &net.UDPAddr{IP:net.ParseIP(remoteIP), Port: remotePort})
	if err != nil {
		log.Panic("Failed ot DialUDP", err)
	}
	conn.Write([]byte("我是:" + tag))

	buf := make([]byte, 256)
	n, _,err := conn.ReadFromUDP(buf)
	if err != nil {
		log.Panic("Failed to ReadFromUDP", err)
	}
	conn.Close()
	toAddr := parseAddr(string(buf[:n]))
	fmt.Print("获得对象地址：", toAddr)
	p2p(&localAddr, &toAddr)
}

func parseAddr(addr string) net.UDPAddr {
	t := strings.Split(addr, ":")
	port, _ := strconv.Atoi(t[1])
	return net.UDPAddr{
		IP:   net.ParseIP(t[0]),
		Port: port,
	}
}

func p2p(srcAddr *net.UDPAddr, dstAddr *net.UDPAddr) {
	conn, _ := net.DialUDP("udp", srcAddr, dstAddr)
	conn.Write([]byte("打洞消息\n"))

	go func() {
		buf := make([]byte, 256)
		for {
			n, _, _ := conn.ReadFromUDP(buf)
			if n > 0 {
				fmt.Printf("收到消息：%sp2p", buf[:n])
			}
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("p2p>")
		data, _ := reader.ReadString('\n')
		conn.Write([]byte(data))
	}
}
