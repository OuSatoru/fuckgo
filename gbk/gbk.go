package gbk

import (
	"bytes"
	"io/ioutil"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func GBKtoUTF8(g string) string {
	gb := []byte(g)
	data, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader(gb), simplifiedchinese.GBK.NewDecoder()))
	return string(data)
}

func UTF8toGBK(u string) string {
	ub := []byte(u)
	data, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader(ub), simplifiedchinese.GBK.NewEncoder()))
	return string(data)
}
