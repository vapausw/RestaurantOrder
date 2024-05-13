package model

type User struct {
	UserId     int    `db:"user_id"`
	UserName   string `db:"nick_name" json:"userName"`
	Password   string `db:"password" json:"password"`
	RePassword string `json:"rePassword"`
	Email      string `db:"email" json:"email"`
	Code       string `json:"code"`
}

type UserWebsocket struct {
	UserId   int64  `json:"user_id"`
	Order_id string `json:"order_id"`
}
