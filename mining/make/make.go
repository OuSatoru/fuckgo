package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	numStr := "1"
	sha := sha256.New()
	sha.Write([]byte(numStr))
	fmt.Printf("%x\n", sha.Sum(nil))
}
