package project

import (
	"annotation/model/tag"
	"annotation/model/upload"
	"annotation/model/user"
	"time"
)

type ProjectType int

const (
	Pcreated ProjectType = iota + 1
	Pworking
	PpendingReview
	Paccept
)

type Project struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"size:20"`
	Description string    `json:"description" gorm:"size:50"`
	CreatedTime time.Time `json:"createdTime"`

	//使用的标记类的Id
	Class   tag.Class `json:"class"  gorm:"foreignKey:ClassId;references:Id"`
	ClassId int       `json:"classId"`
	//认领该任务的人的Id
	Workers []user.User `json:"workers" gorm:"many2many:project_workers;"`
	//该任务包含的图片
	Images []upload.Image `json:"images" gorm:"many2many:project_images;"`
	//该任务产生的Annotation
	Annotations []Annotation `json:"annotations" gorm:"foreignKey:ProjectId;references:Id"`

	Creator   user.User `json:"creator" gorm:"foreignKey:CreatorId;references:UserId"`
	CreatorId int

	Type ProjectType `json:"type"`

	AnnotationMap map[int]Annotation `json:"annotationMap" gorm:"-"`
}

type ProjectCreateReq struct {
	Name        string `json:"name" binding:"required,max=20"`
	Description string `json:"description" binding:"required,max=50"`
	ClassId     int    `json:"classId" binding:"required"`
	UserList    []int  `json:"userList" binding:"required"`
}

type ProjectCreateResp struct {
	Id int `json:"id"`
}

type ProjectListResp struct {
	Id            int         `json:"id"`
	Name          string      `json:"name"`
	Description   string      `json:"description"`
	ClassName     string      `json:"className"`
	ImagesNum     int         `json:"imagesNum"`
	AnnotationNum int         `json:"annotationNum"`
	WorkerNum     int         `json:"workerNum"`
	Type          ProjectType `json:"type"`
	CreatedTime   time.Time   `json:"createdTime"`
}

type ProjectChangeStatusReq struct {
	Type ProjectType `json:"type" binding:"required"`
}

type ProjectAddPublicReq struct {
	ImageId   int `json:"imageId" binding:"required"`
	ProjectId int `json:"projectId" binding:"required"`
}
