package logic

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/dao/redis"
	"RestaurantOrder/model"
)

func GetUserOrderList(id int64) ([]model.OrderInfo, error, co.MyCode) {
	// 直接去redis中查询
	orderList, err := redis.GetUserOrderList(id)
	if err != nil {
		return nil, err, co.CodeServerBusy
	}
	return orderList, nil, co.CodeSuccess
}
