package mysql

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/model"
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func GetMerchantPassword(m *model.Merchant) error {
	if err := db.QueryRow(sqlSelectMerchant, m.Email).Scan(&m.MerchantId, &m.Password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return co.ErrNotFound
		}
		zap.L().Error("GetMerchantPassword db.QueryRow failed", zap.Error(err))
		return co.ErrServerBusy
	}
	return nil
}

func CreateMerchant(m *model.Merchant) error {
	_, err := db.Exec(sqlInsertMerchant, m.MerchantId, m.Email, m.Password)
	if err != nil {
		zap.L().Error("Insert merchant failed", zap.String("username", m.Email), zap.Error(err))
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return co.ErrExistsUser
		}
		return co.ErrServerBusy
	}
	return nil
}

func CreateShop(m *model.Shop) error {
	_, err := db.Exec(sqlInsertShop, m.ShopId, m.ShopName, m.ShopAddr, m.ShopPhone, m.ShopDesc)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			// 直接更新数据即可
			_, err = db.Exec(sqlUpdateShop, m.ShopName, m.ShopAddr, m.ShopPhone, m.ShopDesc)
			if err != nil {
				zap.L().Error("Update shop failed", zap.String("shopName", m.ShopName), zap.Error(err))
				return co.ErrServerBusy
			}
			return nil
		}
		zap.L().Error("Insert shop failed", zap.String("shopName", m.ShopName), zap.Error(err))
		return co.ErrServerBusy
	}
	return nil

}
