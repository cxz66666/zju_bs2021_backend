package handler

import (
	"annotation/handler/ping"
	"annotation/middlware"
	"annotation/utils/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r:=gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//注意cors要在路由前设置
	r.Use(middlware.CorsMiddleware())

	gin.SetMode(setting.ServerSetting.RunMode)

	api:=r.Group("/api",
		middlware.RecoverMiddleware(),
		middlware.ResponseMiddleware(),
		middlware.RewriteToken())


	api.GET("/ping",ping.Pong)




	return r
}