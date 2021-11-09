package project

import (
	"annotation/define"
	"annotation/model/project"
	"annotation/service/project_service"
	"annotation/utils/authUtils"
	"annotation/utils/response"
	"github.com/gin-gonic/gin"
	"time"
)

func CreateProject(c *gin.Context) {

	createReq:=project.ProjectCreateReq{}
	if err:=c.ShouldBind(&createReq);err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	claim,_:=c.Get(define.ANNOTATIONPOLICY)
	userId:=claim.(authUtils.Policy).GetId()

	nowProject:=project.Project{
		Name: createReq.Name,
		Description: createReq.Description,
		CreatedTime: time.Now(),
		CreatorId: userId,
		Type: project.Pcreated,
		ClassId: createReq.ClassId,
	}
	if err:=project_service.CreateProject(&nowProject);err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}

	c.Set(define.ANNOTATIONRESPONSE,response.JSONData(project.ProjectCreateResp{
		Id: nowProject.Id,
	}))
	return

}