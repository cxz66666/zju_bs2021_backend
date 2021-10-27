package authUtils

import (
	"annotation/model/user"
	"annotation/utils/setting"
	"github.com/dgrijalva/jwt-go"
)

func GetStudentToken(user user.User) (string,error) {
	model:=GetClaimFromUser(user)
	tokenClaims:=jwt.NewWithClaims(jwt.SigningMethodHS256,model)

	token,err:=tokenClaims.SignedString([]byte(setting.SecretSetting.JwtKey))
	return token,err
}

func GetSysAdminToken() (string,error) {
	model:=GetClaimFromSysAdmin()
	tokenClaims:=jwt.NewWithClaims(jwt.SigningMethodHS256,model)

	token,err:=tokenClaims.SignedString([]byte(setting.SecretSetting.JwtKey))
	return token,err
}