package main

import (
	"time"
)

// Customer 用户信息
type Customer struct {
	CustomerID  string    // 用户唯一标识
	Name        string    // 用户名
	Email       string    // 邮箱，可用于登录和找回密码
	Password    string    // 加密后的密码
	IsVIP       bool      // 是否为会员
	ResetToken  string    // 重置密码令牌
	TokenExpiry time.Time // 重置密码令牌过期时间
}

var testcustomers Customers // 此处用于测试，实际应用中应该使用数据库存储用户信息

type Customers struct {
	mp map[string]Customer
}

func RegisterUser(name, email, password string) (Customer, error) {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return Customer{}, err
	}
	user := Customer{
		CustomerID: CreateID(), // 生成用户ID的函数
		Name:       name,
		Email:      email, // 其实邮件按理说也是唯一，但此处目前不添加邮件发送令牌检查
		Password:   hashedPassword,
		IsVIP:      false,
	}
	// 存储 user 到数据库
	testcustomers.mp[email] = user
	return user, nil
}

func Login(email, password string) {

}
