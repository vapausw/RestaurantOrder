package mysql

// 存放sql语句
const (
	sqlGetUserPassword    = "SELECT password, user_id FROM users WHERE email = ?"
	sqlInsertUser         = "INSERT INTO users(user_id, email, password, nick_name) VALUES (?,?,?,?)"
	sqlGetShopList        = "SELECT shop_id, shop_name, shop_address from shops"
	sqlGetShop            = "SELECT shop_name, shop_address, shop_phone, shop_desc,shop_id FROM shops WHERE shop_id = ?"
	sqlGetShopMenu        = "SELECT menu_id, menu_name, menu_price, menu_desc, menu_stock FROM menus WHERE shop_id = ? AND menu_id = ?"
	sqlGetShopMenuList    = "SELECT menu_id, menu_name, menu_price FROM menus WHERE shop_id = ?"
	sqlInsertOrder        = "INSERT INTO orders(order_id, user_id, shop_id) VALUES (?,?,?)"
	sqlSelectMerchant     = "SELECT merchant_id, merchant_password FROM merchants WHERE merchant_email = ?"
	sqlInsertMerchant     = "INSERT INTO merchants(merchant_id, merchant_email, merchant_password) VALUES (?,?,?)"
	sqlInsertShop         = "INSERT INTO shops(shop_id, shop_name, shop_address, shop_phone, shop_desc) VALUES (?,?,?,?,?)"
	sqlUpdateShop         = "UPDATE shops SET shop_name = ?, shop_address = ?, shop_phone = ?, shop_desc = ? WHERE shop_id = ?"
	sqlInsertMenu         = "INSERT INTO menus(menu_id, shop_id, menu_name, menu_price, menu_desc, menu_stock) VALUES (?,?,?,?,?,?)"
	sqlUpdateMenu         = "UPDATE menus SET menu_name = ?, menu_price = ?, menu_desc = ?, menu_stock = ? WHERE menu_id = ?"
	sqlUpdateUserPassword = "UPDATE users SET password = ? WHERE email = ?"
)
