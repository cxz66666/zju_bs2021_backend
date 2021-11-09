package project_service

import (
	"annotation/model/project"
	"annotation/utils/db"
)

func CreateProject(nowProject *project.Project) error {
	if err:=db.MysqlDB.Create(nowProject).Error;err!=nil{
		return err
	}
	return nil
}