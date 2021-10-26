package authUtils

import (
	"annotation/utils/setting"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

func ParseToken(token string) (Policy,error) {
	// 解析token
	tokenClaims, err := jwt.ParseWithClaims(token, &Payload{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(setting.SecretSetting.JwtKey),nil
	})

	if err != nil {
		return nil, err
	}
	if payload, ok := tokenClaims.Claims.(*Payload); ok && tokenClaims.Valid { // 校验token
		return payload, nil
	}
	return nil, errors.New("invalid token")
}