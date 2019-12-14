package main

import (
	"fmt"
	"log"

	"github.com/OuSatoru/fuckgo/ldapp"
)

func main() {
	dn, err := ldapp.SearchUser("09800903")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dn)
}
