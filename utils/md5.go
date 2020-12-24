package utils

import (
	"crypto/md5"
	"encoding/hex"
)

const salt = "zhe_shi_salt"

func MD5V(str string) string {
	h := md5.New()
	_, _ = h.Write([]byte(str + salt))
	return hex.EncodeToString(h.Sum(nil))
}
