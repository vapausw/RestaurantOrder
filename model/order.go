package model

type Order struct {
	OrderID int64 `json:"order_id" db:"order_id"`
	UserID  int64 `json:"user_id" db:"user_id"`
	ShopID  int64 `json:"shop_id" db:"shop_id"`
}

type OrderInfo struct {
	OrderID int64  `json:"order_id" db:"order_id"`
	MenuID  int64  `json:"menu_id" db:"menu_id"`
	Count   int    `json:"count" db:"count"`
	Price   int    `json:"price" db:"price"`
	Note    string `json:"note" db:"note"`
}
