package middlware

import (
	"annotation/define"
	"annotation/utils/authUtils"
	"annotation/utils/response"
	"github.com/gin-gonic/gin"
	"strings"
)

// AuthenticationMiddleware 身份验证
func  AuthenticationMiddleware()  gin.HandlerFunc {
	return func(c *gin.Context) {

		// 注意 之前已经使用rewrite中间件，将cookie or header中字段的token统一放倒了`Authorization`字段中，所以这里才可以这样使用
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.Set(define.ANNOTATIONRESPONSE,response.JSONError(response.ERROR_NOT_LOGIN))
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.Set(define.ANNOTATIONRESPONSE,response.JSONError(response.ERROR_AUTH_NO_VALID_HEADER))
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		policy, err := authUtils.ParseToken(parts[1])
		//解析失败
		if err != nil {
			c.Set(define.ANNOTATIONRESPONSE,response.JSONError(response.ERROR_TOKEN_NOT_VAILD))
			c.Abort()
			return
		} else if !policy.CheckExpired(){
			//过期了
			c.Set(define.ANNOTATIONRESPONSE,response.JSONError(response.ERROR_TOKEN_EXPIRED))
			c.Abort()
			return
		}
		// 获取到一个可以拿到policy的接口，存起来
		c.Set(define.ANNOTATIONPOLICY, policy)
		c.Next() // 后续的处理函数可以用过c.Get(define.NOTIFYPOLICY)来获取当前请求上下文
	}

}