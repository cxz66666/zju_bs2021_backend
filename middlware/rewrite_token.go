package middlware

import (
	"annotation/define"
	"github.com/gin-gonic/gin"
	"strings"
)

//RewriteToken 中间件将放在cookie或者header中字段的token 放到`Authorization`中，方便后续的jwt验证和授权
func RewriteToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader:=c.Request.Header.Get("Authorization")
		if !(len(authHeader)>0&&strings.HasPrefix(authHeader,"Bearer ")){
			//先从cookie里找
			cookie,err:=c.Cookie(define.ANNOTATIONTOKEN)
			if err==nil{
				//dont't forget to "add Bearer "
				if strings.Contains(cookie,"Bearer ") {
					c.Request.Header.Set("Authorization",cookie)
				} else {
					c.Request.Header.Set("Authorization", "Bearer "+cookie)
				}
			} else {
				//再从header的部分找
				token:=c.Request.Header.Get(define.ANNOTATIONTOKEN)
				if len(token)>0{
					//dont't forget to "add Bearer "
					if strings.Contains(token,"Bearer ") {
						c.Request.Header.Set("Authorization",token)
					} else {
						c.Request.Header.Set("Authorization", "Bearer "+token)
					}
				}
			}

		}
	}
}
