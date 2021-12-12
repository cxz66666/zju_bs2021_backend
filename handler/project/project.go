package project

import (
	"annotation/define"
	"annotation/model/project"
	"annotation/model/user"
	"annotation/service/project_service"
	"annotation/utils/authUtils"
	"annotation/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func CreateProject(c *gin.Context) {
	createReq := project.ProjectCreateReq{}
	if err := c.ShouldBind(&createReq); err != nil {
		c.Set(define.ANNOTATIONRESPONSE, response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	claim, _ := c.Get(define.ANNOTATIONPOLICY)
	userId := claim.(authUtils.Policy).GetId()
	workers := make([]user.User, 0, len(createReq.UserList))
	for _, m := range createReq.UserList {
		workers = append(workers, user.User{
			UserId: m,
		})
	}
	nowProject := project.Project{
		Name:        createReq.Name,
		Description: createReq.Description,
		CreatedTime: time.Now(),
		CreatorId:   userId,
		Workers:     workers,
		Type:        project.Pcreated,
		ClassId:     createReq.ClassId,
	}
	if err := project_service.CreateProject(&nowProject); err != nil {
		c.Set(define.ANNOTATIONRESPONSE, response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}

	c.Set(define.ANNOTATIONRESPONSE, response.JSONData(project.ProjectCreateResp{
		Id: nowProject.Id,
	}))
	return
}

func ListProject(c *gin.Context) {
	claim, _ := c.Get(define.ANNOTATIONPOLICY)
	policy, _ := claim.(authUtils.Policy)
	pageSizeStr := c.Query("pageSize")
	currentStr := c.Query("current")

	var pageSize, current int
	if len(pageSizeStr) == 0 {
		pageSize = 10
	} else {
		if pageSizeInt, err := strconv.ParseInt(pageSizeStr, 10, 64); err != nil {
			c.Set(define.ANNOTATIONRESPONSE, response.JSONErrorWithMsg("pageSize解析错误"))
			c.Abort()
			return
		} else {
			pageSize = int(pageSizeInt)
		}
	}

	if len(currentStr) == 0 {
		current = 1
	} else {
		if currentInt, err := strconv.ParseInt(currentStr, 10, 64); err != nil {
			c.Set(define.ANNOTATIONRESPONSE, response.JSONErrorWithMsg("current解析错误"))
			c.Abort()
			return
		} else {
			current = int(currentInt)
		}
	}

	projects, total, err := project_service.ListProject(pageSize, current)
	if err != nil {
		c.Set(define.ANNOTATIONRESPONSE, response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	//该用户加入的project数量
	workedProjectNum, err := project_service.QueryProjectByWorker(policy.GetId())
	if err != nil {
		c.Set(define.ANNOTATIONRESPONSE, response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	projectsInfoResp := make([]project.ProjectListResp, 0, len(projects))
	for _, m := range projects {
		if len(m.Class.ClassName) == 0 {
			m.Class.ClassName = "未确认class"
		}
		projectsInfoResp = append(projectsInfoResp, project.ProjectListResp{
			Id:            m.Id,
			Name:          m.Name,
			Description:   m.Description,
			ClassName:     m.Class.ClassName,
			ImagesNum:     len(m.Images),
			AnnotationNum: len(m.Annotations),
			WorkerNum:     len(m.Workers),
			Type:          m.Type,
			CreatedTime:   m.CreatedTime,
		})
	}

	c.Set(define.ANNOTATIONRESPONSE, response.JSONData(gin.H{
		"total":  total,
		"joined": workedProjectNum,
		"number": len(projectsInfoResp),
		"data":   projectsInfoResp,
	}))

	return
}

func GetProject(c *gin.Context) {
	idStr := c.Param("id")
	var id int
	if idInt, err := strconv.ParseInt(idStr, 10, 64); err != nil {
		c.Set(define.ANNOTATIONRESPONSE, response.Failed(http.StatusNotFound))
		c.Abort()
		return
	} else {
		id = int(idInt)
	}
	p, err := project_service.QueryProjectById(id)
	p.AnnotationMap = make(map[int]project.Annotation)

	if err != nil {
		c.Set(define.ANNOTATIONRESPONSE, response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	for k, m := range p.Annotations {
		m.Src = m.Image.GetUrl()
		p.AnnotationMap[m.ImageId] = m
		p.Annotations[k] = m
	}
	c.Set(define.ANNOTATIONRESPONSE, response.JSONData(p))

	return
}

//ChangeStatus 更新项目当前状态
func ChangeStatus(c *gin.Context) {
	idStr := c.Param("id")
	var id int
	if idInt, err := strconv.ParseInt(idStr, 10, 64); err != nil {
		c.Set(define.ANNOTATIONRESPONSE, response.Failed(http.StatusNotFound))
		c.Abort()
		return
	} else {
		id = int(idInt)
	}
	cs := project.ProjectChangeStatusReq{}
	if err := c.ShouldBind(&cs); err != nil {
		c.Set(define.ANNOTATIONRESPONSE, response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	err := project_service.ChangeStatus(id, cs.Type)
	if err != nil {
		c.Set(define.ANNOTATIONRESPONSE, response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	c.Set(define.ANNOTATIONRESPONSE, response.JSONData("success"))
	return

}

// 从公共图片加到项目中
func AddPublicImage(c *gin.Context) {
	publicReq := project.ProjectAddPublicReq{}
	if err := c.ShouldBind(&publicReq); err != nil {
		c.Set(define.ANNOTATIONRESPONSE, response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	err := project_service.AddImageAssociation(publicReq.ProjectId, publicReq.ImageId)
	if err != nil {
		c.Set(define.ANNOTATIONRESPONSE, response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	c.Set(define.ANNOTATIONRESPONSE, response.JSONData("success"))
	return
}

// 新获取若干任务
func GetAnnotationWorks(c *gin.Context) {
	tmp, _ := c.Get(define.ANNOTATIONPOLICY)
	policy := tmp.(authUtils.Policy)
	numStr := c.Query("num")

	var num int
	if len(numStr) == 0 {
		c.Set(define.ANNOTATIONRESPONSE, response.JSONErrorWithMsg("num不能为空"))
		c.Abort()
		return
	} else {
		if pageSizeInt, err := strconv.ParseInt(numStr, 10, 64); err != nil {
			c.Set(define.ANNOTATIONRESPONSE, response.JSONErrorWithMsg("num解析错误"))
			c.Abort()
			return
		} else {
			num = int(pageSizeInt)
		}
	}
	idStr := c.Param("id")
	var id int
	if idInt, err := strconv.ParseInt(idStr, 10, 64); err != nil {
		c.Set(define.ANNOTATIONRESPONSE, response.Failed(http.StatusNotFound))
		c.Abort()
		return
	} else {
		id = int(idInt)
	}
	p, err := project_service.QueryProjectById(id)
	p.AnnotationMap = make(map[int]project.Annotation)
	if err != nil {
		c.Set(define.ANNOTATIONRESPONSE, response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	for _, m := range p.Annotations {
		p.AnnotationMap[m.ImageId] = m
	}
	var newAnnotations []project.Annotation
	for _, m := range p.Images {
		if _, ok := p.AnnotationMap[m.Id]; !ok {
			newAnnotations = append(newAnnotations, project.Annotation{
				ProjectId:    id,
				WorkerId:     policy.GetId(),
				ImageId:      m.Id,
				Regions:      "",
				PixelSize:    "",
				Type:         project.Acreated,
				LastEditTime: time.Now(),
				Src:          m.GetUrl(),
			})
			if len(newAnnotations) == num {
				break
			}
		}
	}
	err = project_service.AddNewAnnotation(id, newAnnotations)
	if err != nil {
		c.Set(define.ANNOTATIONRESPONSE, response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}

	c.Set(define.ANNOTATIONRESPONSE, response.JSONData(gin.H{
		"number": len(newAnnotations),
		"data":   newAnnotations,
	}))
	return
}

func DeleteProject(c *gin.Context) {
	idStr := c.Param("id")
	var id int
	if idInt, err := strconv.ParseInt(idStr, 10, 64); err != nil {
		c.Set(define.ANNOTATIONRESPONSE, response.Failed(http.StatusNotFound))
		c.Abort()
		return
	} else {
		id = int(idInt)
	}
	err := project_service.DeleteProjectByPid(id)
	if err != nil {
		c.Set(define.ANNOTATIONRESPONSE, response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	c.Set(define.ANNOTATIONRESPONSE, response.JSONData("success"))
	return
}
