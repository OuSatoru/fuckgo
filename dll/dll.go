package main

import (
	"syscall"
	"fmt"
)

func main() {
	dll := syscall.NewLazyDLL("embed.dll")
	fmt.Println("Called dll: ", dll.Name)
	p := dll.NewProc("process")
	p.Call()
	fmt.Println("end")
}
