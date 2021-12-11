package project

import (
	"annotation/define"
	"annotation/model/project"
	"annotation/service/project_service"
	"annotation/utils/response"
	"github.com/gin-gonic/gin"
)

//ChangeRegion 修改批量的标注的信息
func ChangeRegion(c *gin.Context) {
	crReq := project.AnnotationRegionReq{}
	if err := c.ShouldBind(&crReq); err != nil {
		c.Set(define.ANNOTATIONRESPONSE, response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	err := project_service.ChangeRegion(crReq)
	if err != nil {
		c.Set(define.ANNOTATIONRESPONSE, response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	c.Set(define.ANNOTATIONRESPONSE, response.JSONData("success"))
	return
}

//ChangeAnnotationType 修改批量的标注的状态
func ChangeAnnotationType(c *gin.Context) {
	ctReq := project.AnnotationTypeReq{}
	if err := c.ShouldBind(&ctReq); err != nil {
		c.Set(define.ANNOTATIONRESPONSE, response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	err := project_service.ChangeAnnotationType(ctReq)
	if err != nil {
		c.Set(define.ANNOTATIONRESPONSE, response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	c.Set(define.ANNOTATIONRESPONSE, response.JSONData("success"))
	return
}
