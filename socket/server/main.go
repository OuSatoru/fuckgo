package main

import (
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	defer func() {
		if err := recover(); err != nil {
			conn.Close()
		}
	}()
	buf := make([]byte, 4096)
	_, err := conn.Read(buf)
	if err != nil {
		log.Print(err)
		return
	}
	conn.Write([]byte("收到"))
}
