package project_service

import (
	"annotation/model/project"
	"annotation/utils/db"
	"gorm.io/gorm/clause"
)

func CreateProject(nowProject *project.Project) error {
	if err:=db.MysqlDB.Create(nowProject).Error;err!=nil{
		return err
	}
	return nil
}

func ListProject(pageSize, current int) ([]project.Project,int,error) {
	var ans []project.Project
	var total int64
	if err:=db.MysqlDB.Model(&project.Project{}).Count(&total).Offset((current-1)*pageSize).Limit(pageSize).Preload(clause.Associations).Order("id desc").Find(&ans).Error;err!=nil{
		return nil,0,err
	}
	return ans,int(total),nil
}

// QueryProjectById 根据id查询class
func QueryProjectById(id int) (project.Project,error)  {
	var ans project.Project
	if err:=db.MysqlDB.Preload(clause.Associations).First(&ans,id).Error;err!=nil{
		return project.Project{},nil
	}
	return ans,nil
}