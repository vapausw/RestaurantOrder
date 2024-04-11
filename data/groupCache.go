package data

import (
	"RestaurantOrder/dao"
	"RestaurantOrder/models"
	"encoding/json"
	"github.com/golang/groupcache"
)

var (
	// CustomerCache 定义一个 Customer的类型的缓存
	CustomerCache = groupcache.NewGroup("CustomerInfoCache", 64<<20, groupcache.GetterFunc(
		func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
			// 从数据库或其他数据源加载数据
			customer, err := loadDataFromDataSource(key)
			if err != nil {
				return err
			}
			// 序列化数据
			data, err := SerializeUserInfo(customer)
			// 将数据写入 dest
			err = dest.SetBytes(data)
			if err != nil {
				return err
			}
			return nil
		}))
	// OrderCache 定义一个 Order的类型的缓存,方便获取一个用户的所有订单信息
	OrderCache = groupcache.NewGroup("OrderInfoCache", 64<<20, groupcache.GetterFunc(
		func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
			// 从数据库或其他数据源加载数据
			orders, err := loadOrderDataFromDataSource(key)
			if err != nil {
				return err
			}
			// 序列化数据
			data, err := SerializeUserInfo(orders)
			// 将数据写入 dest
			err = dest.SetBytes(data)
			if err != nil {
				return err
			}
			return nil
		}))
)

func loadOrderDataFromDataSource(customerEmail string) ([]*models.Order, error) {
	var orders []*models.Order
	// 假设每个订单都有一个字段指向顾客的email
	dao.DB.Where("customer_email = ?", customerEmail).Find(&orders)
	return orders, nil
}

func loadDataFromDataSource(key string) (*models.Customer, error) {
	var customer models.Customer
	dao.DB.Where("email = ?", key).First(&customer)
	return &customer, nil
}

// SerializeUserInfo 将UserInfo对象序列化为字节切片
func SerializeUserInfo(userInfo interface{}) ([]byte, error) {
	return json.Marshal(userInfo)
}

// DeserializeUserInfo 将字节切片反序列化为UserInfo对象
func DeserializeUserInfo(data []byte) (*models.Customer, error) {
	var userInfo models.Customer
	err := json.Unmarshal(data, &userInfo)
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}

// DeserializeOrderInfo 将字节切片反序列化为Order对象列表
func DeserializeOrderInfo(data []byte) ([]*models.Order, error) {
	var orders []*models.Order
	err := json.Unmarshal(data, &orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func GetCustomerInfoFromCache(email string) (*models.Customer, error) {
	var data []byte
	err := CustomerCache.Get(nil, email, groupcache.AllocatingByteSliceSink(&data))
	if err != nil {
		return nil, err
	}
	// 将字节序列反序列化回UserInfo结构体
	customerInfo, err := DeserializeUserInfo(data)
	if err != nil {
		return nil, err
	}

	return customerInfo, nil
}

func GetOrderInfoFromCache(Email string) ([]*models.Order, error) {
	var data []byte
	err := OrderCache.Get(nil, Email, groupcache.AllocatingByteSliceSink(&data))
	if err != nil {
		return nil, err
	}
	// 将字节序列反序列化回UserInfo结构体
	orders, err := DeserializeOrderInfo(data)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
