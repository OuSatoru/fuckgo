package main

import (
	"fmt"
	"strings"
)

// strings.Repeat strings.Title
func Accum(s string) string {
	cums := []string{}
	for i, c := range s {
		str := string(c)
		tempstr := strings.ToUpper(str)
		lower := strings.ToLower(str)
		for j := 0; j < i; j++ {
			tempstr += lower
		}
		cums = append(cums, tempstr)
	}
	return strings.Join(cums, "-")
}

func main() {
	fmt.Println(Accum("RqaEzty"))
}
