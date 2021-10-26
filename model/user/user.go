package user


type User struct {
	User_id     int `gorm:"primaryKey"`
	User_name   string
	User_email  string
	User_type   int
	User_phone  string
	User_secret string
}


type UserResp struct {
	ID     int    `json:"user_id"`
	Name   string `json:"user_name"`
	Email  string `json:"user_email"`
	Type   int    `json:"user_type"`
	Phone  string `json:"user_phone"`
}

type UserReq struct {
	Modify_field string `json:"modify_field" binding:"required"`
	User_name    string `json:"user_name"`
	User_email   string `json:"user_email"`
	User_type    int    `json:"user_type"`
	User_phone   string `json:"user_phone"`
}

type AuthReq struct {
	Account string `json:"account" form:"account" binding:"required"`
	Secret  string `json:"secret" form:"secret" binding:"required"`
}

type TokenAuth struct {
	Token string `json:"token" form:"token" binding:"required"`
}
