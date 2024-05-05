package redis

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/model"
	"RestaurantOrder/pkg/util"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

func SetOrderInfo(id int64, menu *model.CartInfo) error {
	// 拼接key
	key := strings.Join([]string{co.RedisOrderInfoKey, strconv.FormatInt(id, 10)}, "")
	// 存储到redis
	data, err := util.StructToMap(menu)
	if err != nil {
		zap.L().Error("util.StructToMap(menu) failed", zap.Error(err))
		return co.ErrServerBusy
	}
	_, err = rdb.HMSet(key, data).Result()
	if err != nil {
		zap.L().Error("rdb.HMSet failed", zap.Error(err))
		return co.ErrServerBusy
	}
	return nil
}
