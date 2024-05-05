package logic

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/dao/cache"
	"RestaurantOrder/model"
	"errors"
)

func GetShop(shopId int) (data model.Shop, err error, code co.MyCode) {
	// 从数据库中获取商店信息，任然从缓存中读取，没有再去数据库中读取
	data, err = cache.GetShop(shopId)
	if err != nil {
		if errors.Is(err, co.ErrNotFound) {
			return data, co.ErrNotFound, co.CodeNotFound
		}
		return data, co.ErrServerBusy, co.CodeServerBusy
	}
	return data, nil, co.CodeSuccess
}

func GetShopList() ([]model.ShopList, error, co.MyCode) {
	// 从数据库中获取商店信息，任然从缓存中读取，没有再去数据库中读取
	data, err := cache.GetShopList()
	if err != nil {
		if errors.Is(err, co.ErrNotFound) {
			return nil, co.ErrNotFound, co.CodeNotFound
		}
		return nil, co.ErrServerBusy, co.CodeServerBusy
	}
	return data, nil, co.CodeSuccess
}

func GetShopMenuList(id int) ([]model.MenuList, error, co.MyCode) {
	// 从数据库中获取商店信息，任然从缓存中读取，没有再去数据库中读取
	data, err := cache.GetShopMenuList(id)
	if err != nil {
		if errors.Is(err, co.ErrNotFound) {
			return nil, co.ErrNotFound, co.CodeNotFound
		}
		return nil, co.ErrServerBusy, co.CodeServerBusy
	}
	return data, nil, co.CodeSuccess
}

func GetShopMenu(shop_id int, menu_id int) (model.Menu, error, co.MyCode) {
	data, err := cache.GetShopMenu(shop_id, menu_id)
	if err != nil {
		if errors.Is(err, co.ErrNotFound) {
			return data, co.ErrNotFound, co.CodeNotFound
		}
		return data, co.ErrServerBusy, co.CodeServerBusy
	}
	return data, nil, co.CodeSuccess
}
