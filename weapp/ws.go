package main

import (
	"log"

	"golang.org/x/net/websocket"
)

func wxHandler(ws *websocket.Conn) {
	for {
		var recv string
		if err := websocket.Message.Receive(ws, &recv); err != nil {
			log.Fatal(err)
		}
		repl := "from server: " + recv
		if err := websocket.Message.Send(ws, repl); err != nil {
			log.Fatal(err)
		}
	}
}
