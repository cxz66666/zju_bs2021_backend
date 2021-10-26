package middlware

import (
	"annotation/define"
	"annotation/utils/authUtils"
	"annotation/utils/response"
	"github.com/gin-gonic/gin"
)

//StaffOnly check the policy and return ERROR_TOKEN_NOT_VAILD if forbidden
//you need use it after jwt middleware!! very important
func StaffOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		claim,_:=c.Get(define.ANNOTATIONPOLICY)
		if policy,ok:= claim.(authUtils.Policy);!ok||!policy.StaffOnly(){
			c.Set(define.ANNOTATIONRESPONSE,response.JSONError(response.ERROR_TOKEN_NOT_VAILD))
			c.Abort()
			return
		}
	}
}


//AdminOnly check the policy and return ERROR_NOT_ADMIN if forbidden
//you need use it after jwt middleware!! very important
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		claim,_:=c.Get(define.ANNOTATIONPOLICY)
		if policy,ok:= claim.(authUtils.Policy);!ok||!policy.AdminOnly(){
			c.Set(define.ANNOTATIONRESPONSE,response.JSONError(response.ERROR_NOT_ADMIN))
			c.Abort()
			return
		}
	}
}


//SysAdminOnly check the policy and return ERROR_NOT_ADMIN if forbidden
//you need use it after jwt middleware!! very important
func SysAdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		claim,_:=c.Get(define.ANNOTATIONPOLICY)
		if policy,ok:= claim.(authUtils.Policy);!ok||!policy.SysAdminOnly(){
			c.Set(define.ANNOTATIONRESPONSE,response.JSONError(response.ERROR_NOT_ADMIN))
			c.Abort()
			return
		}
	}
}