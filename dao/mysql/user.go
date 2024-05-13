package mysql

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/model"
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func GetUserPasswordFromDB(email string) (user *model.User, err error) {
	user = new(model.User)
	// 查询数据库
	err = db.QueryRow(sqlGetUserPassword, email).Scan(&user.Password, &user.UserId, &user.UserName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, co.ErrNotFound
		}
		zap.L().Error("db.QueryRow failed", zap.Error(err))
		return nil, co.ErrServerBusy
	}
	return
}

func CreateUser(u *model.User) error {
	_, err := db.Exec(sqlInsertUser, u.UserId, u.Email, u.Password, u.UserName)
	if err != nil {
		zap.L().Error("Insert user failed", zap.String("username", u.Email), zap.Error(err))
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return co.ErrExistsUser
		}
		return co.ErrServerBusy
	}
	return nil
}

func UpdateUserPassword(key string, login *model.User) error {
	_, err := db.Exec(sqlUpdateUserPassword, login.Password, key)
	if err != nil {
		zap.L().Error("Update user password failed", zap.String("email", key), zap.Error(err))
		return co.ErrServerBusy
	}
	return nil
}

func GetUserInfo(id int64) (user model.User, err error) {
	err = db.QueryRow(sqlGetUserInfo, id).Scan(&user.UserName, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, co.ErrNotFound
		}
		zap.L().Error("db.QueryRow failed", zap.Error(err))
		return user, co.ErrServerBusy
	}
	return
}
