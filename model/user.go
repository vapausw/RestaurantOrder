package model

type User struct {
	UserId     int    `db:"user_id"`
	UserName   string `db:"nick_name"`
	Password   string `db:"password" json:"password"`
	RePassword string `json:"re_password"`
	Email      string `db:"email" json:"email"`
	Code       string `db:"code" json:"code"`
}

type UserLogin struct {
	UserId   int    `db:"user_id"`
	Password string `db:"password"`
}
