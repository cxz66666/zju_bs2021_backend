package user_service

import "annotation/model/user"


// NewLoginResp return a authResp model with a use model and token
func NewLoginResp(us user.User,_token string) *user.AuthResq {

	return &user.AuthResq{
		UserEmail: us.UserEmail,
		UserName:  us.UserName,
		UserType:  us.UserType,
		UserToken: _token,
	}
}