package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	loginVerifyURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
)

// LoginVerifyReply returned by get loginVerifyURL
type LoginVerifyReply struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid,omitempty"`
}

// DecryptedUserInfoReply unnecessary if only want openid
type DecryptedUserInfoReply struct {
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

// LoginCredentialVerify get openid, sessionkey, unionid by code
func LoginCredentialVerify(appid, appSecret, code string) (LoginVerifyReply, error) {
	var loginReply LoginVerifyReply
	resp, err := http.Get(fmt.Sprintf(loginVerifyURL, appid, appSecret, code))
	if err != nil {
		return loginReply, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return loginReply, err
	}
	if strings.Contains(string(body), "errcode") {
		var e wxError
		err = json.Unmarshal(body, &e)
		if err != nil {
			return loginReply, err
		}
		return loginReply, makeError(e)
	}
	err = json.Unmarshal(body, &loginReply)
	if err != nil {
		return loginReply, err
	}
	return loginReply, nil
}

func UserInfo(key, iv, secure []byte) (DecryptedUserInfoReply, error) {
	var user DecryptedUserInfoReply
	b, err := AESCBCDecrypt(key, iv, secure)
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(b, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}
