package dao

import (
	"github.com/go-redis/redis/v8"
	"server/config"
)

var redisDB = config.GetConfig().RedisConfig.GetClient()

// GetRedis 开放给外部访问redis实例
func GetRedis() *redis.Client {
	return redisDB
}
