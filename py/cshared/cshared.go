package main

import "C"
import "fmt"

//export GoCall
func GoCall(str *C.char) {
	fmt.Println(C.GoString(str))
}

//export Hello
func Hello() {
	fmt.Println("Hello")
}

//export ForPy
func ForPy(a, b int) int {
	return a + b
}

func main() {}
