package model

type Shop struct { // 商店的信息
	ShopId    int    `db:"shop_id" json:"shop_id"`
	ShopName  string `db:"shop_name" json:"shop_name"`
	ShopDesc  string `db:"shop_desc" json:"shop_desc"`
	ShopAddr  string `db:"shop_addr" json:"shop_addr"`
	ShopPhone string `db:"shop_phone" json:"shop_phone"`
	//ShopImage string `db:"shop_image"`
}

type ShopList struct { // 用于返回给前端商户列表
	// 存储商户ID与名称，地址即可用于前端展示，目前尚未实现商户图片的上传，有图片的话也可以加上
	ShopId   int    `json:"shop_id" db:"shop_id"`
	ShopName string `json:"shop_name" db:"shop_name"`
	ShopAddr string `json:"shop_addr" db:"shop_addr"`
}
