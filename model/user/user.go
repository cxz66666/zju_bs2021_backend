package user


type User struct {
	UserId   int `gorm:"primaryKey"`
	UserName string  `gorm:"uniqueIndex;size:30"`
	UserEmail string `gorm:"uniqueIndex;size:30"`
	UserType  Role
	UserPhone  string `gorm:"size:20"`
	UserSecret string
}

type UserInfoResp struct {
	ID     int    `json:"user_id"`
	Name   string `json:"user_name"`
	Email  string `json:"user_email"`
	Type   Role    `json:"user_type"`
	Phone  string `json:"user_phone"`
}

type UserModifyReq struct {
	ModifyField string `json:"modify_field" binding:"required"`
	UserName  string `json:"user_name"`
	UserEmail  string `json:"user_email"`
	UserType  Role   `json:"user_type"`
	UserPhone string `json:"user_phone"`
}

type UserCreateReq struct {
	UserName  string `json:"user_name" form:"user_name" binding:"required,max=30"`
	UserEmail string `json:"user_email" form:"user_email" binding:"required,max=30"`
	UserSecret string `json:"user_secret" form:"user_secret" binding:"required,max=20"`
	UserPhone string `json:"user_phone" form:"user_phone" binding:"required,max=20"`
}


type AuthReq struct {
	Email  string `json:"email" form:"email" binding:"required"`
	Secret string `json:"secret" form:"secret" binding:"required"`
}

type AuthResq struct {
	UserName  string `json:"user_name"`
	UserEmail  string `json:"user_email"`
	UserType  Role   `json:"user_type"`
	UserToken string `json:"user_token"`
}




type TokenAuth struct {
	Token string `json:"token" form:"token" binding:"required"`
}
