package project_service

import (
	"annotation/model/project"
	"annotation/utils/db"
	"gorm.io/gorm"
)

//ChangeRegion 用于改变annotation的状态Region,使用的是数据库事务
func ChangeRegion(crReq project.AnnotationRegionReq) error {
	err := db.MysqlDB.Transaction(func(tx *gorm.DB) error {

		for _, m := range crReq.Data {
			if err := tx.Model(&project.Annotation{
				Id: m.Id,
			}).Update("regions", m.Regions).Error; err != nil {
				return err
			}
		}
		return nil

	})
	return err
}

//ChangeAnnotationType 用于改变annotation的type,使用的是数据库事务
func ChangeAnnotationType(ctReq project.AnnotationTypeReq) error {
	err := db.MysqlDB.Transaction(func(tx *gorm.DB) error {

		for _, m := range ctReq.Ids {
			if err := tx.Model(&project.Annotation{
				Id: m,
			}).Update("type", ctReq.Type).Error; err != nil {
				return err
			}
		}
		return nil
	})
	return err
}
