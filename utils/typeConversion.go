package utils

import (
	"fmt"
	"strconv"
)

//float转float64并保留4位精度
func Float2Float(num float64, format string) float64 {
	floatNum, _ := strconv.ParseFloat(fmt.Sprintf(format, num), 64)
	return floatNum
}

// string 转 int
func String2Int(str string) int {
	intNum, _ := strconv.Atoi(str)
	return intNum
}

// string 转 float64
func String2Float64(str string) float64 {
	floatNum, _ := strconv.ParseFloat(str, 64)
	return floatNum
}

// 字符串转bool
func String2Bool(str string) bool {
	boolVal, _ := strconv.ParseBool(str)
	return boolVal
}
