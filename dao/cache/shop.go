package cache

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/dao/mysql"
	"RestaurantOrder/dao/redis"
	"RestaurantOrder/model"
	"RestaurantOrder/pkg/util"
	"fmt"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

type CacheShop struct {
	model.Shop
}

func (c *CacheShop) RedisGet(key string) (interface{}, error) {
	return redis.GetShop(key)
}

func (c *CacheShop) RedisSet(key string, value interface{}) error {
	return redis.SetShop(key, value.(string), util.GetRandomExpirationInSeconds(1800, 3600))
}

func (c *CacheShop) RedisDel(key string) error {
	return redis.DelShop(key)
}

func (c *CacheShop) MysqlGet(key string) (interface{}, error) {
	return mysql.GetShopFromDB(key)
}

func (c *CacheShop) MysqlSet(key string, value interface{}) error {
	return mysql.UpdateShopToDB(value.(string))
}

func (c *CacheShop) BadData() interface{} {
	return co.BadData
}

func GetShop(id int) (model.Shop, error) {
	key := fmt.Sprintf("%s%s", co.RedisShopKey, strconv.Itoa(id))
	da, err := GetDataFromCache(new(CacheShop), key)
	var data model.Shop
	err = util.Deserialize(da.(string), &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

type CacheShopList struct {
	*model.ShopList
}

func (c *CacheShopList) RedisGet(key string) (interface{}, error) {
	return redis.GetShopList(key)
}

func (c *CacheShopList) RedisSet(key string, value interface{}) error {
	return redis.SetShopList(key, value.([]string), util.GetRandomExpirationInSeconds(1800, 3600))
}

func (c *CacheShopList) RedisDel(key string) error {
	return redis.DelShopList(key)
}

func (c *CacheShopList) MysqlGet(key string) (interface{}, error) {
	return mysql.GetShopListFromDB()
}

func (c *CacheShopList) MysqlSet(key string, value interface{}) error {
	return nil
}

func (c *CacheShopList) BadData() interface{} {
	return []string{co.BadData}
}
func GetShopList() ([]model.ShopList, error) {
	da, err := GetDataFromCache(new(CacheShopList), co.RedisShopListKey)
	var data []model.ShopList
	err = util.DeserializeShops(da.([]string), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type CacheShopMenuList struct {
	*model.MenuList
}

func (c *CacheShopMenuList) RedisGet(key string) (interface{}, error) {
	return redis.GetShopMenuList(key)
}

func (c *CacheShopMenuList) RedisSet(key string, value interface{}) error {
	return redis.SetShopMenuList(key, value.([]string), util.GetRandomExpirationInSeconds(1800, 3600))
}

func (c *CacheShopMenuList) RedisDel(key string) error {
	return nil
}

func (c *CacheShopMenuList) MysqlGet(key string) (interface{}, error) {
	return mysql.GetShopMenuListFromDB(key)
}

func (c *CacheShopMenuList) MysqlSet(key string, value interface{}) error {
	return nil
}

func (c *CacheShopMenuList) BadData() interface{} {
	return []string{co.BadData}
}

func GetShopMenuList(id int) ([]model.MenuList, error) {
	key := strings.Join([]string{co.RedisShopMenuListKey, strconv.Itoa(id)}, "")
	da, err := GetDataFromCache(new(CacheShopMenuList), key)
	if err != nil {
		return nil, err
	}
	var data []model.MenuList
	err = util.DeserializeShopMenus(da.([]string), &data)
	if err != nil {
		zap.L().Error("util.DeserializeShopMenus failed", zap.Error(err))
		return nil, co.ErrServerBusy
	}
	return data, nil
}

// CacheShopMenu 获取一个店家的某一个菜品的详情
type CacheShopMenu struct {
	*model.Menu
}

func (c *CacheShopMenu) RedisGet(key string) (interface{}, error) {
	return redis.GetShopMenu(key)
}

func (c *CacheShopMenu) RedisSet(key string, value interface{}) error {
	return redis.SetShopMenu(key, value.(string), util.GetRandomExpirationInSeconds(1800, 3600))
}

func (c *CacheShopMenu) RedisDel(key string) error {
	return nil
}

func (c *CacheShopMenu) MysqlGet(key string) (interface{}, error) {
	return mysql.GetShopMenuFromDB(key)
}

func (c *CacheShopMenu) MysqlSet(key string, value interface{}) error {
	return nil
}

func (c *CacheShopMenu) BadData() interface{} {
	return co.BadData
}

func GetShopMenu(shopID int, MenuID int) (model.Menu, error) {
	key := strings.Join([]string{co.RedisShopMenuKey, strconv.Itoa(shopID), ":menu:", strconv.Itoa(MenuID)}, "")
	da, err := GetDataFromCache(new(CacheShopMenu), key)
	if err != nil {
		return model.Menu{}, err
	}
	var data model.Menu
	err = util.Deserialize(da.(string), &data)
	if err != nil {
		zap.L().Error("util.Deserialize failed", zap.Error(err))
		return data, co.ErrServerBusy
	}
	return data, nil
}
