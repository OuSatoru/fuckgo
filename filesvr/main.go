package main

import (
	atreugo "github.com/savsgio/atreugo/v10"
)

func main() {
	config := atreugo.Config{
		Addr: "0.0.0.0:2333",
	}
	server := atreugo.New(&config)
	server.Static("/", "./static")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
