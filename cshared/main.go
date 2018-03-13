package main

import "C"
import (
	"fmt"
)

//export Sum
func Sum(a, b int) int {
	fmt.Println(a, b)
	return a + b
}

func main() {

}
