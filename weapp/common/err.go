package common

import (
	"fmt"
)

type wxError struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func makeError(err wxError) error {
	return fmt.Errorf("weixin error: errcode %d errmsg %s", err.ErrCode, err.ErrMsg)
}
