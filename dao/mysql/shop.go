package mysql

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/model"
	"database/sql"
	"errors"
	"go.uber.org/zap"
	"strconv"
)

func GetShopListFromDB() (shopList []model.ShopList, err error) {
	rows, err := db.Query(sqlGetShopList)
	if err != nil {
		return nil, co.ErrServerBusy
	}
	for rows.Next() {
		var shop model.ShopList
		if err = rows.Scan(&shop.ShopId, &shop.ShopName, &shop.ShopAddr); err != nil {
			return nil, co.ErrServerBusy
		}
		shopList = append(shopList, shop)
	}
	if err = rows.Err(); err != nil {
		return nil, co.ErrServerBusy
	}
	if len(shopList) == 0 {
		return nil, co.ErrNotFound
	}
	return
}

func GetShopFromDB(id string) (data model.Shop, err error) {
	// 根据商户id从数据库中获取详细信息
	shop_id, err := strconv.Atoi(id)
	if err != nil {
		zap.L().Error("invalid shop_id", zap.String("id", id))
		return data, co.ErrServerBusy
	}
	err = db.QueryRow(sqlGetShop, shop_id).Scan(&data.ShopName, &data.ShopAddr, &data.ShopPhone, &data.ShopDesc, &data.ShopId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return data, co.ErrNotFound
		}
		zap.L().Error("db.QueryRow failed", zap.Error(err))
		return data, co.ErrServerBusy
	}
	return
}

func GetShopMenuFromDB(id int, idInt int) (model.Menu, error) {
	// 根据商户id和菜单id从数据库中获取详细信息
	var shopMenu model.Menu
	err := db.QueryRow(sqlGetShopMenu, id, idInt).Scan(&shopMenu.MenuID, &shopMenu.MenuName, &shopMenu.MenuPrice, &shopMenu.MenuDesc, &shopMenu.MenuStock)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return shopMenu, co.ErrNotFound
		}
		zap.L().Error("db.QueryRow failed", zap.Error(err))
		return shopMenu, co.ErrServerBusy
	}
	return shopMenu, nil

}

func GetShopMenuListFromDB(id int) ([]model.MenuList, error) {
	rows, err := db.Query(sqlGetShopMenuList, id)
	if err != nil {
		return nil, co.ErrServerBusy
	}
	var shopMenuList []model.MenuList
	for rows.Next() {
		var shopMenu model.MenuList
		if err = rows.Scan(&shopMenu.MenuId, &shopMenu.MenuName, &shopMenu.MenuPrice); err != nil {
			return nil, co.ErrServerBusy
		}
		shopMenuList = append(shopMenuList, shopMenu)
	}
	if err = rows.Err(); err != nil {
		return nil, co.ErrServerBusy
	}
	if len(shopMenuList) == 0 {
		return nil, co.ErrNotFound
	}
	return shopMenuList, nil
}

func UpdateShopToDB(m model.Shop) (err error) {
	// 直接更新数据即可
	_, err = db.Exec(sqlUpdateShop, m.ShopName, m.ShopAddr, m.ShopPhone, m.ShopDesc)
	if err != nil {
		zap.L().Error("Update shop failed", zap.String("shopName", m.ShopName), zap.Error(err))
		return co.ErrServerBusy
	}
	return nil
}
