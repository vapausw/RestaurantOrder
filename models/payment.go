package models

import (
	"RestaurantOrder/dao"
	"time"
)

// PaymentInfo 支付信息
type PaymentInfo struct {
	PaymentID     string    // 支付ID
	OrderId       string    // 订单ID
	CustomerEmail string    // 用户邮箱
	CustomerName  string    // 用户名
	MerchantName  string    // 商家名
	MerchantEmail string    // 商家邮箱
	PaymentMethod string    // 支付方式
	PaymentStatus string    // 支付状态
	PaymentTime   time.Time // 支付时间
	Amount        float64   // 支付金额
}

// CreatePaymentInfo 1.创建一个新的支付信息
func CreatePaymentInfo(paymentInfo *PaymentInfo) error {
	return dao.DB.Create(&paymentInfo).Error
}

// GetPaymentInfoByPaymentID 2.根据支付ID查找某一个支付信息
func GetPaymentInfoByPaymentID(paymentID string) (*PaymentInfo, error) {
	var paymentInfo PaymentInfo
	err := dao.DB.Where("payment_id = ?", paymentID).First(&paymentInfo).Error
	return &paymentInfo, err
}

// UpdatePaymentInfoByPaymentID 3.根据支付ID更新支付信息
func UpdatePaymentInfoByPaymentID(paymentID string) error {
	paymentInfo, _ := GetPaymentInfoByPaymentID(paymentID)
	return dao.DB.Where("payment_id = ?", paymentID).Updates(&paymentInfo).Error
}

// DeletePaymentInfoByPaymentID 4.根据支付ID删除支付信息
func DeletePaymentInfoByPaymentID(paymentID string) error {
	paymentInfo, _ := GetPaymentInfoByPaymentID(paymentID)
	return dao.DB.Delete(&paymentInfo).Error
}
