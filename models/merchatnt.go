package models

import (
	"RestaurantOrder/dao"
	"gorm.io/gorm"
	"time"
)

// Merchant 商家信息
type Merchant struct {
	gorm.Model
	Email            string    `gorm:"primarykey"` // 邮箱
	Name             string    `gorm:"unique"`     // 商家名
	Password         string    `gorm:"password"`   // 密码
	Address          string    // 地址
	Phone            string    // 电话
	MerchantType     string    // 类型
	OperatingStatus  string    // 状态
	OperatingHours   string    // 营业时间
	RegistrationTime time.Time // 注册时间
	//Menu             []Product `gorm:"foreignKey:MerchantEmail"` // 关联菜单
}

// 数据库的一些操作

// CreateMerchant 1.创建一个新的商家信息
func CreateMerchant(merchant *Merchant) error {
	return dao.DB.Create(&merchant).Error
}

// GetMerchantByEmail 2.根据邮箱查找某一个商家信息
func GetMerchantByEmail(email string) (*Merchant, error) {
	var merchant Merchant
	err := dao.DB.Where("email = ?", email).First(&merchant).Error
	return &merchant, err
}

// UpdateMerchantByEmail 3.根据邮箱更新商家信息
func UpdateMerchantByEmail(email string) error {
	merchant, _ := GetMerchantByEmail(email)
	return dao.DB.Where("email = ?", email).Updates(&merchant).Error
}

// DeleteMerchantByEmail 4.根据邮箱删除商家信息
func DeleteMerchantByEmail(email string) error {
	merchant, _ := GetMerchantByEmail(email)
	return dao.DB.Delete(&merchant).Error
}
