package utils

import (
	"math/rand"
	"strconv"
	"time"
)

// GetRandomNumber 生成 6位随机数字验证码
func GetRandomNumber() string {
	rand.Seed(time.Now().UnixNano() - time.Now().Unix())
	randNums := strconv.Itoa(rand.Intn(10)) + strconv.Itoa(rand.Intn(10)) +
		strconv.Itoa(rand.Intn(10)) + strconv.Itoa(rand.Intn(10)) +
		strconv.Itoa(rand.Intn(10)) + strconv.Itoa(rand.Intn(10))
	return randNums
}
