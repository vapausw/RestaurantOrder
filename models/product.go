package models

import "RestaurantOrder/dao"

type Product struct {
	ID            uint    `gorm:"primarykey" json:"id"` // 商品ID
	Name          string  `json:"name"`                 // 名称
	Description   string  `json:"description"`          // 描述
	Price         float64 `json:"price"`                // 价格
	ProductType   string  `json:"productType"`          // 类型
	Stock         int     `json:"stock"`                // 库存
	Sales         int     //销量
	MerchantEmail string  // 外键
}

// CreateProduct 1.创建一个新的商品信息
func CreateProduct(product *Product) error {
	// 将商品名和销量存入以商家主键命名的排序集合
	//dao.Rdb.ZAdd(product.MerchantEmail, &dao.Z{Score: float64(product.Sales), Member: product.Name})
	return dao.DB.Create(&product).Error
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
