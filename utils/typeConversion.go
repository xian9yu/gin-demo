package utils

import (
	"fmt"
	"strconv"
)

// Float2Float float转float64并保留精度
func Float2Float(num float64, format string) float64 {
	floatNum, _ := strconv.ParseFloat(fmt.Sprintf(format, num), 64)
	return floatNum
}

// String2Int string 转 int
func String2Int(str string) int {
	intNum, _ := strconv.Atoi(str)
	return intNum
}

// String2Float64 string 转 float64
func String2Float64(str string) float64 {
	floatNum, _ := strconv.ParseFloat(str, 64)
	return floatNum
}

// String2Bool 字符串转bool
func String2Bool(str string) bool {
	boolVal, _ := strconv.ParseBool(str)
	return boolVal
}
