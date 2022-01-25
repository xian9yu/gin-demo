package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

const md5Key = "thisIsMd5SecureKey"

// GetMd5String 生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	_, _ = h.Write([]byte(s + md5Key))
	return hex.EncodeToString(h.Sum(nil))
}
