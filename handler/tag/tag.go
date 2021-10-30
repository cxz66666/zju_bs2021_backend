package tag

import (
	"annotation/define"
	"annotation/model/tag"
	"annotation/service/tag_service"
	"annotation/service/user_service"
	"annotation/utils/authUtils"
	"annotation/utils/response"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func GetClass(c *gin.Context) {
	pageSizeStr:=c.Query("pageSize")
	currentStr:=c.Query("current")

	var pageSize,current int
	if len(pageSizeStr)==0 {
		pageSize=20
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
	classes,err:=tag_service.QueryClass(pageSize,current)
	if err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	var classInfoResp []tag.ClassInfoResp


	for _,m:=range classes{
		user:=user_service.QueryUserById(m.CreatorId)
		name:=user.UserName
		if len(name)==0{
			name="未找到"
		}

		classInfoResp = append(classInfoResp, tag.ClassInfoResp{
			Id:m.Id,
			ClassName: m.ClassName,
			Description: m.Description,
			CreatorName: name,
			Tags: m.Tags,
		})
	}

	c.Set(define.ANNOTATIONRESPONSE,response.JSONData(gin.H{
		"number":len(classInfoResp),
		"classes":classInfoResp,
	}))

	return

}

func CreateClass(c *gin.Context)  {

	createClass:=tag.ClassCreateReq{}

	if err:=c.ShouldBind(&createClass);err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}

	//获取当前操作者的名称
	claim,_:=c.Get(define.ANNOTATIONPOLICY)
	userId:=claim.(authUtils.Policy).GetId()

	userRec:=user_service.QueryUserById(userId)

	var tags []tag.Tag
	//将post过来的tag模型转为数据库使用的tag.Tag模型
	for _,m:=range createClass.Tags{
		tags = append(tags, tag.Tag{
			Content: m.Content,
		})
	}
	//TODO 检查这玩意能不能使用
	class:=tag.Class{
		ClassName: createClass.ClassName,
		Description: createClass.Description,
		Creator: userRec,
		CreateTime: time.Now(),
		Tags: tags,
	}

	if err:=tag_service.CreateClass(&class);err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	
	c.Set(define.ANNOTATIONRESPONSE,response.JSONData("success"))
	return
}

func UpdateClass(c *gin.Context)  {
	
}