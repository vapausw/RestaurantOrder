package redis

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/model"
	"RestaurantOrder/pkg/util"
	"errors"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
)

func GetShopList(key string) []string {
	return rdb.LRange(key, 0, -1).Val()
}

func SetShopList(key string, shops []string, seconds time.Duration) error {
	rdb.RPush(key, shops)
	return rdb.Expire(key, seconds).Err()
}

func GetShopMenu(sID int, mID int) string {
	key := strings.Join([]string{co.RedisShopMenuKey, strconv.Itoa(sID), "menu:", strconv.Itoa(mID)}, "")
	return rdb.Get(key).Val()
}

func SetShopMenu(id, idInt int, s string, seconds time.Duration) error {
	key := strings.Join([]string{co.RedisShopMenuKey, strconv.Itoa(id), "menu:", strconv.Itoa(idInt)}, "")
	return rdb.Set(key, s, seconds).Err()
}

func GetShopMenuList(id int) []string {
	key := strings.Join([]string{co.RedisShopMenuListKey, strconv.Itoa(id)}, "")
	return rdb.LRange(key, 0, -1).Val()
}

func SetShopMenuList(id int, menus []string, seconds time.Duration) error {
	key := strings.Join([]string{co.RedisShopMenuListKey, strconv.Itoa(id)}, "")
	rdb.RPush(key, menus)
	return rdb.Expire(key, seconds).Err()
}

func GetShop(key string) (model.Shop, error) {
	res := rdb.Get(key).Val()
	if len(res) == 0 {
		return model.Shop{}, co.ErrNotFound
	}
	var da model.Shop
	err := util.Deserialize(res, &da)
	if err != nil {
		zap.L().Error("util.Deserialize failed", zap.Error(err))
		return da, co.ErrServerBusy
	}
	if da.ShopName == co.BadData {
		return model.Shop{}, co.ErrBadRedisData
	}
	return da, nil
}

func SetShop(key string, shop model.Shop, seconds time.Duration) error {
	da, err := util.Serialize(shop)
	if err != nil {
		zap.L().Error("util.Serialize failed", zap.Error(err))
		return co.ErrServerBusy
	}
	err = rdb.Set(key, da, seconds).Err()
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
