package main

import (
	"fmt"

	"github.com/xuyu/goredis"
)

type GetRedisFunc func(key string) string

type Config struct {
	GetRedis GetRedisFunc
}

var c Config

func GetRedisMessage(key string) string {
	r := goredis.Redis{}
	byteValue, _ := r.Get(key)
	return string(byteValue)
}

func GetMessage() string {
	return fmt.Sprintf("Here The Message %s", c.GetRedis("THE_KEY"))
}

func Init() {
	c = Config{
		GetRedis: GetRedisMessage,
	}
}
