package redis

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/model"
	"RestaurantOrder/pkg/util"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

func SetOrderInfo(user_id, order_id int64, menu *model.CartInfo) error {
	// 拼接key
	key := strings.Join([]string{co.RedisOrderInfoKey, strconv.FormatInt(user_id, 10), co.RedisBaseChar, strconv.FormatInt(order_id, 10)}, "")
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

func GetUserOrderList(id int64) (orderList []model.OrderInfo, err error) {
	// 通过用户id进行扫描查询
	keyPattern := strings.Join([]string{co.RedisOrderInfoKey, strconv.FormatInt(id, 10), "*"}, "")
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
			var orderInfo model.OrderInfo
			err = util.MapToStruct(data, &orderInfo)
			if err != nil {
				zap.L().Error("util.MapToStruct(data, &orderInfo) failed", zap.Error(err))
				return nil, co.ErrServerBusy
			}
			orderList = append(orderList, orderInfo)
		}
		if cursor == 0 {
			break
		}
	}
	return
}
