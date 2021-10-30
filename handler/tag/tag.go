package tag

import (
	"annotation/define"
	"annotation/model/tag"
	"annotation/service/tag_service"
	"annotation/service/user_service"
	"annotation/utils/authUtils"
	"annotation/utils/logging"
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
	classInfoResp:=make([]tag.ClassInfoResp,0,len(classes))

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
			CreateTime: m.CreateTime,
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
			Content:m,
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
	updateClass:=tag.ClassUpdateReq{}
	if err:=c.ShouldBind(&updateClass);err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}

	oldClass,err:=tag_service.QueryClassById(updateClass.Id)
	if oldClass.Id<=0 {
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg("未找到该id"))
		c.Abort()
		return
	}
	if err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}

	if len(updateClass.ClassName)>0{
		oldClass.ClassName=updateClass.ClassName
	}

	if len(updateClass.Description)>0{
		oldClass.Description=updateClass.Description
	}


	err=tag_service.UpdateClass(oldClass)
	if err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}

	c.Set(define.ANNOTATIONRESPONSE,response.JSONData("success"))
	return

}

func DeleteClass(c *gin.Context)  {
	deleteClass:=tag.ClassDeleteReq{}
	if err:=c.ShouldBind(&deleteClass);err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}

	err:=tag_service.DeleteClass(deleteClass.Id)
	if err!=nil{
		logging.Error("删除class时出现问题")

		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	c.Set(define.ANNOTATIONRESPONSE,response.JSONData("success"))
	return
}

func CreateTag(c *gin.Context)  {
	createTag:=tag.TagCreateReq{}
	if err:=c.ShouldBind(&createTag);err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}

	err:=tag_service.CreateTag(&tag.Tag{
		ClassId: createTag.ClassId,
		Content: createTag.Content,
	})

	if err!=nil{
		logging.Error("创建tag时出现问题")
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}

	c.Set(define.ANNOTATIONRESPONSE,response.JSONData("success"))
	return
}

func DeleteTag(c *gin.Context)  {
	deleteTag:=tag.DeleteTagReq{}
	if err:=c.ShouldBind(&deleteTag);err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}

	err:=tag_service.DeleteTag(deleteTag.TagId)
	if err!=nil {
		logging.Error("删除tag时出现问题")
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}

	c.Set(define.ANNOTATIONRESPONSE,response.JSONData("success"))
	return
}