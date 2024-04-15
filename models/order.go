package models

import (
	"RestaurantOrder/dao"
	"time"
)

// Order 订单信息
type Order struct {
	OrderID       string      `gorm:"type:varchar(255);primaryKey"` // 订单ID
	MerchantName  string      `json:"merchant_name"`                // 商家名
	CustomerName  string      `json:"customer_name"`                // 用户名
	CustomerEmail string      `json:"customer_email"`               // 用户邮箱
	MerchantEmail string      `json:"merchant_email"`               // 商家邮箱
	Status        string      `json:"status"`                       // 订单状态
	OrderTime     time.Time   // 下单时间
	Items         []OrderItem `gorm:"foreignKey:OrderID"` // 订单项，指定外键
	TotalPrice    float64     // 订单总价
	PaymentMethod string      // 支付方式
}

// OrderItem 订单项
type OrderItem struct {
	OrderItemID uint    `gorm:"primaryKey"` // 订单项ID
	Name        string  `gorm:"type:text"`  // 商品名称
	ProductID   uint    // 商品ID
	Quantity    int     // 商品数量
	UnitPrice   float64 // 商品单价
	OrderID     string  `gorm:"foreignKey:OrderID"` // 订单ID
}

// CreateOrder 1.创建一个新的订单信息
func CreateOrder(order *Order) error {
	return dao.DB.Create(&order).Error
}

// GetOrderByOrderID 2.根据订单ID查找某一个订单信息
func GetOrderByOrderID(orderID string) (*Order, error) {
	var order Order
	err := dao.DB.Where("order_id = ?", orderID).First(&order).Error
	return &order, err
}

// UpdateOrderByOrderID 3.根据订单ID更新订单信息
func UpdateOrderByOrderID(orderID string) error {
	order, _ := GetOrderByOrderID(orderID)
	return dao.DB.Where("order_id = ?", orderID).Updates(&order).Error
}

// DeleteOrderByOrderID 4.根据订单ID删除订单信息
func DeleteOrderByOrderID(orderID string) error {
	order, _ := GetOrderByOrderID(orderID)
	return dao.DB.Delete(&order).Error
}

// GetOrdersByCustomerEmail 5.根据用户邮箱查找所有订单信息
func GetOrdersByCustomerEmail(customerEmail string) ([]Order, error) {
	var orders []Order
	err := dao.DB.Where("customer_email = ?", customerEmail).Find(&orders).Error
	return orders, err
}

// GetOrdersByMerchantEmail 6.根据商家邮箱查找所有订单信息
func GetOrdersByMerchantEmail(merchantEmail string) ([]Order, error) {
	var orders []Order
	err := dao.DB.Where("merchant_email = ?", merchantEmail).Find(&orders).Error
	return orders, err
}
