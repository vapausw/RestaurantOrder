package mysql

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/model"
	"errors"
	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func CreateMenu(m *model.Menu) error {
	_, err := db.Exec(sqlInsertMenu, m.MenuID, m.ShopID, m.MenuName, m.MenuPrice, m.MenuDesc, m.MenuStock)
	zap.L().Info("Insert menu", zap.String("menu_name", m.MenuName), zap.Error(err))
	if err != nil {
		zap.L().Error("Insert menu failed", zap.String("menu_name", m.MenuName), zap.Error(err))
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			// 已经存在该商品，直接更新即可
			_, err = db.Exec(sqlUpdateMenu, m.MenuName, m.MenuPrice, m.MenuDesc, m.MenuStock, m.MenuID)
			if err != nil {
				zap.L().Error("Update menu failed", zap.String("menu_name", m.MenuName), zap.Error(err))
				return co.ErrServerBusy
			}
			return nil
		}
		return co.ErrServerBusy
	}
	return nil
}
