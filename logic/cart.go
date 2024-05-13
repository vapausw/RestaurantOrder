package logic

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/dao/cache"
	"RestaurantOrder/dao/redis"
	"RestaurantOrder/model"
	"errors"
	"go.uber.org/zap"
)

func AddCart(id int64, m *model.CartInfo) (error, co.MyCode) {
	// 先完成加单商品的情况
	// 1. 先查询购物车中是否有该商品,直接去redis中查询
	cart, err := redis.HGetCart(id, m.MenuID)
	// 2. 如果有，更新数量
	if errors.Is(err, co.ErrNotFound) {
		// 不存在，直接插入
		if err = redis.HSetCart(id, m); err != nil {
			return err, co.CodeServerBusy
		}
	} else {
		// 存在，更新数量
		cart.Count += m.Count
		if err = redis.HUpdateCart(id, cart.MenuID, cart.Count); err != nil {
			return err, co.CodeServerBusy
		}
	}
	return nil, co.CodeSuccess
}

func GetCartInfo(id int64) ([]*model.CartInfo, error, co.MyCode) {
	// 直接去redis中查询
	cartList, err := redis.HGetAllCart(id)
	if err != nil {
		return nil, err, co.CodeServerBusy
	}
	return cartList, nil, co.CodeSuccess
}

func CartBuy(info []*model.CartInfo) (error, co.MyCode) {
	// 思考结算请求，判断是否有库存，发起微信支付请求，支付成功后，删除购物车中的商品
	// 1. 判断库存
	for _, menu := range info {
		// 读取数据
		data, err := cache.GetShopMenu(int(menu.ShopID), int(menu.MenuID))
		if err != nil {
			if errors.Is(err, co.ErrNotFound) {
				return co.ErrNotFound, co.CodeNotFound
			}
			return co.ErrServerBusy, co.CodeServerBusy
		}
		if data.MenuStock < menu.Count {
			zap.L().Error("stock not enough")
			return co.ErrStockNotEnough, co.CodeStockNotEnough
		}
	}
	// 2. 发起微信支付请求，此处都默认支付成功因为微信支付api不好调
	// 3. 支付成功后，删除购物车中的商品
	if err := redis.DeleteCart(info); err != nil {
		return co.ErrServerBusy, co.CodeServerBusy
	}
	// 4. 支付成功后还要生成订单信息
	if err := cache.CreateOrder(info); err != nil {
		return co.ErrServerBusy, co.CodeServerBusy
	}
	// 5. 返回成功响应
	return nil, co.CodeSuccess
}

func UpdateCart(cart model.CartInfo) (error, co.MyCode) {
	if err := redis.HUpdateCart(cart.UserID, cart.MenuID, cart.Count); err != nil {
		return err, co.CodeServerBusy
	}
	return nil, co.CodeSuccess
}

func DeleteCart(user_id, menu_id int64) (error, co.MyCode) {
	if err := redis.HDelCart(user_id, menu_id); err != nil {
		return err, co.CodeServerBusy
	}
	return nil, co.CodeSuccess
}
