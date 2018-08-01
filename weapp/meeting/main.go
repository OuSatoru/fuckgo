package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	_ "github.com/asifjalil/cli"
)

const (
	certDir      = "/home/fr/weapp-t/"
	appid        = "wx453dab41a7b0e145"
	appSecret    = "50e1e56159552db5cc4c9731bd84b757"
	loginVaryURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
)

var (
	db *sql.DB
)

type loginVaryReply struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid,omitempty"`
}

type wechatError struct {
	ErrCode int `json:"errcode"`
	ErrMsg  int `json:"errmsg"`
}

type encryptedUserInfoReply struct {
	OpenID    string `json:"openId"`
	NickName  string `json:"nickName"`
	Gender    int    `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarURL string `json:"avatarUrl"`
	UnionID   string `json:"unionId,omitemply"`
	WaterMark struct {
		AppID     string `json:"appid"`
		TimeStamp int64  `json:"timestamp"`
	}
}

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func init() {
	err := errors.New("temp")
	db, err = sql.Open("cli", fmt.Sprintf("DATABASE=%s; HOSTNAME=%s; PORT=%d; PROTOCOL=TCPIP; UID=%s; PWD=%s;",
		"frxms", "172.17.22.1", 60000, "db2inst1", "db2"))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/onLogin", onLogin)
	http.HandleFunc("/checkRegister", checkRegister)
	http.HandleFunc("/register", register)
	http.HandleFunc("/meetingAttendee", meetingAttendee)
	http.HandleFunc("/realName", realName)
	if err := http.ListenAndServeTLS(":9443", certDir+"214552025140331.pem", certDir+"214552025140331.key", nil); err != nil {
		fmt.Println(err)
	}

}

func onLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// fmt.Println(r.Form)
	code := r.Form.Get("code")
	encryptedData, err := url.QueryUnescape(r.Form.Get("encryptedData"))
	if err != nil {
		log.Printf("malformed encryptedData:%v", err)
		return
	}
	signature, err := url.QueryUnescape(r.Form.Get("signature"))
	if err != nil {
		log.Printf("malformed signature:%v", err)
		return
	}
	iv, err := url.QueryUnescape(r.Form.Get("iv"))
	if err != nil {
		log.Printf("malformed iv:%v", err)
		return
	}
	rawData, err := url.QueryUnescape(r.Form.Get("rawData"))
	if err != nil {
		log.Printf("malformed rawData:%v", err)
		return
	}
	resp, err := http.Get(fmt.Sprintf(loginVaryURL, appid, appSecret, code))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	var loginReply loginVaryReply
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(body, &loginReply)
	if err != nil {
		fmt.Println(err)
		return
	}
	openid := loginReply.OpenID
	sessionKey := loginReply.SessionKey
	fmt.Printf("encryptedData:%s\nsignature:%s\niv:%s\nrawData:%s\nopenid:%s\nsessionKey:%s\n",
		encryptedData, signature, iv, rawData, openid, sessionKey)
	h := sha1.New()
	io.WriteString(h, rawData+sessionKey)
	sigCalculated := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(sigCalculated)
	if sigCalculated != signature {
		fmt.Println("sha1 calculated error.")
		return
	}

	// decrypt
	keyDec, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		log.Print(err)
		return
	}
	ivDec, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		log.Print(err)
		return
	}
	secure, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		log.Print(err)
		return
	}
	block, err := aes.NewCipher(keyDec)
	if err != nil {
		log.Printf("sessionkey error:%v", err)
		return
	}
	stream := cipher.NewCBCDecrypter(block, ivDec)
	stream.CryptBlocks(secure, secure)
	length := len(secure)
	unpadding := int(secure[length-1])
	secure = secure[:(length - unpadding)]
	fmt.Printf("%s\n", secure)

	var infoReply encryptedUserInfoReply
	err = json.Unmarshal(secure, &infoReply)
	if err != nil {
		fmt.Println(err)
		return
	}

	// start db exec
	tx, err := db.Begin()
	if err != nil {
		log.Printf("db context begin error:%v", err)
		return
	}
	_, err = tx.Exec("INSERT INTO WEAPP.LOGINS (CODE, OPENID, SESSIONKEY, APPID, TIMESTAMP) VALUES (?, ?, ?, ?, ?)",
		code, openid, sessionKey, infoReply.WaterMark.AppID, infoReply.WaterMark.TimeStamp)
	if err != nil {
		log.Printf("insert login status err: %v", err)
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("commit error:%v", err)
		return
	}

	// return value
	fmt.Fprint(w, string(secure))
}

func checkRegister(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.Form.Get("openid")
	fmt.Println("openid raw:", id)
	var count int
	row := db.QueryRow("SELECT count(*) FROM WEAPP.USERS WHERE OPENID = ?", id)
	err := row.Scan(&count)
	if err != nil {
		log.Print(err)
		return
	}
	if count != 1 {
		fmt.Fprint(w, "0")
	} else {
		fmt.Fprint(w, "1")
	}

}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		fmt.Println(r.Form)
		openID := r.Form.Get("openid")
		name := r.Form.Get("name")
		phone := r.Form.Get("phone")
		_, err := db.Exec("INSERT INTO WEAPP.USERS (OPENID, NAME, PHONE) VALUES (?, ?, ?)", openID, name, phone)
		if err != nil {
			log.Printf("insert user error:%v", err)
			return
		}
	}
}

func meetingAttendee(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT NAME FROM WEAPP.USERS WHERE REMARK = ?", "1")
	if err != nil {
		log.Printf("fetch meeting attendee error:%v", err)
		return
	}
	defer rows.Close()
	var userList []user
	id := 0
	for rows.Next() {
		var u user
		var name string
		err = rows.Scan(&name)
		if err != nil {
			log.Print(err)
			return
		}
		u.ID = id
		u.Name = name
		userList = append(userList, u)
		id++
	}
	ret, err := json.Marshal(userList)
	if err != nil {
		log.Printf("unmarshal user error:%v", err)
		return
	}
	fmt.Fprint(w, string(ret))
}

func realName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	openID := r.Form.Get("openid")
	var name string
	row := db.QueryRow("SELECT NAME FROM WEAPP.USERS WHERE OPENID = ?", openID)
	err := row.Scan(&name)
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Fprintf(w, `{"name": "%s"}`, name)
}
