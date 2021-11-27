package project

import (
	"annotation/define"
	"annotation/model/project"
	"annotation/model/user"
	"annotation/service/project_service"
	"annotation/utils/authUtils"
	"annotation/utils/response"
	"github.com/gin-gonic/gin"
	"strconv"
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
	workers:=make([]user.User,0,len(createReq.UserList))
	for _,m:=range createReq.UserList{
		workers = append(workers, user.User{
			UserId:m,
		})
	}
	nowProject:=project.Project{
		Name: createReq.Name,
		Description: createReq.Description,
		CreatedTime: time.Now(),
		CreatorId: userId,
		Workers: workers,
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

func ListProject(c *gin.Context)  {
	pageSizeStr:=c.Query("pageSize")
	currentStr:=c.Query("current")

	var pageSize,current int
	if len(pageSizeStr)==0 {
		pageSize=10
	} else {
		if pageSizeInt,err:=strconv.ParseInt(pageSizeStr,10,64);err!=nil{
			c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg("pageSize解析错误"))
			c.Abort()
			return
		} else {
			pageSize=int(pageSizeInt)
		}
	}

	if len(currentStr)==0{
		current=1
	} else {
		if currentInt,err:=strconv.ParseInt(currentStr,10,64);err!=nil{
			c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg("current解析错误"))
			c.Abort()
			return
		} else {
			current=int(currentInt)
		}
	}

	projects,total,err:=project_service.ListProject(pageSize,current)
	if err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	
	projectsInfoResp:=make([]project.ProjectListResp,0,len(projects))
	for _,m:=range projects{
		if len(m.Class.ClassName)==0{
			m.Class.ClassName="未确认class"
		}

		projectsInfoResp = append(projectsInfoResp, project.ProjectListResp{
			Id: m.Id,
			Name: m.Name,
			Description: m.Description,
			ClassName: m.Class.ClassName,
			ImagesNum: len(m.Images),
			Type: m.Type,
			CreatedTime: m.CreatedTime,
		})
	}

	c.Set(define.ANNOTATIONRESPONSE,response.JSONData(gin.H{
		"total":total,
		"number":len(projectsInfoResp),
		"data":projectsInfoResp,
	}))

	return
}

func GetProject(c *gin.Context)  {
	idStr:=c.Param("id")
	var id int
	if idInt,err:=strconv.ParseInt(idStr,10,64);err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg("pageSize解析错误"))
		c.Abort()
		return
	} else {
		id=int(idInt)
	}
	p,err:=project_service.QueryProjectById(id)
	if err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}

	c.Set(define.ANNOTATIONRESPONSE,response.JSONData(p))

	return
}

//ChangeStatus 更新项目当前状态
func ChangeStatus(c *gin.Context)  {
	idStr:=c.Param("id")
	var id int
	if idInt,err:=strconv.ParseInt(idStr,10,64);err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg("pageSize解析错误"))
		c.Abort()
		return
	} else {
		id=int(idInt)
	}
	cs:=project.ProjectChangeStatusReq{}
	if err:=c.ShouldBind(&cs);err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	err:=

}
