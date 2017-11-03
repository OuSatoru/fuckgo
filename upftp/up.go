package main

import (
	"log"
	"net/textproto"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
)

func main() {
	conn, err := ftp.DialTimeout("32.185.32.78:21", 5*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Quit()
	err = conn.Login("agree", "agree1234")
	if err != nil {
		log.Fatal(err)
	}
	err = conn.NoOp()
	if err != nil {
		log.Fatal(err)
	}
	err = conn.ChangeDir("/tmp")
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	err = conn.Stor("totest.txt", file)
	if err != nil {
		log.Fatal(err)
	}
	err = conn.Logout()
	if err != nil {
		if protoErr := err.(*textproto.Error); protoErr != nil {
			if protoErr.Code != ftp.StatusNotImplemented {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	}
}
