package utils

// RemoveLastRune 删除 s 最后 i长度的字符
func RemoveLastRune(s string, i int) string {
	r := []rune(s)
	return string(r[:len(r)-i])
}
