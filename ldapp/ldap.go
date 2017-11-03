package main

import (
	"fmt"

	ldap "gopkg.in/ldap.v2"
)

func main() {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", "dtrcb.net", 389))
	if err != nil {
		panic(err)
	}
	defer l.Close()
}
