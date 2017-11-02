package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s host:port\n", os.Args[0])
	}
	tcpAddr, err := net.ResolveTCPAddr("tcp", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	// _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	result, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(result))
}
