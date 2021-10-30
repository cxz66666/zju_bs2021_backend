package user

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
	"strings"
)

func GetInfo(c *gin.Context) {
	claim,_:=c.Get(define.ANNOTATIONPOLICY)
	policy,_:=claim.(authUtils.Policy)
	//如果是系统管理员 直接特判返回
	if policy.SysAdminOnly() {
		sysAdminResp:=user.UserInfoResp{
			ID: setting.AdminSetting.UserId,
			Name: setting.AdminSetting.Name,
			Email: setting.AdminSetting.Email,
			Type: user.SysAdmin,
			Phone: "1008611",
		}
		c.Set(define.ANNOTATIONRESPONSE,response.JSONData(sysAdminResp))
		return
	}
	userID:=policy.GetId()

	userRec:=user_service.QueryUserById(userID)

	if userRec.UserId <0{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONError(response.ERROR_USER_NOT_FOUND))
		c.Abort()
		return
	}
	userResp := user.UserInfoResp{
		userRec.UserId,
		userRec.UserName,
		userRec.UserEmail,
		userRec.UserType,
		userRec.UserPhone,
	}
	c.Set(define.ANNOTATIONRESPONSE,response.JSONData(userResp))
	return
}

func ModifyInfo	(c *gin.Context)  {
	claim,_:=c.Get(define.ANNOTATIONPOLICY)
	policy,_:=claim.(authUtils.Policy)


	//特判sysAdmin情况
	if policy.SysAdminOnly() {
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg("系统管理员不支持修改资料"))
		c.Abort()
		return
	}


	userID:=policy.GetId()

	userRec:=user_service.QueryUserById(userID)
	//后续清除缓存用到oldUser
	oldUser:=userRec

	if userRec.UserId <0{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONError(response.ERROR_USER_NOT_FOUND))
		c.Abort()
		return
	}

	req:=user.UserModifyReq{}
	if err:=c.ShouldBind(&req);err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONError(response.ERROR_PARAM_FAIL))
		c.Abort()
		return
	}

	if strings.Contains(req.ModifyField, "name") {
		userRec.UserName = req.UserName
	}
	if strings.Contains(req.ModifyField, "email") {
		userRec.UserEmail = req.UserEmail
	}
	if strings.Contains(req.ModifyField, "type") {
		userRec.UserType = req.UserType
	}
	if strings.Contains(req.ModifyField, "phone") {
		userRec.UserPhone = req.UserPhone
	}
	err:=user_service.UpdateUser(userRec)

	if err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONError(response.ERROR_UPDATE_FAIL))
		c.Abort()
		return
	}

	user_service.CleanUserCache(oldUser)

	c.Set(define.ANNOTATIONRESPONSE,response.JSONData(gin.H{
		"description":"success",
	}))
	return
}

func CreateUser(c *gin.Context)  {
	var userCreate user.UserCreateReq
	if  err:=c.ShouldBind(&userCreate);err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONError(response.ERROR_PARAM_FAIL))
		c.Abort()
		return
	}

	if !user_service.ValidUser(userCreate){
		c.Set(define.ANNOTATIONRESPONSE,response.JSONError(response.ERROR_NOT_VALID_USER_PARAM))
		c.Abort()
		return
	}

	userCreate.UserSecret=crypto.Password2Secret(userCreate.UserSecret)

	if err:=user_service.CreateUser(&userCreate);err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}

	logging.InfoF("create a new user: %v\n",userCreate)

	c.Set(define.ANNOTATIONRESPONSE,response.JSONData("success"))
	return
}