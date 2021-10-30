package tag

import (
	"annotation/model/user"
	"time"
)

type Class struct {
	Id int  `gorm:"primaryKey"`
	ClassName string `json:"className" gorm:"size:30"`
	Description string `json:"description" gorm:"size:40"`
	Creator user.User `json:"creator" gorm:"foreignKey:CreatorId;references:UserId"`
	CreatorId int `json:"creatorId" `
	CreateTime time.Time

	Tags []Tag `gorm:"foreignKey:ClassId"`
}



type Tag struct {
	Id int `gorm:"primaryKey" json:"id"`
	ClassId int `json:"classId"`
	Content string `gorm:"size:20" json:"content"`
}



type ClassCreateReq struct {
	ClassName string `json:"className" binding:"required,max=30"`
	Description string `json:"description" binding:"max=40"`
	Tags []TagCreateReq
}

type TagCreateReq struct {
	Content string `json:"content" binding:"required,max=20"`
}

type ClassInfoResp struct {
	Id int `json:"id"`
	ClassName string `json:"className"`
	Description string `json:"description"`
	CreatorName string `json:"creatorName"`
	Tags []Tag
}