package cache

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/dao/mysql"
	"RestaurantOrder/dao/redis"
	"RestaurantOrder/model"
	"RestaurantOrder/pkg/util"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"strconv"
)

type CacheShop struct {
	model.Shop
}

func (c *CacheShop) RedisGet(key string) (interface{}, error) {
	keys := fmt.Sprintf("%s%s", co.RedisShopKey, key)
	return redis.GetShop(keys)
}

func (c *CacheShop) RedisSet(key string, value interface{}) error {
	keys := fmt.Sprintf("%s%s", co.RedisShopKey, key)
	return redis.SetShop(keys, value.(model.Shop), util.GetRandomExpirationInSeconds(1800, 3600))
}

func (c *CacheShop) RedisDel(key string) error {
	keys := fmt.Sprintf("%s%s", co.RedisShopKey, key)
	return redis.DelShop(keys)
}

func (c *CacheShop) MysqlGet(key string) (interface{}, error) {
	return mysql.GetShopFromDB(key)
}

func (c *CacheShop) MysqlSet(key string, value interface{}) error {
	return mysql.UpdateShopToDB(value.(model.Shop))
}

func (c *CacheShop) BadData() interface{} {
	return model.Shop{
		ShopName: co.BadData,
	}
}

func GetShop(id int) (model.Shop, error) {
	cache_shop := new(CacheShop)
	data, err := GetDataFromCache(cache_shop, strconv.Itoa(id))
	if err != nil {
		return model.Shop{}, err
	}
	return data.(model.Shop), nil
}

//func GetShop(id int) (*model.Shop, error) {
//	// 从缓存中获取商店信息
//	shop := redis.GetShop(id)
//	if shop == co.BadData {
//		return nil, co.ErrNotFound
//	}
//	if len(shop) == 0 {
//		// 从数据库中获取商店信息
//		mshop, err := mysql.GetShopFromDB(id)
//		//zap.L().Info("mysql.GetShopFromDB", zap.Any("mshop", mshop))
//		if err != nil {
//			if errors.Is(err, co.ErrNotFound) {
//				// 将缓存中的数据设置为空，防止缓存穿透
//				err = redis.SetShop(id, co.BadData, util.GetRandomExpirationInSeconds(1800, 3600))
//				return &mshop, co.ErrNotFound
//			}
//			return &mshop, co.ErrServerBusy
//		}
//		// 将结构体序列化为字符串
//		da, _ := util.Serialize(mshop)
//		//zap.L().Info("util.Serialize", zap.Any("da", da))
//		// 将商店信息缓存起来
//		err = redis.SetShop(id, da, util.GetRandomExpirationInSeconds(1800, 3600))
//		if err != nil {
//			zap.L().Error("redis.SetShop failed", zap.Error(err))
//			return &mshop, co.ErrServerBusy
//		}
//		return &mshop, nil
//	}
//	var da model.Shop
//	err := util.Deserialize(shop, &da)
//	if err != nil {
//		return &da, err
//	}
//	return &da, nil
//}

func GetShopList() ([]model.ShopList, error) {
	// 从缓存中获取商店信息
	shops := redis.GetShopList(co.RedisShopListKey)
	if len(shops) > 0 && shops[0] == co.BadData {
		return nil, co.ErrNotFound
	}
	if len(shops) == 0 {
		// 从数据库中获取商店信息
		mshops, err := mysql.GetShopListFromDB()
		if err != nil {
			if errors.Is(err, co.ErrNotFound) {
				// 将缓存中的数据设置为空，防止缓存穿透
				err = redis.SetShopList(co.RedisShopListKey, []string{co.BadData}, util.GetRandomExpirationInSeconds(1800, 3600))
				return nil, co.ErrNotFound
			}
			return nil, co.ErrServerBusy
		}
		// 将结构体数组序列化为字符串数组
		da, _ := util.SerializeShops(mshops)
		// 将商店信息缓存起来
		err = redis.SetShopList(co.RedisShopListKey, da, util.GetRandomExpirationInSeconds(1800, 3600))
		if err != nil {
			zap.L().Error("redis.SetShopList failed", zap.Error(err))
			return nil, co.ErrServerBusy
		}
		return mshops, nil
	}
	var data []model.ShopList
	err := util.DeserializeShops(shops, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetShopMenuList(id int) ([]model.MenuList, error) {
	// 从缓存中获取商店信息
	menus := redis.GetShopMenuList(id)
	if len(menus) > 0 && menus[0] == co.BadData {
		return nil, co.ErrNotFound
	}
	if len(menus) == 0 {
		// 从数据库中获取商店信息
		m_menus, err := mysql.GetShopMenuListFromDB(id)
		if err != nil {
			if errors.Is(err, co.ErrNotFound) {
				// 将缓存中的数据设置为空，防止缓存穿透
				err = redis.SetShopMenuList(id, []string{co.BadData}, util.GetRandomExpirationInSeconds(1800, 3600))
				return nil, co.ErrNotFound
			}
			return nil, co.ErrServerBusy
		}
		// 将结构体数组序列化为字符串数组
		da, _ := util.SerializeMenus(m_menus)
		// 将商店信息缓存起来
		err = redis.SetShopMenuList(id, da, util.GetRandomExpirationInSeconds(1800, 3600))
		if err != nil {
			zap.L().Error("redis.SetShopList failed", zap.Error(err))
			return nil, co.ErrServerBusy
		}
		return m_menus, nil
	}
	var data []model.MenuList
	err := util.DeserializeShopMenus(menus, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetShopMenu(id int, idInt int) (model.Menu, error) {
	// 从缓存中获取商店信息
	menu := redis.GetShopMenu(id, idInt)
	if menu == co.BadData {
		return model.Menu{}, co.ErrNotFound
	}
	//zap.L().Info("redis.GetShopMenu", zap.Any("menu", menu))
	if len(menu) == 0 {
		// 从数据库中获取商店信息
		m_menu, err := mysql.GetShopMenuFromDB(id, idInt)
		//zap.L().Info("mysql.GetShopMenuFromDB", zap.Any("m_menu", m_menu))
		if err != nil {
			if errors.Is(err, co.ErrNotFound) {
				// 将缓存中的数据设置为空，防止缓存穿透
				err = redis.SetShopMenu(id, idInt, co.BadData, util.GetRandomExpirationInSeconds(1800, 3600))
				return m_menu, co.ErrNotFound
			}
			return m_menu, co.ErrServerBusy
		}
		// 将结构体序列化为字符串
		da, _ := util.Serialize(m_menu)
		//zap.L().Info("util.Serialize", zap.Any("da", da))
		// 将商店信息缓存起来
		err = redis.SetShopMenu(id, idInt, da, util.GetRandomExpirationInSeconds(1800, 3600))
		if err != nil {
			zap.L().Error("redis.SetShop failed", zap.Error(err))
			return m_menu, co.ErrServerBusy
		}
		return m_menu, nil
	}
	var da model.Menu
	err := util.Deserialize(menu, &da)
	if err != nil {
		return da, err
	}
	return da, nil
}
