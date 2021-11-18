package upload_service

import (
	"annotation/model/upload"
	"annotation/utils/db"
	"gorm.io/gorm/clause"
)

func QueryImage(pid int,crc32 int,fileName string) (upload.Image,error) {
	var ans upload.Image
	if err:=db.MysqlDB.Where("project_id = ? and crc32_hash = ? and name = ?",pid,crc32,fileName).First(&ans).Error;err!=nil{
		return ans,err
	}
	return ans,nil
}


func QueryListImages(pageSize int,current int)([]upload.Image,int,error){
	var ans []upload.Image
	var total int64
	if err:=db.MysqlDB.Model(&upload.Image{}).Count(&total).Offset((current-1)*pageSize).Limit(pageSize).
		Preload(clause.Associations).Order("id desc").Find(&ans).Error;err!=nil{
			return nil,0,err
	}
	return ans,int(total),nil
}