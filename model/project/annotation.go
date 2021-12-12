package project

import (
	"annotation/model/upload"
	"annotation/model/user"
	"time"
)

type AnnotateType int

const (
	Acreated AnnotateType = iota
	Afinish
	AReview
	Aaccept
)

type Annotation struct {
	Id        int          `json:"id" gorm:"primaryKey"`
	ProjectId int          `json:"projectId"`
	Worker    user.User    `json:"worker" gorm:"foreignKey:WorkerId;references:UserId"`
	WorkerId  int          `json:"workerId"`
	Image     upload.Image `json:"image" gorm:"foreignKey:ImageId;references:Id"`
	ImageId   int          `json:"imageId"`

	Regions   string `json:"regions"`
	PixelSize string `json:"pixelSize"`

	Type         AnnotateType `json:"type"`
	LastEditTime time.Time    `json:"lastEditTime"`

	Src string `json:"src" gorm:"-"`
}

//用于update
type AnnotationRegionReq struct {
	Data []AnnotationReginUpdate `json:"data"`
}

//用于单个标注内容的修改
type AnnotationReginUpdate struct {
	Id        int    `json:"id" binding:"required"`
	Regions   string `json:"regions"`
	PixelSize string `json:"pixelSize"`
}

type AnnotationTypeReq struct {
	Ids  []int        `json:"ids" binding:"required"`
	Type AnnotateType `json:"type" binding:"required"`
}
