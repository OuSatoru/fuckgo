package ldapp

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/go-ldap/ldap"
	// _ "github.com/mattn/go-oci8"
	"golang.org/x/text/encoding/unicode"
)

type config struct {
	LDAPAdmin    string `json:"ldap_admin"`
	LDAPPassword string `json:"ldap_password"`
	OAAdmin      string `json:"oa_admin"`
	OAPassword   string `json:"oa_password"`
}

var (
	ldapAdmin    string
	ldapPassword string
	oaAdmin      string
	oaPassword   string

	l  *ldap.Conn
	ls *ldap.Conn
)

func init() {
	cb, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalf("read config file err: %v", err)
	}
	var c config
	err = json.Unmarshal(cb, &c)
	if err != nil {
		log.Fatalf("read from config err: %v", err)
	}
	ldapAdmin = c.LDAPAdmin
	ldapPassword = c.LDAPPassword
	oaAdmin = c.OAAdmin
	oaPassword = c.OAPassword
	fmt.Println(ldapAdmin, ldapPassword, oaAdmin, oaPassword)

	l, err = ldap.Dial("tcp", fmt.Sprintf("%s:%d", "dtrcb.net", 389))
	if err != nil {
		log.Fatal(err)
	}

	err = l.Bind(ldapAdmin, ldapPassword)
	if err != nil {
		log.Fatal(err)
	}

	pool := x509.NewCertPool()
	pem, err := ioutil.ReadFile("dtrcb-root.pem")
	if err != nil {
		log.Fatal(err)
	}
	pool.AppendCertsFromPEM(pem)

	ls, err = ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", "dtrcb.net", 636), &tls.Config{
		InsecureSkipVerify: true,
		RootCAs:            pool,
	})
	if err != nil {
		log.Fatal(err)
	}

	err = ls.Bind(ldapAdmin, ldapPassword)
	if err != nil {
		log.Fatal(err)
	}
}

// SearchUser return DN of user
func SearchUser(username string) (string, error) {
	search := ldap.NewSearchRequest(
		"dc=dtrcb, dc=net",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(cn=%s))", username),
		[]string{"dn", "cn"}, nil,
	)
	sr, err := ls.Search(search)
	if err != nil {
		return "", fmt.Errorf("111 %v", err)
	}
	if len(sr.Entries) == 0 {
		return "", errors.New("no user")
	}
	for _, attr := range sr.Entries[0].Attributes {
		fmt.Println(attr.Name, ": ", attr.Values)
	}
	return sr.Entries[0].DN, nil
}

func VerifyUser(username, password string) error {
	defer l.Bind(ldapAdmin, ldapPassword)
	userdn, err := SearchUser(username)
	if err != nil {
		return err
	}
	fmt.Println(userdn)
	err = l.Bind(userdn, password)
	if err != nil {
		return err
	}
	return nil
}

func AddUser(cn, ou1, ou2, ou3, sn, gn, mobile string) error {
	dn := fmt.Sprintf("cn=%s,ou=%s,ou=%s,ou=%s,dc=dtrcb,dc=net", cn, ou1, ou2, ou3)
	add := ldap.NewAddRequest(dn, nil)
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
	modReq := ldap.NewModifyRequest(dn, nil)
	modReq.Replace("userAccountControl", []string{"512"})
	// modReq := &ldap.ModifyRequest{
	// 	DN: dn,
	// 	ReplaceAttributes: []ldap.PartialAttribute{
	// 		{Type: "userAccountControl", Vals: []string{"512"}},
	// 	},
	// }
	return l.Modify(modReq)
}

func DelUser(l *ldap.Conn, cn string) error {
	toDel, err := SearchUser(cn)
	if err != nil {
		return err
	}
	del := ldap.NewDelRequest(toDel, nil)
	return l.Del(del)
}

func ModifyUser(l *ldap.Conn, method, user, attr string, val []string) error {
	modify := ldap.NewModifyRequest(user, nil)
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

func ModifyPasswordAD(username, oldPwd, newPwd string) error {
	// https://github.com/go-ldap/ldap/issues/106
	utf16 := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	pwdEncoded, err := utf16.NewEncoder().String("\"" + newPwd + "\"")
	if err != nil {
		return err
	}
	toMod, err := SearchUser(username)
	if err != nil {
		return err
	}
	passReq := ldap.NewModifyRequest(toMod, nil)
	passReq.Replace("unicodePwd", []string{pwdEncoded})
	// passReq := &ldap.ModifyRequest{
	// 	DN: toMod,
	// 	ReplaceAttributes: []ldap.PartialAttribute{
	// 		{Type: "unicodePwd", Vals: []string{pwdEncoded}},
	// 	},
	// }
	return ls.Modify(passReq)
}

func ModifyPasswordLDAP(username, oldPwd, newPwd string) error {
	defer l.Bind(ldapAdmin, ldapPassword)
	toMod, err := SearchUser(username)
	if err != nil {
		return err
	}
	err = l.Bind(toMod, oldPwd)
	if err != nil {
		return err
	}
	passwordModifyRequest := ldap.NewPasswordModifyRequest(toMod, oldPwd, newPwd)
	_, err = ls.PasswordModify(passwordModifyRequest)
	return err
}

func ModifyDN(l *ldap.Conn, username, ou1, ou2, ou3 string) error {
	toMod, err := SearchUser(username)
	if err != nil {
		return err
	}
	modRequest := ldap.NewModifyDNRequest(toMod, fmt.Sprintf("cn=%s", username), true, fmt.Sprintf(
		"ou=%s,ou=%s,ou=%s,dc=dtrcb,dc=net", ou1, ou2, ou3))
	return l.ModifyDN(modRequest)
}
