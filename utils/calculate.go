package utils

import "github.com/shopspring/decimal"

//go 精度计算

// Add 精确加法
func Add(f1, f2 float64) float64 {
	res, _ := decimal.NewFromFloat(f1).Add(decimal.NewFromFloat(f2)).Float64()
	return res
}

// Sub 精确减法
func Sub(f1, f2 float64) float64 {
	res, _ := decimal.NewFromFloat(f1).Sub(decimal.NewFromFloat(f2)).Float64()
	return res
}

// Mul 精确乘法
func Mul(f1, f2 float64) float64 {
	res, _ := decimal.NewFromFloat(f1).Mul(decimal.NewFromFloat(f2)).Float64()
	return res
}

// Div 精确除法
func Div(f1, f2 float64) float64 {
	res, _ := decimal.NewFromFloat(f1).Div(decimal.NewFromFloat(f2)).Float64()
	return res
}

// IntPart 返回小数的整数部分
func IntPart(f float64) int64 {
	return decimal.NewFromFloat(f).IntPart()
}
