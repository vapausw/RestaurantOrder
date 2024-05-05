package model

type Merchant struct {
	MerchantId int64  `db:"merchant_id"`
	Email      string `json:"email" db:"merchant_email"`
	Password   string `json:"password" db:"merchant_password"`
	RePassword string `json:"re_password"`
	Code       string `json:"code"`
}
