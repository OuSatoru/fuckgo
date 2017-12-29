package main

import (
	"fmt"
	"log"

	"gopkg.in/ldap.v2"
)

const (
	admin    = "-"
	adminpwd = "-"
)

func main() {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", "dtrcb.net", 389))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	err = l.Bind(admin, adminpwd)
	if err != nil {
		log.Fatal(err)
	}

	// ls, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", "dtrcb.net", 636))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer ls.Close()

	// err = ls.StartTLS(&tls.Config{InsecureSkipVerify: true})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = ls.Bind(admin, adminpwd)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println(VerifyUser(l, "09800903", "13401766862"))
}

func SearchUser(l *ldap.Conn, username string) string {
	search := ldap.NewSearchRequest(
		"dc=dtrcb, dc=net",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(cn=%s))", username),
		[]string{"dn", "cn"}, nil,
	)
	sr, err := l.Search(search)
	if err != nil {
		log.Fatal(err)
	}
	if len(sr.Entries) == 0 {
		return ""
	}
	for _, attr := range sr.Entries[0].Attributes {
		fmt.Println(attr.Name, ": ", attr.Values)
	}
	return sr.Entries[0].DN
}

func VerifyUser(l *ldap.Conn, username, password string) bool {
	defer l.Bind(admin, adminpwd)
	search := ldap.NewSearchRequest(
		"dc=dtrcb, dc=net",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(cn=%s))", username),
		[]string{"dn", "cn"},
		nil,
	)
	sr, err := l.Search(search)
	if err != nil {
		log.Fatal(err)
	}
	if len(sr.Entries) != 1 {
		log.Fatal("User does not exist or too many entries returned")
	}
	userdn := sr.Entries[0].DN
	fmt.Println(userdn)
	err = l.Bind(userdn, password)
	if err == nil {
		return true
	}
	return false
}

func AddUser(l *ldap.Conn, cn, ou1, ou2, ou3 string) error {
	add := ldap.NewAddRequest(fmt.Sprintf(
		"cn=%s,ou=%s,ou=%s,ou=%s,dc=dtrcb,dc=net", cn, ou1, ou2, ou3))
	add.Attribute("cn", []string{"0601002"})
	add.Attribute("objectClass", []string{"user"})
	add.Attribute("sn", []string{"金"})
	add.Attribute("givenName", []string{"金刚"})
	add.Attribute("displayName", []string{"金刚 (0601002)"})
	add.Attribute("userPrincipalName", []string{"0601002@dtrcb.net"})
	add.Attribute("sAMAccountname", []string{"0601002"})
	add.Attribute("userpassword", []string{"13401766862"})
	return l.Add(add)
}

func DelUser(l *ldap.Conn, cn string) error {
	toDel := SearchUser(l, cn)
	del := ldap.NewDelRequest(toDel, nil)
	return l.Del(del)
}

func ModifyUser(l *ldap.Conn, method, user, attr string, val []string) error {
	modify := ldap.NewModifyRequest(user)
	switch method {
	case "add":
		modify.Add(attr, val)
	case "delete":
		modify.Delete(attr, val)
	case "replace":
		modify.Replace(attr, val)
	}
	err := l.Modify(modify)
	if err != nil {
		return err
	}
	return nil
}

func ModifyPassword(l *ldap.Conn, username, newPwd string) error {
	toMod := SearchUser(l, username)
	passwordModifyRequest := ldap.NewPasswordModifyRequest(toMod, "", newPwd)
	_, err := l.PasswordModify(passwordModifyRequest)
	return err
}

func ModifyDN(l *ldap.Conn, username, ou1, ou2, ou3 string) error {
	toMod := SearchUser(l, username)
	modRequest := ldap.NewModifyDNRequest(toMod, fmt.Sprintf("cn=%s", username), true, fmt.Sprintf(
		"ou=%s,ou=%s,ou=%s,dc=dtrcb,dc=net", ou1, ou2, ou3))
	return l.ModifyDN(modRequest)
}
