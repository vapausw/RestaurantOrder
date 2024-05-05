package model

type CartInfo struct {
	// 商品id, 商品数量，商品价格， 商品备注信息
	// 用户id，方便后续查询
	MenuID int64  `json:"menu_id" db:"menu_id"`
	Count  int    `json:"count" db:"count"`
	Price  int    `json:"price" db:"price"`
	Note   string `json:"note" db:"note"`
	ShopID int64  `json:"shop_id" db:"shop_id"`
	UserID int64  `db:"user_id"`
}
