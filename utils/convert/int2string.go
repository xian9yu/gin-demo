package convert

import "strconv"

type intConvert interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~int | ~int8 | ~int16 | ~int32 | ~int64
}

// IntToString int to string
func IntToString[INT intConvert](num INT) string {
	switch interface{}(num).(type) {
	case uint:
		return strconv.Itoa(int(interface{}(num).(uint)))
	case uint8:
		return strconv.Itoa(int(interface{}(num).(uint8)))
	case uint16:
		return strconv.Itoa(int(interface{}(num).(uint16)))
	case uint32:
		return strconv.Itoa(int(interface{}(num).(uint32)))
	case uint64:
		return strconv.Itoa(int(interface{}(num).(uint64)))
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
