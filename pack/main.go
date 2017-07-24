package main

import (
	"encoding/xml"
	"fmt"
)

type Result struct {
	// XMLName xml.Name `xml:"persons"`
	Person []Person `xml:"person"`
}
type Person struct {
	Name      string    `xml:"name,attr"`
	Age       int       `xml:"age,attr"`
	Career    string    `xml:"career"`
	Interests Interests `xml:"interests"`
}
type Interests struct {
	Interest []string `xml:"interest"`
}

func main() {
	raw := `<?xml version="1.0" encoding="utf-8"?>
<persons>
    <person name="polaris" age="28">
        <career>无业游民</career>
        <interests>
            <interest>编程</interest>
            <interest>下棋</interest>
        </interests>
    </person>
</persons>`
	var m Result
	xml.Unmarshal([]byte(raw), &m)
	fmt.Printf("%v\n", m)
}
