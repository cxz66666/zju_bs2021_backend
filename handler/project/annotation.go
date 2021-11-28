package project

import (
	"annotation/define"
	"annotation/model/project"
	"annotation/service/project_service"
	"annotation/utils/response"
	"github.com/gin-gonic/gin"
)

func ChangeRegion(c *gin.Context)  {
	regionReq:=project.AnnotationRegionReq{}
	if err:=c.ShouldBind(&regionReq);err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	err:=project_service.ChangeRegion(regionReq.Id,regionReq.Regions)
	if err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	c.Set(define.ANNOTATIONRESPONSE,response.JSONData("success"))
	return
}

func ChangeAnnotationStatus(c *gin.Context)  {
	//TODO
}