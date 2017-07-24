package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

const xmlStr = `<?xml version="1.0" encoding="UTF-8"?>
<Persons>
    <Person name="polaris" age="28">
        <Career>无业游民</Career>
        <Interests>
            <Interest>编程</Interest>
            <Interest>下棋</Interest>
        </Interests>
    </Person>
    <Person name="studygolang" age="27">
        <Career>码农</Career>
        <Interests>
            <Interest>编程</Interest>
            <Interest>下棋</Interest>
        </Interests>
    </Person>
</Persons>`

type Result struct {
	Person []Person `xml:"Person"`
}

type Person struct {
	Name      string    `xml:"name,attr"`
	Age       int       `xml:"age,attr"`
	Career    string    `xml:"Career"`
	Interests Interests `xml:"Interests"`
}

type Interests struct {
	Interest []string `xml:"Interest"`
}

func main() {
	var res Result
	err := xml.Unmarshal([]byte(xmlStr), &res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", res)
}
