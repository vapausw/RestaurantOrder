package logic

//// PlaceOrder 用户点餐生成订单
//func PlaceOrder(db *gorm.DB, customerName, paymentMethod string, cartItems []CartItem) (Order, error) {
//	// CartItem 是前端发送的一个包含菜品ID和数量的结构体
//	var items []orderItem
//	var total float64
//
//	// 遍历购物车项目，转换为订单项
//	for _, cartItem := range cartItems {
//		var product Product
//		if err := db.Where("id = ?", cartItem.ProductID).First(&product).Error; err != nil {
//			return Order{}, err // 未找到菜品
//		}
//
//		// 创建订单项
//		item := orderItem{
//			Name:      product.Name,
//			Quantity:  cartItem.Quantity,
//			UnitPrice: product.Price,
//		}
//
//		items = append(items, item)
//		total += item.UnitPrice * float64(item.Quantity)
//	}
//
//	// 创建订单
//	order, err := CreateOrder(db, customerName, paymentMethod, items)
//	if err != nil {
//		return Order{}, err
//	}
//
//	// 可选：更新订单总价（如果CreateOrder中没有计算）
//	order.TotalPrice = total
//	if err := db.Save(&order).Error; err != nil {
//		return Order{}, err
//	}
//
//	return order, nil
//}
//
//func CreateOrder(db *gorm.DB, customerName, paymentMethod string, items []orderItem) (Order, error) {
//	order := Order{
//		OrderID:       CreateID(), // 生成或指定一个唯一的订单ID
//		CustomerName:  customerName,
//		Status:        "Pending", // 初始状态可能是"Pending"
//		OrderTime:     time.Now(),
//		Items:         items,
//		PaymentMethod: paymentMethod,
//	}
//
//	// 计算总价
//	var totalPrice float64
//	for _, item := range items {
//		totalPrice += item.UnitPrice * float64(item.Quantity)
//	}
//	order.TotalPrice = totalPrice
//
//	// 存储到数据库
//	if err := db.Create(&order).Error; err != nil {
//		return Order{}, err
//	}
//
//	return order, nil
//}
//
//func UpdateOrderStatus(db *gorm.DB, orderID, status string) error {
//	var order Order
//	if err := db.Where("order_id = ?", orderID).First(&order).Error; err != nil {
//		return err
//	}
//
//	order.Status = status
//	return db.Save(&order).Error
//}
