package mysql

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/model"
	"RestaurantOrder/pkg/util"
	"database/sql"
	"errors"
	"go.uber.org/zap"
	"strconv"
)

func GetShopListFromDB() (data []string, err error) {
	shopList := make([]model.ShopList, 0, 100)
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
	data, err = util.SerializeShops(shopList)
	if err != nil {
		return nil, co.ErrServerBusy
	}
	return
}

func GetShopFromDB(key string) (da string, err error) {
	var shopID int
	for i := len(key) - 1; i >= 0; i-- {
		if key[i] == ':' {
			shopID, _ = strconv.Atoi(key[i+1:])
			break
		}
	}
	var data model.Shop
	err = db.QueryRow(sqlGetShop, shopID).Scan(&data.ShopName, &data.ShopAddr, &data.ShopPhone, &data.ShopDesc, &data.ShopId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", co.ErrNotFound
		}
		zap.L().Error("db.QueryRow failed", zap.Error(err))
		return "", co.ErrServerBusy
	}
	da, err = util.Serialize(data)
	if err != nil {
		zap.L().Error("util.Serialize failed", zap.Error(err))
		return "", co.ErrServerBusy
	}
	return
}

func GetShopMenuFromDB(key string) (string, error) {
	var shopId, menuId int
	for i := len(key) - 1; i >= 0; i-- {
		if key[i] == ':' {
			menuId, _ = strconv.Atoi(key[i+1:])
			for j := i - 6; j >= 0; j-- {
				if key[j] == ':' {
					shopId, _ = strconv.Atoi(key[j+1 : i-5])
					break
				}
			}
			break
		}
	}
	// 根据商户id和菜单id从数据库中获取详细信息
	var shopMenu model.Menu
	err := db.QueryRow(sqlGetShopMenu, shopId, menuId).Scan(&shopMenu.MenuID, &shopMenu.MenuName, &shopMenu.MenuPrice, &shopMenu.MenuDesc, &shopMenu.MenuStock)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", co.ErrNotFound
		}
		zap.L().Error("db.QueryRow failed", zap.Error(err))
		return "", co.ErrServerBusy
	}
	data, err := util.Serialize(shopMenu)
	if err != nil {
		zap.L().Error("util.Serialize failed", zap.Error(err))
		return "", co.ErrServerBusy
	}
	return data, nil

}

func GetShopMenuListFromDB(key string) ([]string, error) {
	// 从key中获取id
	for i := len(key) - 1; i >= 0; i-- {
		if key[i] == ':' {
			key = key[i+1:]
			break
		}
	}
	id, err := strconv.Atoi(key)
	if err != nil {
		zap.L().Error("invalid shop_id", zap.String("key", key))
		return nil, co.ErrServerBusy
	}
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
	data, err := util.SerializeMenus(shopMenuList)
	if err != nil {
		zap.L().Error("util.SerializeMenus failed", zap.Error(err))
		return nil, co.ErrServerBusy
	}
	return data, nil
}

func UpdateShopToDB(da string) (err error) {
	var m model.Shop
	err = util.Deserialize(da, &m)
	if err != nil {
		zap.L().Error("util.Deserialize failed", zap.Error(err))
		return co.ErrServerBusy
	}
	// 直接更新数据即可
	_, err = db.Exec(sqlUpdateShop, m.ShopName, m.ShopAddr, m.ShopPhone, m.ShopDesc)
	if err != nil {
		zap.L().Error("Update shop failed", zap.String("shopName", m.ShopName), zap.Error(err))
		return co.ErrServerBusy
	}
	return nil
}
