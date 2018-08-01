package common

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"fmt"
	"io"
)

// Sha1Verify sha1 signature verification on rawData+sessionKey with provided signature
func Sha1Verify(rawData, sessionKey, signature string) bool {
	h := sha1.New()
	io.WriteString(h, rawData+sessionKey)
	sigCalculated := fmt.Sprintf("%x", h.Sum(nil))
	return sigCalculated == signature
}

// AESCBCDecrypt AES-128-CBC PKCS#7 base64-decoded key iv secure -> decrypted
func AESCBCDecrypt(key, iv, secure []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCBCDecrypter(block, iv)
	stream.CryptBlocks(secure, secure)
	length := len(secure)
	unpadding := int(secure[length-1])
	secure = secure[:(length - unpadding)]
	return secure, nil
}
