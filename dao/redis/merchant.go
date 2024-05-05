package redis

import (
	co "RestaurantOrder/constant"
	"fmt"
	"time"
)

func SetMerchantCode(email string, code string) error {
	key := fmt.Sprintf("%s%s", co.RedisMerchantCodeKey, email)
	return rdb.Set(key, code, time.Minute*5).Err()
}

func GetMerchantCode(email string) (string, error) {
	key := fmt.Sprintf("%s%s", co.RedisMerchantCodeKey, email)
	return rdb.Get(key).Result()
}
