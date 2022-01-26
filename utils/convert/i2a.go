package convert

import "strconv"

type intConvert interface {
	~uint | ~int | ~int8 | ~int16 | ~int32 | ~int64
}

// MoreInt2String int to string
func MoreInt2String[INT intConvert](num INT) string {
	switch interface{}(num).(type) {
	case uint:
		return strconv.Itoa(int(interface{}(num).(uint)))
	case int:
		return strconv.Itoa(interface{}(num).(int))
	case int8:
		return strconv.Itoa(int(interface{}(num).(int8)))
	case int16:
		return strconv.Itoa(int(interface{}(num).(int16)))
	case int32:
		return strconv.Itoa(int(interface{}(num).(int32)))
	case int64:
		return strconv.Itoa(int(interface{}(num).(int64)))
	}
	return ""
}
