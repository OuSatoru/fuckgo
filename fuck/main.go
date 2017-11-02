package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

func main() {
	str := "0000WA"
	sh := hex.EncodeToString([]byte(str))
	fmt.Println(sh)
	buf := new(bytes.Buffer)
	var data = []int8{2, 3}
	for _, v := range data {
		err := binary.Write(buf, binary.LittleEndian, v)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Printf("%v\n", buf.Bytes())
}
