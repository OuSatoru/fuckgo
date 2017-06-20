package main

import "C"
import "fmt"

//export GoCall
func GoCall(buf *C.char) {
	fmt.Println(C.GoString(buf))
}

//export ForPy
func ForPy(a, b int) int {
	return a + b
}

func main() {}
