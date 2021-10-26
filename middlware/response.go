package middlware

import (
	"annotation/define"
	"annotation/utils/logging"
	"annotation/utils/response"
	"github.com/gin-gonic/gin"
)

func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		value,exist:=c.Get(define.ANNOTATIONRESPONSE)
		if !exist{
			logging.Warn("response not set")
			return
		}
		resp,ok:=value.(response.Response)
		if !ok{
			logging.Warn("response type invalid!")
			return
		}
		resp.Write(c)	}
}