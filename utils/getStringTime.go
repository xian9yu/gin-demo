package utils

import "time"

// 获取字符串日期
func StrDate() string {
	return time.Now().Format("2006-01-02")
}

func StrDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
