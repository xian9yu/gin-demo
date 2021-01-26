package encryption

import (
	"encoding/base64"
)

// base64 编码
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

//  base64 解码
func Base64Decode(str string) (string, []byte) {
	resBytes, _ := base64.StdEncoding.DecodeString(str)
	return string(resBytes), resBytes
}
