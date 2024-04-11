package models

import "RestaurantOrder/dao"

// Customer 定义用户模型
type Customer struct {
	Email    string `gorm:"primarykey"` // 邮箱，可用于登录和找回密码，也作为用户唯一标识
	Name     string `gorm:"name"`       // 用户名
	Password string `gorm:"password"`   // 加密后的密码
}

//数据库的一些操作

// CreateCustomer 1.创建一个新的用户信息
func CreateCustomer(customer *Customer) error {
	return dao.DB.Create(&customer).Error
}

// GetCustomerByEmail 2.根据邮箱查找某一个用户信息
func GetCustomerByEmail(email string) (*Customer, error) {
	var customer Customer
	err := dao.DB.Where("email = ?", email).First(&customer).Error
	return &customer, err
}

// UpdateCustomerByEmail 3.根据邮箱更新用户信息
func UpdateCustomerByEmail(email string) error {
	customer, _ := GetCustomerByEmail(email)
	return dao.DB.Where("email = ?", email).Updates(&customer).Error
}

// DeleteCustomerByEmail 4.根据邮箱删除用户信息
func DeleteCustomerByEmail(email string) error {
	customer, _ := GetCustomerByEmail(email)
	return dao.DB.Delete(&customer).Error
}
