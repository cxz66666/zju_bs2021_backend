package middlware

import (
	"annotation/define"
	"annotation/service/ipblocker"
	"annotation/utils/logging"
	"annotation/utils/response"
	"github.com/gin-gonic/gin"
)

func IPBlock() gin.HandlerFunc  {
	return func(c *gin.Context) {
		ip:=c.ClientIP()
		if result:=ipblocker.IsLoginable(ip);!result{
			c.Set(define.ANNOTATIONRESPONSE,response.JSONError(response.ERROR_IP_BLOCK))
			c.Abort()
			return
		}
		c.Next()
		status,exist:=c.Get(define.ANNOTATIONLOGINSTATUS)
		if statusBool,ok:=status.(bool);!exist||!ok||!statusBool{
			logging.InfoF("ip %s try to login and fail\n",ip)
			ipblocker.Fail(ip)
		} else {
			ipblocker.Success(ip)
		}

	}
}