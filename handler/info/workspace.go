package info

import (
	"annotation/define"
	"annotation/service/project_service"
	"annotation/service/user_service"
	"annotation/utils/authUtils"
	"annotation/utils/response"
	"github.com/gin-gonic/gin"
)

func GetWorkspace(c *gin.Context) {
	claim, _ := c.Get(define.ANNOTATIONPOLICY)
	userId := claim.(authUtils.Policy).GetId()

	participateProjectNum, err := project_service.QueryProjectByWorker(userId)
	if err != nil {
		c.Set(define.ANNOTATIONRESPONSE, response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}

	projects, projectsNum, err := project_service.ListProject(6, 0)
	if err != nil {
		c.Set(define.ANNOTATIONRESPONSE, response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	userStaffNum, userAdminNum, err := user_service.GetNums()
	if err != nil {
		c.Set(define.ANNOTATIONRESPONSE, response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	c.Set(define.ANNOTATIONRESPONSE, response.JSONData(gin.H{
		"participateNum":   participateProjectNum,
		"proejcts":         projects,
		"totalProjectsNum": projectsNum,
		"totalUser":        userStaffNum + userAdminNum,
	}))
	return
}
