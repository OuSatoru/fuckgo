package main

import (
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":2333")
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
	// buf := make([]byte, 1024)
	// _, err := conn.Read(buf)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	conn.Write([]byte("傻逼"))
}
