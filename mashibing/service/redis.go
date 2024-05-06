package service

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var rdb *redis.Client

const CookieTime = time.Second * 20

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatal(err)
	}
}

func SetCookie(key string, value string) {
	rdb.Set(context.Background(), key, value, CookieTime)
}

func GetCookie(key string) string {
	return rdb.Get(context.Background(), key).Val()
}
