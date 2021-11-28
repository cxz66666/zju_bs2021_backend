package project

import (
	"annotation/model/upload"
	"annotation/model/user"
	"time"
)

type AnnotateType int
const (
	Acreated AnnotateType =iota
	Afinish
	ApendingReview
	Aaccept

)

type Annotation struct {
	Id int `json:"id" gorm:"primaryKey"`
	Project Project `json:"project" gorm:"foreignKey:ProjectId;references:Id"`
	ProjectId  int `json:"projectId"`
	Worker user.User `json:"worker" gorm:"foreignKey:WorkerId;references:UserId"`
	WorkerId int `json:"workerId"`
	Image  upload.Image `json:"image" gorm:"foreignKey:ImageId;references:Id"`
	ImageId int `json:"imageId"`

	Regions string `json:"regions"`
	Type AnnotateType `json:"type"`
	LastEditTime time.Time `json:"lastEditTime"`

	Src string `json:"src" gorm:"-"`

}

//用于update
type AnnotationRegionReq struct {
	Id int 	`json:"id" binding:"required"`
	Regions string `json:"regions" binding:"required"`
}

type AnnotationTypeReq struct {
	Id int `json:"id" binding:"required"`
	Type AnnotateType `json:"type" binding:"required"`

}