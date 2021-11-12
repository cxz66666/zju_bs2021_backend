package upload_service

import (
	"annotation/model/upload"
	"annotation/utils/db"
)

func QueryImage(pid int,crc32 int,fileName string) (upload.Image,error) {
	var ans upload.Image
	if err:=db.MysqlDB.Where("project_id = ? and crc32_hash = ? and name = ?",pid,crc32,fileName).First(&ans).Error;err!=nil{
		return ans,err
	}
	return ans,nil
}
