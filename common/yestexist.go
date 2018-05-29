package common

import (
	"os"
	"time"
)

func Exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		// os.IsNotExist(err)
		return false
	}
	return true
}

func Yesterday() string {
	now := time.Now()
	return now.AddDate(0, 0, -1).Format("20060102")
}
