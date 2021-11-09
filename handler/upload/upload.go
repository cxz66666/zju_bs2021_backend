package upload

import (
	"annotation/define"
	"annotation/service/upload_service"
	"annotation/utils/authUtils"
	"annotation/utils/logging"
	"annotation/utils/numberu"
	"annotation/utils/response"
	"github.com/gin-gonic/gin"
)


//用于上传图片
func UploadImage(c *gin.Context)  {

	claim,_:=c.Get(define.ANNOTATIONPOLICY)
	userId:=claim.(authUtils.Policy).GetId()

	form,err:=c.MultipartForm();
	if err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}

	files:=form.File["images"]
	projectIdSlice:=form.Value["id"]
	if len(projectIdSlice)==0||len(projectIdSlice)>1{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONError(response.ERROR_UPLOAD_NOT_ID))
		c.Abort()
		return
	}
	projectId:=numberu.ToInt(projectIdSlice[0])

	errs:=make([]string,0,0)
	count:=0
	for _,file:=range files{
		err=upload_service.SaveUploadedImage(projectId,userId,file)
		if err!=nil{
			logging.Info(err)
			errs = append(errs, err.Error())
		} else {
			count++
		}
	}
	c.Set(define.ANNOTATIONRESPONSE,response.JSONData(gin.H{
		"count":count,
		"errs":errs,
	}))
}


// UploadVideo 用于上传视频
func UploadVideo(c *gin.Context)  {

	claim,_:=c.Get(define.ANNOTATIONPOLICY)
	userId:=claim.(authUtils.Policy).GetId()
	form,err:=c.MultipartForm();
	if err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}

	videos:=form.File["videos"]
	projectIdSlice:=form.Value["id"]
	if len(projectIdSlice)==0||len(projectIdSlice)>1{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONError(response.ERROR_UPLOAD_NOT_ID))
		c.Abort()
		return
	}
	projectId:=numberu.ToInt(projectIdSlice[0])

	if len(videos)>1||len(videos)==0{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg("只能上传一个视频文件"))
		c.Abort()
		return
	}
	err=upload_service.SaveUploadedVideo(projectId,userId,videos[0])

	c.Set(define.ANNOTATIONRESPONSE,response.JSONData(gin.H{
		"count":1,
		"errs":err,
	}))

}