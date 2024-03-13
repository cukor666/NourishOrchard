package config

import (
	"github.com/go-redis/redis/v8"
	"strconv"
)

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

func (r RedisConfig) address() string {
	port := strconv.Itoa(r.Port)
	return r.Host + ":" + port
}

func (r RedisConfig) GetClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     r.address(),
		Password: r.Password,
		DB:       r.Db,
	})
}
