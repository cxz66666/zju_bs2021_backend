package tag

import (
	"annotation/model/user"
	"time"
)

type Class struct {
	Id int  `json:"id" gorm:"primaryKey"`
	ClassName string `json:"className" gorm:"size:30"`
	Description string `json:"description" gorm:"size:40"`
	Creator user.User `json:"creator" gorm:"foreignKey:CreatorId;references:UserId"`
	CreatorId int `json:"creatorId" `
	CreateTime time.Time `json:"createTime"`

	Tags []Tag `json:"tags" gorm:"foreignKey:ClassId"`
}



type Tag struct {
	Id int `gorm:"primaryKey" json:"id"`
	ClassId int `json:"classId"`
	Content string `gorm:"size:20" json:"content"`
}

type DeleteTagReq struct {
	TagId int `json:"tagId" binding:"required"`
}


type ClassCreateReq struct {
	ClassName string `json:"className" binding:"required,max=30"`
	Description string `json:"description" binding:"max=40"`
	Tags []string `json:"tags" binding:"required"`
}

type TagCreateReq struct {
	ClassId int `json:"classId" binding:"required"`
	Content string `json:"content" binding:"required,max=20"`
}

type ClassInfoResp struct {
	Id int `json:"id"`
	ClassName string `json:"className"`
	Description string `json:"description"`
	CreatorName string `json:"creatorName"`
	CreateTime time.Time `json:"createTime"`
	Tags []Tag `json:"tags"`
}

type ClassDeleteReq struct {
	Id int `json:"id" binding:"required"`
}

type ClassUpdateReq struct {
	Id int `json:"id" binding:"required"`
	ClassName string `json:"className" binding:"max=30"`
	Description string `json:"description" binding:"max=40"`
}

type ClassChooseResp struct {
	Label string  `json:"label"`
	Value int 	  `json:"value"`
}