package upload

import (
	"annotation/define"
	"annotation/service/upload_service"
	"annotation/utils/authUtils"
	"annotation/utils/logging"
	"annotation/utils/numberu"
	"annotation/utils/response"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"sync"
)

//用于上传图片
func UploadImage(c *gin.Context) {

	claim, _ := c.Get(define.ANNOTATIONPOLICY)
	userId := claim.(authUtils.Policy).GetId()

	form, err := c.MultipartForm()
	if err != nil {
		c.Set(define.ANNOTATIONRESPONSE, response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}

	files := form.File["images"]
	if len(files) == 0 {
		c.Set(define.ANNOTATIONRESPONSE, response.JSONError(response.ERROR_UPLOAD_NOT_ID))
		c.Abort()
		return
	}
	projectIdSlice := form.Value["id"]
	if len(projectIdSlice) > 1 {
		c.Set(define.ANNOTATIONRESPONSE, response.JSONError(response.ERROR_UPLOAD_NOT_ID))
		c.Abort()
		return
	}
	var projectId int
	if len(projectIdSlice) > 0 {
		projectId = numberu.ToInt(projectIdSlice[0])
	} else {
		projectId = 0
	}

	errs := make([]string, 0, 0)
	count := 0
	var wg sync.WaitGroup
	wg.Add(len(files))
	chans := make(chan error, len(files))
	for _, file := range files {
		go func(localfile *multipart.FileHeader) {
			chans <- upload_service.SaveUploadedImage(projectId, userId, localfile)
		}(file)
	}
	total := 0
	for err = range chans {
		total++
		if err != nil {
			logging.Info(err)
			errs = append(errs, err.Error())
		} else {
			count++
		}
		if total == len(files) {
			close(chans)
		}
	}
	c.Set(define.ANNOTATIONRESPONSE, response.JSONData(gin.H{
		"count": count,
		"errs":  errs,
	}))
}

// UploadVideo 用于上传视频
func UploadVideo(c *gin.Context) {

	claim, _ := c.Get(define.ANNOTATIONPOLICY)
	userId := claim.(authUtils.Policy).GetId()
	form, err := c.MultipartForm()
	if err != nil {
		c.Set(define.ANNOTATIONRESPONSE, response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}

	videos := form.File["videos"]
	projectIdSlice := form.Value["id"]
	if len(projectIdSlice) > 1 {
		c.Set(define.ANNOTATIONRESPONSE, response.JSONError(response.ERROR_UPLOAD_NOT_ID))
		c.Abort()
		return
	}
	var projectId int

	if len(projectIdSlice) > 0 {
		projectId = numberu.ToInt(projectIdSlice[0])
	} else {
		projectId = 0
	}

	if len(videos) > 1 || len(videos) == 0 {
		c.Set(define.ANNOTATIONRESPONSE, response.JSONErrorWithMsg("只能上传一个视频文件"))
		c.Abort()
		return
	}
	count, err := upload_service.SaveUploadedVideo(projectId, userId, videos[0])
	if err != nil {
		c.Set(define.ANNOTATIONRESPONSE, response.JSONErrorWithMsg(err.Error()))
		c.Abort()
		return
	}
	c.Set(define.ANNOTATIONRESPONSE, response.JSONData(count))
	return
}
