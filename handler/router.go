package handler

import (
	"annotation/handler/ping"
	"annotation/handler/project"
	"annotation/handler/tag"
	"annotation/handler/token"
	"annotation/handler/upload"
	"annotation/handler/user"
	"annotation/middlware"
	"annotation/utils/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//注意cors要在路由前设置
	r.Use(middlware.CorsMiddleware())

	gin.SetMode(setting.ServerSetting.RunMode)

	api := r.Group("/api",
		middlware.RecoverMiddleware(),
		middlware.ResponseMiddleware(),
		middlware.RewriteToken())

	api.GET("/ping", ping.Pong)

	userMod := api.Group("/user")
	{
		userMod.GET("/me", middlware.AuthenticationMiddleware(), middlware.StaffOnly(), user.GetInfo)
		userMod.POST("/register", user.CreateUser)
		userMod.PUT("/me", middlware.AuthenticationMiddleware(), middlware.StaffOnly(), user.ModifyInfo)
	}
	usersMod := api.Group("/users").Use(middlware.AuthenticationMiddleware(), middlware.SysAdminOnly())
	{
		usersMod.GET("/list", user.GetUsers)
		usersMod.DELETE("/user", user.DeleteUsers)
		usersMod.PUT("/role", user.ChangeUserRole)
		usersMod.GET("/num", user.GetNum)
		usersMod.GET("/alluser", user.GetAllUser)

	}
	tokenMod := api.Group("/token")
	{
		tokenMod.POST("/login", token.Login)
		tokenMod.GET("/logout", middlware.AuthenticationMiddleware(), middlware.StaffOnly(), token.Logout)
		tokenMod.POST("/refresh", middlware.AuthenticationMiddleware(), middlware.StaffOnly(), token.Refresh)

	}

	tagMod := api.Group("/class").Use(middlware.AuthenticationMiddleware(), middlware.StaffOnly())
	{
		tagMod.GET("/allclass", tag.GetAllClass)
		tagMod.GET("/list", tag.GetClass)
		tagMod.POST("/create", tag.CreateClass)
		tagMod.PUT("/update", tag.UpdateClass)
		tagMod.DELETE("/delete", tag.DeleteClass)
		tagMod.POST("/tag/create", tag.CreateTag)
		tagMod.DELETE("/tag/delete", tag.DeleteTag)
	}

	uploadMod := api.Group("/upload").Use(middlware.AuthenticationMiddleware(), middlware.StaffOnly())
	{
		uploadMod.POST("/image", upload.UploadImage)
		uploadMod.POST("/video", upload.UploadVideo)

	}

	settingMod := api.Group("/setting").Use(middlware.AuthenticationMiddleware(), middlware.AdminOnly())
	{
		settingMod.GET("/setting", upload.GetSetting)
		settingMod.PUT("/setting", upload.UpdateSetting)
	}

	projectMod := api.Group("/project").Use(middlware.AuthenticationMiddleware(), middlware.StaffOnly())
	{
		projectMod.POST("/new", middlware.AdminOnly(), project.CreateProject)
		projectMod.GET("/list", project.ListProject)
		projectMod.PUT("/cs/:id", middlware.AdminOnly(), project.ChangeStatus)
		projectMod.POST("/addpi", project.AddPublicImage)
		projectMod.GET("/cn/:id", project.GetAnnotationWorks)
		projectMod.POST("/cr", project.ChangeRegion)
		projectMod.POST("/ct", project.ChangeAnnotationType)
		projectMod.GET("/:id", project.GetProject)
		projectMod.DELETE("/:id", middlware.AdminOnly(), project.DeleteProject)

	}

	imageMod := api.Group("/image")
	{
		imageMod.GET("/list/:id", upload.ListImage)
		imageMod.GET("/:pid/:crc32/*name", upload.GetImage)
	}

	return r
}
