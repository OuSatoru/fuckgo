package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	str := "0000WA"
	sh := hex.EncodeToString([]byte(str))
	fmt.Println(sh)
}
