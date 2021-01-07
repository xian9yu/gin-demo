package models

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client

func initClient() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//检测心跳
	_, err := Rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}
