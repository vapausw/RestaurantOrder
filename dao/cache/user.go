package cache

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/dao/mysql"
	"RestaurantOrder/dao/redis"
	"RestaurantOrder/model"
	"RestaurantOrder/pkg/util"
)

// 使用redis缓存，减轻数据库的压力，提高性能
// 防止缓存穿透，缓存击穿，缓存雪崩

// 1. 缓存穿透：查询一个一定不存在的数据，由于缓存不命中，每次都要去查询数据库，导致数据库压力过大
// 解决方案：将数据库中存在的数据缓存起来，将不存在的数据也缓存起来，但是设置一个较短的过期时间

// 2. 缓存击穿：一个key非常热点，在某个时间点，这个key过期了，此时有大量的请求并发访问这个key，导致数据库压力过大
// 解决方案：设置热点数据永不过期，或者加锁，让其他请求等待，直到第一个请求完成数据查询

// 3. 缓存雪崩：大量的key在同一时间过期，导致大量的请求直接访问数据库
// 解决方案：设置不同的过期时间，让缓存失效的时间点尽量均匀

type CacheUser struct {
	model.User
}

func (c *CacheUser) RedisSet(key string, value interface{}) error {
	return redis.SetUserPassword(key, value.(*model.User), util.GetRandomExpirationInSeconds(300, 600))
}
func (c *CacheUser) RedisGet(key string) (interface{}, error) {
	return redis.GetUserPassword(key)
}
func (c *CacheUser) RedisDel(key string) error {
	return redis.DelUserPassword(key)
}
func (c *CacheUser) MysqlGet(key string) (interface{}, error) {
	return mysql.GetUserPasswordFromDB(key)
}
func (c *CacheUser) MysqlSet(key string, value interface{}) error {
	return mysql.UpdateUserPassword(key, value.(*model.User))
}
func (c *CacheUser) BadData() interface{} {
	return &model.User{
		Password: co.BadData,
	}
}

// GetUserPassword 获取用户加密的密码用于登录验证
func GetUserPassword(email string) (user *model.User, err error) {
	data, err := GetDataFromCache(new(CacheUser), email)
	if err != nil {
		return nil, co.ErrServerBusy
	}
	user = data.(*model.User)
	return
}

//func GetUserPassword(email string) (user *model.UserLogin, err error) {
//	// 从缓存中获取用户密码
//	user, err = redis.GetUserPassword(email)
//	//zap.L().Info("user", zap.Any("user", user))
//	//zap.L().Info("err", zap.Any("err", err))
//	if err != nil && !errors.Is(err, co.ErrNotFound) {
//		zap.L().Error("redis.GetUserPassword failed", zap.Error(err))
//		return nil, co.ErrServerBusy
//	}
//	if user.Password == co.BadData {
//		return nil, co.ErrNotFound
//	}
//	// 如果缓存中没有，从数据库中获取
//	if len(user.Password) == 0 {
//		// 从数据库中获取用户密码
//		user, err = mysql.GetUserPasswordFromDB(email)
//		//zap.L().Info("err", zap.Any("err", err))
//		if err == nil {
//			// 将用户密码缓存起来
//			err = redis.SetUserPassword(email, user, util.GetRandomExpirationInSeconds(300, 600))
//			if err != nil {
//				zap.L().Error("redis.SetUserPassword failed", zap.Error(err))
//				return nil, co.ErrServerBusy
//			}
//		} else if user.Password == "" && errors.Is(err, co.ErrNotFound) { // 数据库中也没有将其设置为nil，防止缓存穿透
//			user.Password = co.BadData
//			err = redis.SetUserPassword(email, user, util.GetRandomExpirationInSeconds(300, 600))
//			if err != nil {
//				zap.L().Error("redis.SetUserPassword failed", zap.Error(err))
//				return nil, co.ErrServerBusy
//			}
//			err = nil
//		}
//		zap.L().Error("redis get failed", zap.Error(err))
//	}
//	return
//}
