package mysql

import (
	"RestaurantOrder/setting"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

var db *sql.DB

func Init(config *setting.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username, config.Password, config.Host, config.Port, config.DbName)
	db, err = sql.Open("mysql", dsn)
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	if err != nil {
		return
	}
	err = db.Ping()
	return
}

func Close() {
	if err := db.Close(); err != nil {
		zap.L().Error("mysql close error", zap.Error(err))
		panic(err)
	}
}
