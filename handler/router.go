package handler

import (
	"annotation/handler/ping"
	"annotation/handler/tag"
	"annotation/handler/token"
	"annotation/handler/upload"
	"annotation/handler/user"
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

	userMod:=api.Group("/user")
	{
		userMod.GET("/me",middlware.AuthenticationMiddleware(),middlware.StaffOnly(), user.GetInfo)
		userMod.POST("/register",user.CreateUser)
		userMod.PUT("/me",middlware.AuthenticationMiddleware(),middlware.StaffOnly(),user.ModifyInfo)
	}

	tokenMod:=api.Group("/token")
	{
		tokenMod.POST("/login", token.Login)
		tokenMod.GET("/logout", middlware.AuthenticationMiddleware(),middlware.StaffOnly(), token.Logout)
		tokenMod.POST("/refresh", middlware.AuthenticationMiddleware(),middlware.StaffOnly(), token.Refresh)

	}

	tagMod:=api.Group("/class").Use(middlware.AuthenticationMiddleware(),middlware.StaffOnly())
	{
		tagMod.GET("/list",tag.GetClass)
		tagMod.POST("/create",tag.CreateClass)
		tagMod.PUT("/update",tag.UpdateClass)
		tagMod.DELETE("/delete",tag.DeleteClass)
		tagMod.POST("/tag/create",tag.CreateTag)
		tagMod.DELETE("/tag/delete",tag.DeleteTag)
	}

	uploadMod:=api.Group("/upload").Use(middlware.AuthenticationMiddleware(),middlware.StaffOnly())
	{
		uploadMod.POST("/image",upload.UploadImage)
	}

	settingMod:=api.Group("/setting").Use(middlware.AuthenticationMiddleware(),middlware.AdminOnly())
	{
		settingMod.GET("/setting",upload.GetSetting)
		settingMod.PUT("/setting",upload.UpdateSetting)
	}

	return r
}