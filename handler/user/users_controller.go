package user

import (
	"annotation/define"
	"annotation/model/user"
	"annotation/service/user_service"
	"annotation/utils/response"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	users,total,err:=user_service.QueryUsers()
	if err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}

	userInfoResp:=make([]user.UserInfoResp,0,len(users))

	for _,m:=range users{
		userInfoResp = append(userInfoResp, user.UserInfoResp{
			ID:m.UserId,
			Name: m.UserName,
			Email: m.UserEmail,
			Type: m.UserType,
			Phone: m.UserPhone,
		} )
	}

	c.Set(define.ANNOTATIONRESPONSE,response.JSONData(gin.H{
		"total":total,
		"data":userInfoResp,
	}))
	return
}

func DeleteUsers(c *gin.Context)  {
	deleteReq:= user.UserDeleteReq{}
	if err:=c.ShouldBind(&deleteReq);err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	err:=user_service.DeleteUser(deleteReq.ID)
	if err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	c.Set(define.ANNOTATIONRESPONSE,response.JSONData("success"))
	return
}

func ChangeUserRole(c *gin.Context)  {
	changeRoleReq:=user.UserChangeRoleReq{}
	if err:=c.ShouldBind(&changeRoleReq);err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	err:=user_service.ChangeRole(changeRoleReq.ID,changeRoleReq.Type)
	if err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	c.Set(define.ANNOTATIONRESPONSE,response.JSONData("success"))
	return
}

func GetNum(c *gin.Context)  {
	normal,admin,err:=user_service.GetNums();
	if err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	c.Set(define.ANNOTATIONRESPONSE,response.JSONData(gin.H{
		"normal":normal,
		"admin":admin,
		"total":normal+admin,
	}))
	return
}