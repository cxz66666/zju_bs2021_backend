package upload_service

import (
	"annotation/model/project"
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


func QueryListImages(pid int,pageSize int,current int)([]upload.Image,int,error){
	var ans []upload.Image
	var total int64
	//如果是查询公共图片集
	if pid ==0 {
		if err:=db.MysqlDB.Model(&upload.Image{}).Where("project_id = ?",pid).Count(&total).Offset((current-1)*pageSize).Limit(pageSize).Order("id desc").Find(&ans).Error;err!=nil{
			return nil,0,err
		}
		return ans,int(total),nil
	}
	//count和详情分开查

	total=db.MysqlDB.Model(&project.Project{Id: pid}).Association("Images").Count()
	if err:=db.MysqlDB.Model(&project.Project{Id: pid}).Offset((current-1)*pageSize).Limit(pageSize).Order("id desc").Association("Images").Find(&ans);err!=nil{
		return nil,0,err
	}
	return ans,int(total),nil
}