package tag_service

import (
	"annotation/model/tag"
	"annotation/utils/db"
)


// QueryClass 根据pageSize和current 进行查询，由于各种原因 无法进行缓存
func QueryClass(pageSize,current int) ([]tag.Class,error)  {
	var ans []tag.Class
	if err:=db.MysqlDB.Offset((current-1)*pageSize).Limit(current).Preload("Tag").Find(&ans).Error;err!=nil{
		return nil,err
	}
	return ans,nil
}

// CreateClass 创建一个class，同时包含创建者ID和一并的tags
func CreateClass(class *tag.Class) error {
	if err:=db.MysqlDB.Create(class).Error;err!=nil{
		return err
	}
	return nil
}


