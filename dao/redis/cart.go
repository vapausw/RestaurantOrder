package redis

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/model"
	"RestaurantOrder/pkg/util"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

func HGetCart(id int64, menuID int64) (cart *model.CartInfo, err error) {
	// 从redis中获取购物车的数据
	// 1. 拼接key
	key := strings.Join([]string{co.RedisCartKey, strconv.FormatInt(id, 10), co.RedisBaseChar, strconv.FormatInt(menuID, 10)}, "")
	// 2. 查询
	data, err := rdb.HGetAll(key).Result()
	if err != nil {
		zap.L().Error("rdb.Get(key).Result() failed", zap.Error(err))
		return nil, co.ErrServerBusy
	}
	if len(data) == 0 {
		return nil, co.ErrNotFound
	}
	err = util.MapToStruct(data, cart)
	if err != nil {
		zap.L().Error("util.MapToStruct(data, &cart) failed", zap.Error(err))
		return nil, co.ErrServerBusy
	}
	return
}

func HSetCart(id int64, m *model.CartInfo) (err error) {
	// 拼接key
	key := strings.Join([]string{co.RedisCartKey, strconv.FormatInt(id, 10), co.RedisBaseChar, strconv.FormatInt(m.MenuID, 10)}, "")
	// 存储到redis
	data, err := util.StructToMap(m)
	if err != nil {
		zap.L().Error("util.StructToMap(m) failed", zap.Error(err))
		return co.ErrServerBusy
	}
	_, err = rdb.HMSet(key, data).Result()
	// 异步将数据更新到mysql中保存
	//go dao.AddCart(m)
	return
}

func HUpdateCart(id, menuID int64, count int) (err error) {
	key := strings.Join([]string{co.RedisCartKey, strconv.FormatInt(id, 10), co.RedisBaseChar, strconv.FormatInt(menuID, 10)}, "")
	return rdb.HSet(key, "count", count).Err()
}

func HDelCart(user_id, menu_id int64) error {
	key := strings.Join([]string{co.RedisCartKey, strconv.FormatInt(user_id, 10), co.RedisBaseChar, strconv.FormatInt(menu_id, 10)}, "")
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
func HGetAllCart(id int64) (cartList []*model.CartInfo, err error) {
	// 模糊匹配
	keyPattern := fmt.Sprintf("%s%d%s*", co.RedisCartKey, id, co.RedisBaseChar)
	zap.L().Info("keyPattern", zap.String("keyPattern", keyPattern))
	var cursor uint64
	for {
		keys, cur, err := rdb.Scan(cursor, keyPattern, 10).Result()
		zap.L().Info("rdb.Scan(cursor, keyPattern, 10).Result()", zap.Any("keys", keys), zap.Uint64("cursor", cursor))
		if err != nil {
			zap.L().Error("rdb.Scan(cursor, keyPattern, 10).Result() failed", zap.Error(err))
			return nil, co.ErrServerBusy
		}
		cursor = cur
		for _, key := range keys {
			data, err := rdb.HGetAll(key).Result()
			if err != nil {
				zap.L().Error("rdb.HGetAll(key).Result() failed", zap.Error(err))
				return nil, co.ErrServerBusy
			}
			// 解析值并构建 CartInfo
			cartInfo := new(model.CartInfo)
			err = util.MapToStruct(data, cartInfo)
			if err != nil {
				zap.L().Error("util.MapToStruct(data, &cart) failed", zap.Error(err))
				return nil, co.ErrServerBusy
			}
			cartList = append(cartList, cartInfo)
		}
		if cursor == 0 {
			break
		}
	}
	return
}

func DeleteCart(info []*model.CartInfo) error {
	for _, menu := range info {
		key := strings.Join([]string{co.RedisCartKey, strconv.FormatInt(menu.UserID, 10), co.RedisBaseChar, strconv.FormatInt(menu.MenuID, 10)}, "")
		if err := rdb.Del(key).Err(); err != nil {
			zap.L().Error("rdb.Del(key).Err() failed", zap.Error(err))
			return co.ErrServerBusy
		}
	}
	return nil
}
