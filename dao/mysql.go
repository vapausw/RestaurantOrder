package dao

import (
	"RestaurantOrder/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitMySQL(cfg *setting.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}

func MySQLClose() error {
	sqlDB, _ := DB.DB()
	return sqlDB.Close()
}
