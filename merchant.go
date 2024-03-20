package main

import "time"

// Merchant 商家信息
type Merchant struct {
	MerchantID       string    // 商家ID
	Name             string    // 商家名
	Email            string    // 邮箱，可用于登录和找回密码
	Password         string    // 加密后的密码
	Address          string    // 商家地址
	Phone            string    // 商家电话
	MerchantType     string    // 商家类型
	OperatingStatus  string    // 商家状态
	OperatingHours   string    // 营业时间
	Rating           float64   // 评分
	RegistrationTime time.Time // 注册时间
}
