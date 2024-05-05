package logic

//
//func LoginCheck(email, password, modelsType string) (err error) {
//	if email == "" || password == "" {
//		return EmailOrPasswordNUll
//	}
//
//	if modelsType == "customer" {
//		v, _ := mysql.GetCustomerByEmail(email) // 从数据库中获取用户信息
//		if !utils.CheckPasswordHash(password, v.Password) {
//			return EmailOrPasswordWrong
//		}
//	} else {
//		v, _ := mysql.GetMerchantByEmail(email)
//		if !utils.CheckPasswordHash(password, v.Password) {
//			return EmailOrPasswordWrong
//		}
//	}
//	return nil
//}
//
//func RegisterSendCode(email string) error {
//	token := utils.GenerateSecureToken()
//	redis.Set(email, token, 10*time.Minute) // 此令牌10分钟有效
//	return utils.SendEmail(email, token)
//}
//
//func RegisterCheck(repeatPassword, token string, v interface{}) error {
//	switch x := v.(type) {
//	case models.Customer:
//		if x.Password != repeatPassword {
//			return EquivalentPassword
//		}
//		if token != redis.Get(x.Email).Val() {
//			return ErrorCaptcha
//		}
//		// 将密码加密后将数据保存到mysql数据库中
//		hashedPassword, err := utils.HashPassword(x.Password)
//		if err != nil {
//			return err
//		}
//		x.Password = hashedPassword
//		x.ID = snowflake.GenID()
//		err = mysql.CreateCustomer(&x)
//		if err != nil {
//			return err
//		}
//	case models.Merchant:
//		if x.Password != repeatPassword {
//			return EquivalentPassword
//		}
//		if token != redis.Get(x.Email).Val() {
//			return ErrorCaptcha
//		}
//		// 将密码加密后将数据保存到mysql数据库中
//		hashedPassword, err := utils.HashPassword(x.Password)
//		if err != nil {
//			return err
//		}
//		x.Password = hashedPassword
//		x.ID = snowflake.GenID()
//		err = mysql.CreateMerchant(&x)
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}
