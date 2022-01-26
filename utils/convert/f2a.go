package convert

import (
	"fmt"
)

type floatConvert interface {
	~float32 | ~float64
}

func MoreFloat2String[T floatConvert](num T) string {
	switch interface{}(num).(type) {
	case float32:
		// 精度小数点后 6位
		f := fmt.Sprintf("%.6f", interface{}(num).(float32))
		return f
	case float64:
		// 精度小数点后 14位
		f := fmt.Sprintf("%.14f", interface{}(num).(float64))
		return f
	}
	return ""
}
