package redis

import (
	co "RestaurantOrder/constant"
	"errors"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"time"
)

func GetShopList(key string) ([]string, error) {
	res, err := rdb.LRange(key, 0, -1).Result()
	if errors.Is(err, redis.Nil) || len(res) == 0 {
		return nil, co.ErrNotFound

	} else if err != nil {
		zap.L().Error("redis.GetShopList failed", zap.Error(err))
		return nil, co.ErrServerBusy
	} else if res[0] == co.BadData {
		return nil, co.ErrBadRedisData
	}
	return res, err
}

func SetShopList(key string, shops []string, seconds time.Duration) error {
	rdb.RPush(key, shops)
	return rdb.Expire(key, seconds).Err()
}

func DelShopList(key string) error {
	err := rdb.Del(key).Err()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil
		}
		zap.L().Error("redis.DelShopList failed", zap.Error(err))
		return co.ErrServerBusy
	}
	return nil
}

func GetShopMenu(key string) (string, error) {
	res, err := rdb.Get(key).Result()
	if errors.Is(err, redis.Nil) {
		return "", co.ErrNotFound
	}
	if err != nil {
		zap.L().Error("redis.GetShopMenu failed", zap.Error(err))
		return "", co.ErrServerBusy
	}
	if res == co.BadData {
		return "", co.ErrBadRedisData
	}
	return res, nil
}

func SetShopMenu(key, s string, seconds time.Duration) error {
	return rdb.Set(key, s, seconds).Err()
}

func GetShopMenuList(key string) ([]string, error) {
	res, err := rdb.LRange(key, 0, -1).Result()
	if errors.Is(err, redis.Nil) || len(res) == 0 {
		return nil, co.ErrNotFound
	}
	if err != nil {
		zap.L().Error("redis.GetShopMenuList failed", zap.Error(err))
		return nil, co.ErrServerBusy
	}
	if res[0] == co.BadData {
		return nil, co.ErrBadRedisData
	}
	return res, nil
}

func SetShopMenuList(key string, menus []string, seconds time.Duration) error {
	rdb.RPush(key, menus)
	return rdb.Expire(key, seconds).Err()
}

func GetShop(key string) (string, error) {
	res, err := rdb.Get(key).Result()
	if errors.Is(err, redis.Nil) {
		return res, co.ErrNotFound
	} else if err != nil {
		zap.L().Error("redis.GetShop failed", zap.Error(err))
		return res, co.ErrServerBusy
	} else if res == co.BadData {
		return res, co.ErrBadRedisData
	}
	return res, nil
}

func SetShop(key, da string, seconds time.Duration) error {
	err := rdb.Set(key, da, seconds).Err()
	if err != nil {
		zap.L().Error("redis.SetShop failed", zap.Error(err))
		return co.ErrServerBusy
	}
	return nil
}

func DelShop(key string) error {
	err := rdb.Del(key).Err()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil
		}
		zap.L().Error("redis.DelShop failed", zap.Error(err))
		return co.ErrServerBusy
	}
	return nil
}
