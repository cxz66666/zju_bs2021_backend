package project_service

import (
	"annotation/model/project"
	"annotation/model/upload"
	"annotation/utils/db"
	"gorm.io/gorm/clause"
)

func CreateProject(nowProject *project.Project) error {
	if err := db.MysqlDB.Create(nowProject).Error; err != nil {
		return err
	}
	return nil
}

func ListProject(pageSize, current int) ([]project.Project, int, error) {
	var ans []project.Project
	var total int64
	if err := db.MysqlDB.Model(&project.Project{}).Count(&total).Offset((current - 1) * pageSize).Limit(pageSize).Preload(clause.Associations).Order("id desc").Find(&ans).Error; err != nil {
		return nil, 0, err
	}
	return ans, int(total), nil
}

//QueryProjectByWorker 根据workerId查他参与了多少项目
func QueryProjectByWorker(workerId int) (int, error) {
	var count int
	if err := db.MysqlDB.Raw("select COUNT(*) from project_workers where user_user_id = ?", workerId).Scan(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

// QueryProjectById 根据id查询class
func QueryProjectById(id int) (project.Project, error) {
	var ans project.Project
	if err := db.MysqlDB.Preload("Class.Tags").Preload("Annotations.Image").Preload(clause.Associations).First(&ans, id).Error; err != nil {
		return project.Project{}, err
	}
	return ans, nil
}

//ChangeStatus 用于改变project的状态Type
func ChangeStatus(id int, newType project.ProjectType) error {

	if err := db.MysqlDB.Model(&project.Project{
		Id: id,
	}).Update("type", newType).Error; err != nil {
		return err
	}
	return nil
}

//AddImageAssociation 用于添加公共图片到本项目
func AddImageAssociation(pid int, imageId int) error {
	if err := db.MysqlDB.Model(&project.Project{Id: pid}).Association("Images").Append(&upload.Image{Id: imageId}); err != nil {
		return err
	}
	return nil
}

// AddNewAnnotation 创建新的标注，此时标注仍然是空，type为Acreated
func AddNewAnnotation(pid int, newAnnotations []project.Annotation) error {
	if err := db.MysqlDB.Model(&project.Project{Id: pid}).Association("Annotations").Append(newAnnotations); err != nil {
		return err
	}
	return nil
}
