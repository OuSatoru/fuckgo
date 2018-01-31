package main

import (
	"fmt"
	"plugin"
	"time"
)

func main() {
	for {
		p, err := plugin.Open("./toopen.so")
		if err != nil {
			panic(err)
		}
		add, err := p.Lookup("Add")
		if err != nil {
			panic(err)
		}
		sum := add.(func(int, int) int)(11, 22)
		fmt.Println(sum)

		time.Sleep(5 * time.Second)
	}

}
