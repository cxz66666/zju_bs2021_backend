package upload

import (
	"annotation/define"
	"annotation/model/upload"
	"annotation/service/upload_service"
	"annotation/utils/file"
	"annotation/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetImage(c *gin.Context)  {
	pidStr:=c.Param("pid")
	crc32Str:=c.Param("crc32")
	fileName:=c.Param("name")

	if fileName[0]=='/' {
		fileName=fileName[1:]
	}
	var pid,crc32 int
	if pidInt,err:=strconv.ParseInt(pidStr,10,64);err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.Failed(http.StatusNotFound))
		c.Abort()
		return
	} else {
		pid=int(pidInt)
	}

	if crc32Int,err:=strconv.ParseInt(crc32Str,10,64);err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.Failed(http.StatusNotFound))
		c.Abort()
		return
	} else {
		crc32=int(crc32Int)
	}

	if len(fileName)==0 {
		c.Set(define.ANNOTATIONRESPONSE,response.Failed(http.StatusNotFound))
		c.Abort()
		return
	}

	images,err:=upload_service.QueryImage(pid,crc32,fileName)
	if err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.Failed(http.StatusNotFound))
		c.Abort()
		return
	}

	if images.Type==upload.Backend {
		if !file.CheckExist(images.StorePath) {
			c.Set(define.ANNOTATIONRESPONSE,response.Failed(http.StatusNotFound))
			c.Abort()
			return
		} else {
			c.Set(define.ANNOTATIONRESPONSE,response.Image(images.StorePath))
			c.Abort()
			return
		}
	} else if images.Type==upload.OSS {
		c.Set(define.ANNOTATIONRESPONSE,response.Redirect(301,images.StorePath))
		c.Abort()
		return
	} else {
		c.Set(define.ANNOTATIONRESPONSE,response.Failed(http.StatusNotFound))
		c.Abort()
		return
	}
}

func ListImage(c *gin.Context)  {
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
	images,total,err:=upload_service.QueryListImages(pageSize,current)
	if err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	imageInfoResp:=make([]upload.ImageResp,0,len(images))

	for _,m:=range images{
		name:=m.Creator.UserName
		if len(name)==0{
			name="未知者"
		}
		imageInfoResp = append(imageInfoResp, upload.ImageResp{
			Id: m.Id,
			Name: m.Name,
			ProjectId: m.ProjectId,
			Type: m.Type,
			Url: m.GetUrl(),
			CreatorId: m.CreatorId,
			CreatorName: name,
			UploadTime: m.UploadTime.Format("2006-01-02"),
		})
	}

	c.Set(define.ANNOTATIONRESPONSE,response.JSONData(gin.H{
		"total":total,
		"number":len(imageInfoResp),
		"data":imageInfoResp,
	}))
	return
}
