package models

import "RestaurantOrder/dao"

type Product struct {
	ID            uint    `gorm:"primarykey"` // 商品ID
	Name          string  // 名称
	Description   string  // 描述
	Price         float64 // 价格
	ProductType   string  // 类型
	Stock         int     // 库存
	Sales         int     //销量
	MerchantEmail string  // 外键
}

// CreateProduct 1.创建一个新的商品信息
func CreateProduct(product *Product) error {
	return dao.DB.Create(&product).Error
}

// GetProductByID 2.根据商品ID查找某一个商品信息
func GetProductByID(productID uint) (*Product, error) {
	var product Product
	err := dao.DB.Where("id = ?", productID).First(&product).Error
	return &product, err
}

// UpdateProductByID 3.根据商品ID更新商品信息
func UpdateProductByID(productID uint) error {
	product, _ := GetProductByID(productID)
	return dao.DB.Where("id = ?", productID).Updates(&product).Error
}

// DeleteProductByID 4.根据商品ID删除商品信息
func DeleteProductByID(productID uint) error {
	product, _ := GetProductByID(productID)
	return dao.DB.Delete(&product).Error
}

// GetAllProduct 5.根据商家email获取所有菜品
func GetAllProduct(email string) (*[]Product, error) {
	var v *[]Product
	err := dao.DB.Where("merchant_email=?", email).Find(&v).Error
	return v, err
}
