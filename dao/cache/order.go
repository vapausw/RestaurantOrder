package cache

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/dao/mysql"
	"RestaurantOrder/dao/redis"
	"RestaurantOrder/model"
	"RestaurantOrder/pkg/snowflake"
)

func CreateOrder(info []*model.CartInfo) error {
	// 生成订单信息
	// 1. 生成订单号,将order表的信息存储到redis中
	// 2. 生成订单信息
	// 3. 生成订单详情信息
	for _, menu := range info {
		order := new(model.Order)
		order.OrderID = snowflake.GenID()
		order.UserID = menu.UserID
		order.ShopID = menu.ShopID
		// 将订单信息存储到mysql中
		if err := mysql.CreateOrder(order); err != nil {
			return co.ErrServerBusy
		}
		// 将订单详细信息存储到redis中
		if err := redis.SetOrderInfo(order.OrderID, menu); err != nil {
			return co.ErrServerBusy
		}
	}
	return nil
}
