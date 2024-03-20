package main

import "time"

type orderItem struct { // 订单项
	ProductId string  // 商品ID
	Quantity  int     // 商品数量
	UnitPrice float64 // 商品单价
}
type Order struct {
	OrderID       string      // 订单ID
	CustomerId    string      // 用户ID
	Status        string      // 订单状态
	OrderTime     time.Time   // 下单时间
	Items         []orderItem // 订单项
	TotalPrice    float64     // 订单总价
	PaymentMethod string      // 支付方式
}
