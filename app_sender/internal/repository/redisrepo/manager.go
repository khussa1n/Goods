package redisrepo

import (
	"github.com/go-redis/redis/v8"
)

const (
	GoodsList = "goods_list"
)

type Redis struct {
	RedisClient *redis.Client
}

func New(redisClient *redis.Client) *Redis {
	return &Redis{
		RedisClient: redisClient,
	}
}
