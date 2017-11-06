package main

import (
	"fmt"
	"log"

	"gopkg.in/ldap.v2"
)

func main() {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", "dtrcb.net", 389))
	if err != nil {
		panic(err)
	}
	defer l.Close()

	err = l.Bind("-", "-")
	if err != nil {
		log.Fatal(err)
	}

	search := ldap.NewSearchRequest(
		"dc=dtrcb, dc=net",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(&(objectClass=organizationalPerson))", []string{"dn", "cn"}, nil,
	)
	sr, err := l.Search(search)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range sr.Entries {
		fmt.Printf("%s: %v\n", entry.DN, entry.GetAttributeValue("cn"))
	}
}
