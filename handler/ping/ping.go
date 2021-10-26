package ping

import (
	"annotation/define"
	"annotation/utils/response"
	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context)  {
	c.Set(define.ANNOTATIONRESPONSE,response.JSONData(gin.H{
		"message":"pong!",
	}))
}