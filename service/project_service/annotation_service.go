package project_service

import (
	"annotation/model/project"
	"annotation/utils/db"
)

//ChangeRegion 用于改变annotation的状态Region
func ChangeRegion(id int,regions string) error {
	if err:=db.MysqlDB.Model(&project.Annotation{
		Id: id,
	}).Update("regions",regions).Error;err!=nil{
		return err
	}
	return nil
}