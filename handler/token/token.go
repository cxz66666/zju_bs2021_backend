package token

import (
	"annotation/define"
	"annotation/model/user"
	"annotation/service/user_service"
	"annotation/utils/authUtils"
	"annotation/utils/crypto"
	"annotation/utils/logging"
	"annotation/utils/response"
	"annotation/utils/setting"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context){
	//从请求中获取账号密码
	var userAuth user.AuthReq
	if  err:=c.ShouldBind(&userAuth);err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONError(response.ERROR_PARAM_FAIL))
		c.Abort()
		return
	}
	loginStatus:=false
	defer c.Set(define.ANNOTATIONLOGINSTATUS,loginStatus)

	if userAuth.Email==setting.AdminSetting.Email{
		if userAuth.Secret==setting.AdminSetting.Password{
			loginStatus=true
			adminToken,err:=authUtils.GetSysAdminToken()
			if err!=nil{
				c.Set(define.ANNOTATIONRESPONSE,response.JSONError(response.ERROR_TOKEN_GENERATE_FAIL))
				c.Abort()
				return
			}
			c.SetCookie(define.ANNOTATIONTOKEN,"Bearer "+adminToken,int(setting.ServerSetting.JwtExpireTime.Seconds()),"/","",false,true)
			c.Set(define.ANNOTATIONRESPONSE,response.JSONData(user_service.NewLoginResp(user.User{
				UserEmail: setting.AdminSetting.Email,
				UserName:  setting.AdminSetting.Name,
				UserType:  user.SysAdmin,
			},adminToken)))
			return
		} else {
			c.Set(define.ANNOTATIONRESPONSE,response.JSONError(response.ERROR_NOT_ADMIN))
			c.Abort()
			return
		}
	}

	secret:=crypto.Password2Secret(userAuth.Secret)

	//根据 account（stuid）查找用户
	queryUser := user_service.QueryUserByEmail(userAuth.Email)

	//账号不存在
	if queryUser.UserId ==0 {
		c.Set(define.ANNOTATIONRESPONSE,response.JSONError(response.ERROR_USERID))
		return
	}
	//判断账号密码
	if  secret== queryUser.UserSecret {
		loginStatus=true
		jwt, err := authUtils.GetStudentToken(queryUser)
		if err != nil{
			logging.ErrorF("generate token error for user:%+v\n",queryUser)
			c.Set(define.ANNOTATIONRESPONSE,response.JSONError(response.ERROR_TOKEN_GENERATE_FAIL))
			c.Abort()
		} else {
			c.SetCookie(define.ANNOTATIONTOKEN,"Bearer "+jwt,int(setting.ServerSetting.JwtExpireTime.Seconds()),"/","",false,true)
			c.Set(define.ANNOTATIONRESPONSE,response.JSONData(user_service.NewLoginResp(queryUser,jwt)))
		}
	}else{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONError(response.ERROR_PASSWORD))
		c.Abort()
	}
}

func Logout(c *gin.Context) {
	c.SetCookie(define.ANNOTATIONTOKEN,"",-1,"/","",false,true)
	c.Set(define.ANNOTATIONRESPONSE,response.JSONData("login out!"))
}

func Refresh(c *gin.Context)  {
	tmp,_:=c.Get(define.ANNOTATIONPOLICY)
	policy:=tmp.(authUtils.Policy)
	var token string
	if policy.SysAdminOnly() {
		token,_=authUtils.GetSysAdminToken()
	} else {
		token,_=authUtils.GetStudentToken(policy.ConvertToUser())
	}
	c.SetCookie(define.ANNOTATIONTOKEN,"Bearer "+token,int(setting.ServerSetting.JwtExpireTime.Seconds()),"/","",false,true)
	c.Set(define.ANNOTATIONRESPONSE,response.JSONData(user_service.NewLoginResp(policy.ConvertToUser(),token)))
	return
}