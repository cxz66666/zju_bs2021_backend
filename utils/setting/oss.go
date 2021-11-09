package setting

import (
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var Client *oss.Client
var Bucket *oss.Bucket

func  SetupBucket()  error {
	var err error
	Client, err = oss.New(UploadSetting.Region,UploadSetting.AccessKeyId, UploadSetting.AccessKeySecret)
	if err!=nil{
		Client=nil
		return err
	}
	Bucket,err=Client.Bucket(UploadSetting.Bucket)
	if err!=nil{
		Bucket=nil
		return err
	}
	return nil
}

func UploadImage(path string,content []byte) (error,string) {
	err:= Bucket.PutObject(path, bytes.NewReader(content) )
	if err != nil {
		fmt.Println("文件上传失败",path)
		return err,""
	}
	url:="https://"+UploadSetting.Bucket+"."+UploadSetting.Region+"/"+path
	return nil,url
}