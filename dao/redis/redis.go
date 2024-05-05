package redis

import (
	"RestaurantOrder/setting"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var (
	rdb *redis.Client
)

func Init(config *setting.RedisConfig) error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.Db,
		PoolSize: config.PoolSize,
	})
	_, err := rdb.Ping().Result()
	return err
}

func Close() {
	if err := rdb.Close(); err != nil {
		zap.L().Error("close redis failed", zap.Error(err))
		panic(err)
	}
}
