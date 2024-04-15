package models

import (
	"RestaurantOrder/dao"
	"github.com/go-redis/redis"
)

type Product struct {
	ID            uint    `gorm:"primarykey" json:"id"` // 商品ID
	Name          string  `json:"name"`                 // 名称
	Description   string  `json:"description"`          // 描述
	Price         float64 `json:"price"`                // 价格
	ProductType   string  `json:"productType"`          // 类型
	Stock         int     `json:"stock"`                // 库存
	Sales         int     //销量
	MerchantEmail string  `gorm:"index"` // 指明这是外键/
}

// CreateProduct 1.创建一个新的商品信息
func CreateProduct(product *Product) error {
	// 首先, 将商品信息存入数据库
	err := dao.DB.Create(&product).Error
	if err != nil {
		return err
	}

	// 使用商家的邮箱作为key, 将商品名和销量存入Redis的排序集合
	// 这里假设 `dao.Rdb` 是Redis连接的实例, `Z` 是结构体用于表示成员和分数
	score := float64(product.Sales) // 将销量作为排序分数
	member := product.Name          // 商品名作为成员
	key := product.MerchantEmail    // 使用商家邮箱作为排序集名称

	// 将商品名和销量加入到Redis排序集合
	dao.Rdb.ZAdd(key, redis.Z{Score: score, Member: member})
	if err != nil {
		return err
	}
	return nil
}

// GetProductByID 2.根据商品ID查找某一个商品信息
func GetProductByID(productID int) (*Product, error) {
	var product Product
	err := dao.DB.Where("id = ?", productID).First(&product).Error
	return &product, err
}

// UpdateProduct 3.更新商品信息
func UpdateProduct(p *Product) error {
	err := dao.DB.Save(p).Error
	return err
}

// DeleteProductByID 4.根据商品ID删除商品信息
func DeleteProductByID(productID int) error {
	product, _ := GetProductByID(productID)
	return dao.DB.Delete(&product).Error
}

// GetAllProduct 5.根据商家email获取所有菜品
func GetAllProduct(email string) (*[]Product, error) {
	var v *[]Product
	err := dao.DB.Where("merchant_email=?", email).Find(&v).Error
	return v, err
}
