package user

type User struct {
	UserId     int    `gorm:"primaryKey"`
	UserName   string `gorm:"uniqueIndex;size:30"`
	UserEmail  string `gorm:"uniqueIndex;size:30"`
	UserType   Role
	UserPhone  string `gorm:"size:20"`
	UserSecret string
}

type UserInfoResp struct {
	ID    int    `json:"userId"`
	Name  string `json:"userName"`
	Email string `json:"userEmail"`
	Type  Role   `json:"userType"`
	Phone string `json:"userPhone"`
}

type UserModifyReq struct {
	UserName  string `json:"userName" binding:"required"`
	UserEmail string `json:"userEmail" binding:"required"`
	UserPhone string `json:"userPhone" binding:"required"`
}

type UserCreateReq struct {
	UserName   string `json:"userName" form:"userName" binding:"required,max=30"`
	UserEmail  string `json:"userEmail" form:"userEmail" binding:"required,max=30"`
	UserSecret string `json:"userSecret" form:"userSecret" binding:"required,max=20"`
	UserPhone  string `json:"userPhone" form:"userPhone" binding:"required,max=20"`
	NoCookie   bool   `json:"noCookie" form:"noCookie"`
}

type UserDeleteReq struct {
	ID int `json:"id" binding:"required"`
}

type UserChangeRoleReq struct {
	ID   int  `json:"userId" binding:"required"`
	Type Role `json:"userType" binding:"required"`
}

type AuthReq struct {
	Type    string `json:"type" binding:"required,oneof=account email"`
	Account string `json:"account" `
	Email   string `json:"email" form:"email"`
	Secret  string `json:"secret" form:"secret" binding:"required"`
}

type AuthResq struct {
	UserName  string `json:"userName"`
	UserEmail string `json:"userEmail"`
	UserType  Role   `json:"userType"`
	UserToken string `json:"userToken"`
	LoginType string `json:"loginType"`
}

type TokenAuth struct {
	Token string `json:"token" form:"token" binding:"required"`
}
