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

}
