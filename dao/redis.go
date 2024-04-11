package dao

import (
	"RestaurantOrder/setting"
	"github.com/go-redis/redis"
)

var (
	Rdb *redis.Client
)

func InitRedis(config *setting.RedisConfig) {
	// TODO
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
		PoolSize: config.PoolSize,
	})
}

func RdbClose() error {
	return Rdb.Close()
}
