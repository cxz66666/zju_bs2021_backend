package project

import (
	"annotation/model/upload"
	"annotation/model/user"
	"time"
)

type AnnotateType int
const (
	Acreated AnnotateType =iota
	ApendingReview
	Aaccept

)

type Annotation struct {
	Id int `json:"id" gorm:"primaryKey"`
	Project Project `json:"projectId" gorm:"foreignKey:ProjectId;references:Id"`
	ProjectId  int
	Worker user.User `json:"worker" gorm:"foreignKey:WorkerId;references:UserId"`
	WorkerId int
	Image  upload.Image `json:"image" gorm:"foreignKey:ImageId;references:Id"`
	ImageId int

	Content string `json:"content"`

	Type AnnotateType
	LastEditTime time.Time `json:"lastEditTime"`
}