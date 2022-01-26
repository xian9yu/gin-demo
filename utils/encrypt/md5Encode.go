package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

// GetMd5String 生成32位md5字串
func GetMd5String(value, md5Key string) string {
	h := md5.New()
	_, _ = h.Write([]byte(value + md5Key))
	return hex.EncodeToString(h.Sum(nil))
}
