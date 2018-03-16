package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

// NATP格式：字段长度 + 字段 + 0 0 0 + 值长度 + 值。

func main() {

	// j := new(bytes.Buffer)
	// j.WriteByte(8)
	// j.WriteString("f.h.jkdm")
	// j.Write([]byte{0, 0, 0})
	// j.WriteByte(10)
	// j.WriteString("1234567890")
	// j.WriteByte(4)
	// j.WriteString("sdfd")
	// j.Write([]byte{0, 0, 0})
	// j.WriteByte(0)
	// // j.WriteString("dd")
	// fmt.Printf("%x\n", j)
	// fmt.Printf("%s\n", string(Encode(j.Bytes())))

	j := []byte(`{"f.h.jkdm":12345,"sdfd":""}`)
	fmt.Printf("%s\n", Decode(j))
	fmt.Printf("%x\n", Decode(j))
}

// Encode : encoding NATP to json
func Encode(n []byte) []byte {
	j := bytes.NewBufferString("{")
	i := 0
	for i < len(n) {
		if n[i] == 0 {
			i += 3
			v := n[i+1 : i+1+int(n[i])]
			_, err := strconv.ParseFloat(string(v), 64)
			if err != nil {
				j.WriteString("\"")
				j.Write(v)
				j.WriteString("\"")
			} else {
				j.Write(v)
			}

			j.WriteString(",")
		} else {
			j.WriteString("\"")
			fmt.Println(int(n[i]))
			j.Write(n[i+1 : i+1+int(n[i])])
			j.WriteString("\":")
		}

		i += 1 + int(n[i])
	}
	j.Truncate(j.Len() - 1)
	j.WriteString("}")
	return j.Bytes()
}

// Decode : json to natp
func Decode(j []byte) []byte {
	var msr map[string]interface{}
	n := new(bytes.Buffer)
	err := json.Unmarshal(j, &msr)
	if err != nil {
		panic(err)
	}
	for k, v := range msr {
		n.WriteByte(byte(len(k)))
		n.WriteString(k)
		n.Write([]byte{0, 0, 0})
		str := fmt.Sprintf("%v", v)
		fmt.Println(str)
		n.WriteByte(byte(len(str)))
		n.WriteString(str)
	}
	return n.Bytes()
}
