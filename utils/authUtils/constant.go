package authUtils

import (
	"annotation/model/user"
	"annotation/utils/setting"
	"github.com/dgrijalva/jwt-go"
	"time"
)


 type Policy interface {
	 AdminOnly() bool
	 CheckExpired() bool
	 SysAdminOnly() bool
	 StaffOnly() bool
	 GetId() int
	 GetName() string
	 GetEmail()string
	 ConvertToUser() user.User
 }



 type Payload struct {
	 Name string `json:"name"`
	 UserId int `json:"user_id"`
	 Email string `json:"email"`
	 Role user.Role `json:"role"`
	 jwt.StandardClaims
 }

func (p *Payload) AdminOnly() bool {
	return p.Role==user.Admin || p.Role==user.SysAdmin
}

func (p *Payload) CheckExpired() bool {
	return time.Now().Unix()<p.ExpiresAt
}


func (p *Payload) SysAdminOnly() bool {
	return p.Role==user.SysAdmin
}

func (p *Payload) StaffOnly() bool {
	return p.Role==user.SysAdmin || p.Role==user.Admin || p.Role==user.Staff
}

func (p *Payload) GetId() int {
	return p.UserId
}

func (p *Payload) GetName() string {
	return p.Name
}

func (p *Payload) GetEmail() string {
	return p.Email
}

func (p *Payload) ConvertToUser() user.User {
	return user.User{
		UserEmail: p.Email,
		UserType:  p.Role,
		UserName:  p.Name,
		UserId:    p.UserId,
	}
}



//GetClaimFromUser convert token.LoginReq to Payload (already init the jwt.StandardClaims)
func GetClaimFromUser(user user.User) *Payload {
	nowTime:=time.Now()
	expireTime:=nowTime.Add(setting.ServerSetting.JwtExpireTime)

	return &Payload{
		Name: user.UserName,
		UserId: user.UserId,
		Email: user.UserEmail,
		Role: user.UserType,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt:expireTime.Unix(),
			Issuer: setting.SecretSetting.JwtIssuer,
		},
	}
}



//GetClaimFromSysAdmin get a payload for admin
func GetClaimFromSysAdmin() *Payload  {
	nowTime:=time.Now()
	expireTime:=nowTime.Add(setting.ServerSetting.JwtExpireTime)

	return &Payload{
		Name: setting.AdminSetting.Name,
		UserId: setting.AdminSetting.UserId,
		Email: setting.AdminSetting.Email,
		Role: user.SysAdmin,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: setting.SecretSetting.JwtIssuer,
		},
	}
}

