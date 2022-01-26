package convert

import "time"

// DateTime2Timestamp 日期时间字符串转时间戳（秒）
func DateTime2Timestamp(datetime string) int64 {
	local, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation("2006-01-02 15:04:05", datetime, local)
	return tmp.Unix() //转化为时间戳 类型是int64

}

// Date2Timestamp 纯日期字符串转时间戳（秒）
func Date2Timestamp(datetime string) int64 {
	local, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation("2006-01-02", datetime, local)
	return tmp.Unix() //转化为时间戳 类型是int64

}

// Timestamp2Date 时间戳(秒)转时间字符串
func Timestamp2Date(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}

// FormatSecond 秒转换为时分秒
func FormatSecond(seconds int64) (day, hour, minute, second int64) {
	day = seconds / (24 * 3600)
	hour = (seconds - day*3600*24) / 3600
	minute = (seconds - day*24*3600 - hour*3600) / 60
	second = seconds - day*24*3600 - hour*3600 - minute*60
	return day, hour, minute, second
}
