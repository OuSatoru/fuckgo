package main

import "C"
import "fmt"

//export GoCall
func GoCall(buf *C.char) {
	fmt.Println(C.GoString(buf))
}

func main() {}
