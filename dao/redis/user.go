package redis

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/model"
	"RestaurantOrder/pkg/util"
	"errors"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"strings"
	"time"
)

func GetUserPassword(email string) (user *model.User, err error) {
	key := strings.Join([]string{co.RedisUserPasswordKey, email}, "")
	data, err := rdb.HGetAll(key).Result()
	if err != nil {
		zap.L().Error("rdb.Get(key).Result() failed", zap.Error(err))
		return &model.User{}, co.ErrServerBusy
	}
	if len(data) == 0 {
		return &model.User{}, co.ErrNotFound
	}
	user = new(model.User)
	err = util.MapToStruct(data, user)
	if err != nil {
		zap.L().Error("util.MapToStruct(data, &cart) failed", zap.Error(err))
		return &model.User{}, co.ErrServerBusy
	}
	if user.Password == co.BadData {
		return &model.User{}, co.ErrBadRedisData
	}
	return
}
func SetUserPassword(email string, user *model.User, expiration time.Duration) error {
	// 拼接key
	key := strings.Join([]string{co.RedisUserPasswordKey, email}, "")
	// 存储到redis
	data, err := util.StructToMap(user)
	if err != nil {
		zap.L().Error("util.StructToMap(m) failed", zap.Error(err))
		return co.ErrServerBusy
	}
	_, err = rdb.HMSet(key, data).Result()
	if err != nil {
		zap.L().Error("rdb.HMSet(key, data).Result() failed", zap.Error(err))
		return co.ErrServerBusy
	}
	err = rdb.Expire(key, expiration).Err()
	return nil
}

func SetUserCode(email, code string) error {
	key := strings.Join([]string{co.RedisUserCodeKey, email}, "")
	return rdb.Set(key, code, time.Minute*5).Err()
}

func GetUserCode(email string) (string, error) {
	key := strings.Join([]string{co.RedisUserCodeKey, email}, "")
	res, err := rdb.Get(key).Result()
	if errors.Is(err, redis.Nil) {
		return "", co.ErrNotFound
	}
	if err != nil {
		zap.L().Error("rdb.Get(key).Result() failed", zap.Error(err))
		return "", co.ErrServerBusy
	}
	return res, err
}

func DelUserPassword(email string) error {
	key := strings.Join([]string{co.RedisUserCodeKey, email}, "")
	err := rdb.Del(key).Err()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil
		}
		zap.L().Error("rdb.Del(key).Err() failed", zap.Error(err))
		return co.ErrServerBusy
	}
	return nil
}
