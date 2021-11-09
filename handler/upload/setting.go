package upload

import (
	"annotation/define"
	"annotation/model/upload"
	"annotation/utils/logging"
	"annotation/utils/response"
	"annotation/utils/setting"
	"github.com/gin-gonic/gin"
)


// GetSetting 返回当前的配置
func GetSetting(c *gin.Context)  {
	c.Set(define.ANNOTATIONRESPONSE,response.JSONData(setting.UploadSetting))
	return
}

func UpdateSetting(c *gin.Context){
	newSetting :=setting.Upload{}
	oldSetting:=setting.UploadSetting
	if err:=c.ShouldBind(&newSetting);err!=nil{
		c.Set(define.ANNOTATIONRESPONSE,response.JSONError(response.ERROR_PARAM_FAIL))
		c.Abort()
		return
	}

	if newSetting.Type!=0&& newSetting.Type!=setting.UploadSetting.Type{
		setting.UploadSetting.Type= newSetting.Type
	}

	if len(newSetting.BackendPath)>0 {
		setting.UploadSetting.BackendPath= newSetting.BackendPath
	}

	if len(newSetting.Region)>0 {
		setting.UploadSetting.Region= newSetting.Region
	}
	if len(newSetting.AccessKeyId)>0 {
		setting.UploadSetting.AccessKeyId= newSetting.AccessKeyId
	}
	if len(newSetting.AccessKeySecret)>0 {
		setting.UploadSetting.AccessKeySecret= newSetting.AccessKeySecret
	}
	if len(newSetting.Bucket)>0 {
		setting.UploadSetting.Bucket= newSetting.Bucket
	}
	if len(newSetting.OSSPath)>0 {
		setting.UploadSetting.OSSPath= newSetting.OSSPath
	}
	//如果使用oss 需要更新
	if setting.UploadSetting.Type==upload.OSS {
		if err:=setting.SetupBucket();err!=nil{
			logging.Info("change setting error")
			setting.UploadSetting=oldSetting
			c.Set(define.ANNOTATIONRESPONSE,response.JSONErrorWithMsg("设置的参数(账号密码等)错误，更新失败"))
			c.Abort()
			return
		}
	}
	setting.SaveUploadSetting()

	c.Set(define.ANNOTATIONRESPONSE,response.JSONData("success"))
	return

}