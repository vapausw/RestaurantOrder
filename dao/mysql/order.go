package mysql

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/model"
	"errors"
	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func CreateOrder(order *model.Order) error {

	if _, err := db.Exec(sqlInsertOrder, order.OrderID, order.UserID, order.ShopID); err != nil {
		zap.L().Error("Insert order failed", zap.Int64("order_id", order.OrderID), zap.Error(err))
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return co.ErrExistsData
		}
		return co.ErrServerBusy
	}
	return nil
}
