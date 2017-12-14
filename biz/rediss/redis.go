package rediss

import (
	"github.com/mediocregopher/radix.v2/redis"
	"gaecharge/config"
)

var client *redis.Client

func init() {
	redisClient, err := redis.Dial("tcp", config.AppConfig.Redis.Hosts)
	if nil != err {
		panic(err)
	}

	client = redisClient
}

// 扣钱
func Charge(key string, cost int64) (int, error) {
	left, err := client.Cmd("DECRBY", key, cost).Int()
	if nil != err {
		return -1, err
	}

	return left, nil
}
