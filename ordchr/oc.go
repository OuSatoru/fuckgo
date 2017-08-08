package main

import (
	"fmt"
)

func main() {
	by := []byte{97, 98, 99}
	fmt.Println(by)
	fmt.Println(string(by[0]))
	fmt.Println(int(by[1]))
}
