package utils

import (
	"math/rand"
	"strconv"
	"time"
)

const ExpireTime int64 = 86400 // token有效期 (单位: s)

func NewToken(userName, userId string) string {
	return userName + "_" + userId + "_" + strconv.Itoa(rand.Intn(10)) + strconv.Itoa(rand.Intn(10)) + strconv.FormatInt(time.Now().Unix(), 10)
}
