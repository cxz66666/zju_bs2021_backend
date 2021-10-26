package authUtils

import (
	"annotation/utils/setting"
	"github.com/dgrijalva/jwt-go"
	"annotation/model/user"
	"time"
)

type Role int


const (
	Staff Role =1
	Admin Role =2
	SysAdmin Role=3
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
	 Role Role `json:"role"`
	 jwt.StandardClaims
 }

func (p *Payload) AdminOnly() bool {
	return p.Role==Admin || p.Role==SysAdmin
}

func (p *Payload) CheckExpired() bool {
	return time.Now().Unix()<p.ExpiresAt
}


func (p *Payload) SysAdminOnly() bool {
	return p.Role==SysAdmin
}

func (p *Payload) StaffOnly() bool {
	return p.Role==SysAdmin || p.Role==Admin || p.Role==Staff
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
		User_email: p.Email,
		User_type: int(p.Role),
		User_name: p.Name,
		User_id: p.UserId,
	}
}



//GetClaimFromUser convert token.LoginReq to Payload (already init the jwt.StandardClaims)
func GetClaimFromUser(user user.User) *Payload {
	nowTime:=time.Now()
	expireTime:=nowTime.Add(setting.ServerSetting.JwtExpireTime)

	return &Payload{
		Name: user.User_name,
		UserId: user.User_id,
		Email: user.User_email,
		Role: Role(user.User_type),
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
		Name: "系统管理员",
		UserId: 10086,
		Email: "cxz@zjueva.net",
		Role: SysAdmin,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: setting.SecretSetting.JwtIssuer,
		},
	}
}

