package setting

import "github.com/aliyun/aliyun-oss-go-sdk/oss"

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
