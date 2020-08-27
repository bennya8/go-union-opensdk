package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

type CryptHelper struct {
}

func CryptMD5(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

func CryptMD5Base64(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	return string(base64.StdEncoding.EncodeToString(hash.Sum(nil)))
}
