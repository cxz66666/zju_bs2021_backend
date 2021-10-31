package tag_service

import (
	"annotation/model/tag"
	"annotation/utils/db"
	"errors"
	"gorm.io/gorm/clause"
)


// QueryClass 根据pageSize和current 进行查询，由于各种原因 无法进行缓存
func QueryClass(pageSize,current int) ([]tag.Class,int,error)  {
	var ans []tag.Class
	var total int64
	if err:=db.MysqlDB.Model(&tag.Class{}).Count(&total).Offset((current-1)*pageSize).Limit(pageSize).Preload("Tags").Order("id desc").Find(&ans).Error;err!=nil{
		return nil,0,err
	}
	return ans,int(total),nil
}

// QueryClassById 根据id查询class
func QueryClassById(classId int) (tag.Class,error)  {
	var ans tag.Class
	if err:=db.MysqlDB.First(&ans,classId).Error;err!=nil{
		return tag.Class{},nil
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

// DeleteClass 删除改id的所有tag和内容 使用clause.Association
func DeleteClass(id int) error {
	tmp:=db.MysqlDB.Select(clause.Associations).Delete(&tag.Class{},id)
	if err:=tmp.Error;err!=nil{
		return err
	}
	if tmp.RowsAffected==0{
		return errors.New("不存在该id")
	}
	return nil
}


// UpdateClass 更新class的两个字段，只有"class_name","description"是允许被更新的
func UpdateClass(newClass tag.Class) error {
	if err:=db.MysqlDB.Model(&newClass).Select("class_name","description").Updates(newClass).Error;err!=nil{
		return err
	}
	return nil
}

// CreateTag 根据tag模型创建tag
func CreateTag(tag *tag.Tag) error  {
	if err:=db.MysqlDB.Create(tag).Error;err!=nil{
		return err
	}
	return nil
}

// DeleteTag 根据id删除tag
func DeleteTag(tagId int) error {
	tmp:=db.MysqlDB.Delete(&tag.Tag{},tagId)
	if err:=tmp.Error;err!=nil{
		return err
	}
	if tmp.RowsAffected==0{
		return errors.New("不存在该id")
	}
	return nil
}