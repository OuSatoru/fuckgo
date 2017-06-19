package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	var buf [1024]byte
	service := os.Args[1]
	addr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	rAddr := conn.RemoteAddr()
	for {
		var str string
		fmt.Scanln(&str)
		if str == "exit" {
			return
		}
		n, err := conn.Write([]byte(str))
		if err != nil {
			log.Fatal(err)
		}
		n, err = conn.Read(buf[0:])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Reply from server ", rAddr.String(), string(buf[0:n]))
	}

}
