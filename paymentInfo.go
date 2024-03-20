package main

import "time"

// PaymentInfo 支付信息
type PaymentInfo struct {
	PaymentID     string    // 支付ID
	OrderId       string    // 订单ID
	PaymentMethod string    // 支付方式
	PaymentStatus string    // 支付状态
	PaymentTime   time.Time // 支付时间
	Amount        float64   // 支付金额
}
