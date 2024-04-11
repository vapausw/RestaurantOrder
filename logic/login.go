package logic

import (
	"RestaurantOrder/dao"
	"RestaurantOrder/models"
	"RestaurantOrder/utils"
	"errors"
	"time"
)

func LoginCheck(email, password, modelsType string) (err error) {
	if email == "" || password == "" {
		return errors.New("email and password are required")
	}

	if modelsType == "customer" {
		v, _ := models.GetCustomerByEmail(email)
		if !utils.CheckPasswordHash(password, v.Password) {
			return errors.New("incorrect password or email")
		}
	} else {
		v, _ := models.GetMerchantByEmail(email)
		if !utils.CheckPasswordHash(password, v.Password) {
			return errors.New("incorrect password or email")
		}
	}
	return nil
}

func RegisterSendCode(email string) error {
	token := utils.GenerateSecureToken()
	dao.Rdb.Set(email, token, 10*time.Minute) // 此令牌10分钟有效
	return utils.SendEmail(email, token)
}

func RegisterCheck(email, token, password, repeatPassword, name, modelsType string) error {
	if password != repeatPassword {
		return errors.New("passwords do not match")
	}
	if token != dao.Rdb.Get(email).Val() {
		return errors.New("incorrect token")
	}
	// 将密码加密后将数据保存到mysql数据库中
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	switch modelsType {
	case "customer":
		customer := models.Customer{
			Name:     name,
			Email:    email,
			Password: hashedPassword,
		}
		models.CreateCustomer(&customer)
	case "merchant":
		merchant := models.Merchant{
			Name:     name,
			Email:    email,
			Password: hashedPassword,
		}
		models.CreateMerchant(&merchant)
	default:
		return errors.New("type assertion failed")
	}
	return nil
}
