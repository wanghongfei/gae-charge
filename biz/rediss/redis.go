package rediss

import (
	"github.com/mediocregopher/radix.v2/redis"
	"gaecharge/config"
	"log"
)

var client *redis.Client

func InitRedis() {
	redisClient, err := redis.Dial("tcp", config.AppConfig.Redis.Hosts)
	if nil != err {
		panic(err)
	}

	client = redisClient

	log.Printf("redis initialized")
}

// 扣钱
func Charge(key string, cost int64) (int, error) {
	left, err := client.Cmd("DECRBY", key, cost).Int()
	if nil != err {
		return -1, err
	}

	return left, nil
}
