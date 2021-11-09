package project

import (
	"annotation/model/tag"
	"annotation/model/upload"
	"annotation/model/user"
	"time"
)

type ProjectType int

const (
	Pcreated ProjectType=iota
	Pworking
	PpendingReview
	Paccept
)
type Project struct {
	Id int 	`json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"size:20"`
	Description string `json:"description" gorm:"size:50"`
	CreatedTime time.Time `json:"createdTime"`

	//使用的标记类的Id
	Class tag.Class `json:"class"  gorm:"foreignKey:ClassId;references:Id"`
	ClassId int `json:"classId"`
	//认领该任务的人的Id
	Workers []user.User `json:"workers" gorm:"many2many:project_workers;"`
	//该任务包含的图片
	Images []upload.Image `json:"images" gorm:"many2many:project_images;"`
	//该任务产生的Annotation
	Annotations []Annotation `json:"annotations" gorm:"many2many:project_annotation"`

	Creator user.User `json:"creator" gorm:"foreignKey:CreatorId;references:UserId"`
	CreatorId int

	Type ProjectType `json:"type"`
}