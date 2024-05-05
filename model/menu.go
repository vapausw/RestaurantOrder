package model

type Menu struct { // 商家菜品详情
	MenuID    int     `json:"menu_id" db:"menu_id"`
	ShopID    int     `json:"shop_id" db:"shop_id"`
	MenuName  string  `json:"menu_name" db:"menu_name"`
	MenuDesc  string  `json:"menu_desc" db:"menu_desc"`
	MenuPrice float64 `json:"menu_price" db:"menu_price"`
	MenuStock int     `json:"menu_stock" db:"menu_stock"`
	//MenuImage string `json:"menu_image" db:"menu_image"`
}

type MenuList struct { // 商家菜品列表
	MenuId    int     `json:"menu_id" db:"menu_id"`
	MenuName  string  `json:"menu_name" db:"menu_name"`
	MenuPrice float64 `json:"menu_price" db:"menu_price"`
}
