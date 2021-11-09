package upload_service

import (
	"annotation/model/upload"
	"annotation/utils/conv"
	"annotation/utils/crypto"
	"annotation/utils/db"
	file2 "annotation/utils/file"
	"annotation/utils/setting"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"path"
	"path/filepath"
	"strings"
	"time"
)

var ImageTypes =[]string{"image/png","image/jpeg"}
var VideoTypes = []string{"video/mp4"}
func Contains(slice []string, s string) int {
	for index, value := range slice {
		if value == s {
			return index
		}
	}
	return -1
}
func SaveUploadedImage(id int, userId int, file *multipart.FileHeader) error {
	//fmt.Println(file.Header)
	var content=file.Header.Get("Content-Type")
	if Contains(ImageTypes,content)==-1 {
		return errors.New("不是合法的文件类型，目前只支持png,jpg")
	}
	src, err := file.Open()
	if err != nil {
		return errors.New("文件打开错误")
	}
	defer src.Close()

	data, err := ioutil.ReadAll(src)
	if err!=nil{
		return err
	}
	crc32:=crypto.EncodeCrc32(string(data))
	//filepath是文件夹地址，dst是整个路径
	filePath:=setting.UploadSetting.BackendPath+ string(filepath.Separator)+ conv.Int2Str(id)+string(filepath.Separator)
	fileSuffix := path.Ext(file.Filename) //获取文件后缀
	filenameOnly:= strings.TrimSuffix(file.Filename, fileSuffix)//获取文件名
	dst:=filePath+filenameOnly+"-"+conv.Int2Str(int(crc32))+fileSuffix

	err=file2.IsNotExistMkDir(filePath)
	if err!=nil{
		return fmt.Errorf("file.IsNotExistMkDir src:%s, err: %v",src,err)
	}
	//如果存在则不做任何操作
	if file2.CheckExist(dst){
		return nil
	}

	err=ioutil.WriteFile(dst,data,0666);
	if err!=nil{
		return err
	}

	image:=upload.Image{
		Name: file.Filename,
		Type: setting.UploadSetting.Type ,
		ProjectId: id,
		StorePath: dst,
		Crc32Hash: crc32,
		CreatorId: userId,
		UploadTime: time.Now(),
	}

	return SaveImage(&image)
}

func SaveImage(image *upload.Image) error {
	if err:=db.MysqlDB.Create(image).Error;err!=nil{
		return err
	}
	return nil
}

// SaveUploadedVideo 用于转化视频文件为指定的图片，同时保存到相应的目录和db中
func SaveUploadedVideo(id int, userId int, file *multipart.FileHeader)  error {
	var content=file.Header.Get("Content-Type")
	if Contains(VideoTypes,content)==-1 {
		return errors.New("不是合法的类型，目前只支持mp4类型文件")
	}
	return nil
}