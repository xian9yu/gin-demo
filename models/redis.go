package models

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var (
	Rdb *redis.Client
	ctx = context.Background()
	cfg = c.InitConfig() //初始化配置文件
)

func InitClient() *redis.Client {
	hp := cfg.GetString("Redis.HP")
	pw := cfg.GetString("Redis.Password")

	Rdb = redis.NewClient(&redis.Options{
		Addr:     hp,  // host:port
		Password: pw,  // set password
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})

	//检测心跳
	_, err := Rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalln("redis 连接失败 error:", err)
	}
	return Rdb
}

//执行任意/自定义命令
func StrDo(function, key string) (interface{}, error) {
	//val, err := Rdb.Do(ctx, "get", "key").Result()
	val, err := Rdb.Do(ctx, function, key).Result()
	if err != nil {
		return nil, err
	}
	return val, nil
}

//Set 方法的最后一个参数表示过期时间，0 表示永不过期
func StrSet(key, value string) error {
	err := Rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

/* NOTE:
SetNX()与SetEX()的区别是，SexNX()仅当key不存在的时候才设置，
如果key已经存在则不做任何操作，而SetEX()方法不管该key是否已经存在缓存中直接覆盖
*/
//设置键的同时，设置过期时间（ SetEX()方法不管该 key是否已经存在缓存中直接覆盖过期时间）
func StrSetEX(key, value string, expiration time.Duration) error {
	err := Rdb.SetEX(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

// SetNX()仅当 key不存在的时候才设置过期时间
func StrSetNX(key, value string, expiration time.Duration) (bool, error) {
	res, err := Rdb.SetNX(ctx, key, value, expiration).Result()
	if err != nil {
		return res, err
	}
	return res, nil
}

//如果要获取的key在缓存中并不存在，Get()方法将会返回redis.Nil
func StrGet(key string) (string, error) {
	val, err := Rdb.Get(ctx, key).Result()
	if err != nil {
		return val, err
	}
	return val, nil
}

// 批量查询key的值
func StrMGet(key string) ([]interface{}, error) {
	val, err := Rdb.MGet(ctx, key).Result()
	if err != nil {
		return val, err
	}
	return val, nil
}

func StrExists(key string) (res bool, err error) {
	val, err := Rdb.Exists(ctx, key).Result()
	if err != nil {
		res = val == 0
		return res, err
	}
	res = val == 1
	return res, nil
}

func StrDel(key string) (int64, error) {
	res, err := Rdb.Del(ctx, key).Result()
	if err != nil {
		return res, err
	}
	return res, nil
}
func StrGetRange(key string) (string, error) {
	res, err := Rdb.GetRange(ctx, key, 0, -1).Result()
	if err != nil {
		return res, err
	}
	return res, nil
}

func StrSetExpireAt(key string, expireTime int64) error {
	err := Rdb.ExpireAt(ctx, key, time.Unix(expireTime, 0)).Err()
	if err != nil {
		return err
	}
	return nil
}
func Ttl(key string) (val int, err error) {
	duration, err := Rdb.TTL(ctx, key).Result()
	if err != nil {
		return
	}
	val = int(duration / time.Second)
	return
}
