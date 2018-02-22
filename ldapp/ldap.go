package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/text/encoding/unicode"
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

	pool := x509.NewCertPool()
	pem, err := ioutil.ReadFile("dtrcb-root.pem")
	if err != nil {
		log.Fatal(err)
	}
	pool.AppendCertsFromPEM(pem)

	ls, err := ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", "dtrcb.net", 636), &tls.Config{
		InsecureSkipVerify: true,
		RootCAs:            pool,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer ls.Close()

	err = ls.Bind(admin, adminpwd)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(SearchUser(l, "09800903"))
	// err = DelUser(l, "sbdsb")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = AddUser(l, "09801010", "010602000", "010000000", "Enterprise Staffs")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	err = ModifyPasswordAD(ls, "09800903", "17625094474", "13401766862")
	if err != nil {
		log.Fatal(err)
	}
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

func AddUser(l *ldap.Conn, cn, ou1, ou2, ou3, sn, gn, mobile string) error {
	dn := fmt.Sprintf("cn=%s,ou=%s,ou=%s,ou=%s,dc=dtrcb,dc=net", cn, ou1, ou2, ou3)
	add := ldap.NewAddRequest(dn)
	add.Attribute("cn", []string{cn})
	add.Attribute("objectClass", []string{"user"})
	add.Attribute("sn", []string{sn})
	add.Attribute("givenName", []string{gn})
	add.Attribute("displayName", []string{fmt.Sprintf("%s (%s)", gn, cn)})
	add.Attribute("userPrincipalName", []string{fmt.Sprintf("%s@dtrcb.net", cn)})
	add.Attribute("sAMAccountname", []string{cn})
	add.Attribute("userpassword", []string{mobile})
	err := l.Add(add)
	if err != nil {
		return err
	}
	// https://github.com/go-ldap/ldap/issues/106
	modReq := &ldap.ModifyRequest{
		DN: dn,
		ReplaceAttributes: []ldap.PartialAttribute{
			{Type: "userAccountControl", Vals: []string{"512"}},
		},
	}
	return l.Modify(modReq)
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

func ModifyPasswordAD(l *ldap.Conn, username, oldPwd, newPwd string) error {
	// https://github.com/go-ldap/ldap/issues/106
	utf16 := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	pwdEncoded, err := utf16.NewEncoder().String("\"" + newPwd + "\"")
	if err != nil {
		return err
	}
	toMod := SearchUser(l, username)
	passReq := &ldap.ModifyRequest{
		DN: toMod,
		ReplaceAttributes: []ldap.PartialAttribute{
			{Type: "unicodePwd", Vals: []string{pwdEncoded}},
		},
	}
	return l.Modify(passReq)
}

func ModifyPasswordLDAP(l *ldap.Conn, username, oldPwd, newPwd string) error {
	defer l.Bind(admin, adminpwd)
	toMod := SearchUser(l, username)
	err := l.Bind(toMod, oldPwd)
	if err != nil {
		return err
	}
	passwordModifyRequest := ldap.NewPasswordModifyRequest(toMod, oldPwd, newPwd)
	_, err = l.PasswordModify(passwordModifyRequest)
	return err
}

func ModifyDN(l *ldap.Conn, username, ou1, ou2, ou3 string) error {
	toMod := SearchUser(l, username)
	modRequest := ldap.NewModifyDNRequest(toMod, fmt.Sprintf("cn=%s", username), true, fmt.Sprintf(
		"ou=%s,ou=%s,ou=%s,dc=dtrcb,dc=net", ou1, ou2, ou3))
	return l.ModifyDN(modRequest)
}
