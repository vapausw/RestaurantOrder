package constant

const (
	RedisUserPasswordKey = "RestaurantOrder:user:password:"  // 存储用户哈希密码
	RedisUserCodeKey     = "RestaurantOrder:user:code:"      // 存储验证码
	RedisMerchantCodeKey = "RestaurantOrder:merchant:code:"  // 存储商家验证码
	RedisShopListKey     = "RestaurantOrder:shop:list"       // 存储商家信息链表
	RedisShopKey         = "RestaurantOrder:shop:"           // 存储商家信息
	RedisShopMenuListKey = "RestaurantOrder:shop:menu:list:" // 存储菜品信息链表
	RedisShopMenuKey     = "RestaurantOrder:shop:"           // 存储菜品信息
	RedisCartKey         = "RestaurantOrder:cart:"           // 存储用户的购物车信息
	RedisBaseChar        = "%"                               // 用于拼接一些特殊的key，例如用户id与菜品id的拼接
	RedisOrderInfoKey    = "RestaurantOrder:order:info:"     // 存储订单信息
)

const (
	BadData = "CachePenetration"
)
