package cache

import (
	co "RestaurantOrder/constant"
	"errors"
)

// Cache 实现一个缓存接口
type Cache interface {
	RedisSet(key string, value interface{}) error
	RedisGet(key string) (interface{}, error)
	RedisDel(key string) error
	MysqlGet(key string) (interface{}, error)
	MysqlSet(key string, value interface{}) error
	BadData() interface{}
}

//  定义一些通用的缓存方法，用于支撑redis和mysql的服务

// GetDataFromCache 获取数据
func GetDataFromCache(c Cache, key string) (interface{}, error) {
	//var mysqlSync sync.Mutex
	// 1.先从redis中获取，如果获取到了，直接返回
	// 2.如果没有获取到，从mysql中获取
	// 3.将mysql中获取到的数据存储到redis中，mysql中没有数据的话，redis中存储一个提前设置好的坏数据
	// 4.返回数据
	data, err := c.RedisGet(key)
	if errors.Is(err, co.ErrNotFound) {
		// 从mysql中获取数据，加锁，防止没有缓存期间全部请求都去mysql中获取数据然后进行更新
		data, err = c.MysqlGet(key)
		if errors.Is(err, co.ErrNotFound) {
			// 将缓存中的数据设置为提前设置好的坏数据，防止缓存穿透
			err = c.RedisSet(key, c.BadData())
			if err != nil {
				return nil, co.ErrServerBusy
			}
			return nil, co.ErrNotFound
		} else if err != nil {
			return nil, co.ErrServerBusy
		}
		err = c.RedisSet(key, data)
		if err != nil {
			return nil, co.ErrServerBusy
		}
		return data, nil
	} else if errors.Is(err, co.ErrBadRedisData) {
		return nil, co.ErrNotFound
	} else if err != nil {
		return nil, co.ErrServerBusy
	}
	return data, nil
}

// UpdateCache 实现一个更新缓存的方法，保证数据的一致性
func UpdateCache(c Cache, key string, value interface{}) error {
	// 1.将数据先更新到mysql中
	// 2.然后删除redis缓存中的数据
	err := c.MysqlSet(key, value)
	if err != nil {
		return co.ErrServerBusy
	}
	err = c.RedisDel(key)
	if err != nil {
		return co.ErrServerBusy
	}
	return nil
}
