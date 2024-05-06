package service

import (
	"fmt"
	"github.com/go-redis/redis"
)

func init() {
	c, err := redis.Dial("tcp", "localhost:6379")
	defer c.Close()
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	fmt.Println("redis conn success")

}
