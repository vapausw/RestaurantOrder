package logic

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/dao/mysql"
	"RestaurantOrder/model"
)

func AddMenu(m *model.Menu) (error, co.MyCode) {
	// 将当前商品添加到数据库
	if err := mysql.CreateMenu(m); err != nil {
		return err, co.CodeServerBusy
	}
	return nil, co.CodeSuccess
}
