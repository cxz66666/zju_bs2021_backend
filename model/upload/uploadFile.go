package upload

import (
	"annotation/model/user"
	"fmt"
	"time"
)

type StoreType int
const (
	Backend StoreType=iota+1
	OSS
)
//Image 存储文件的格式为 Md5Hash.jpg/png Crc32的作用是相当一级索引加速查询，只查md5非常的慢
type Image struct {
	Id int `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"size:50"`
	ProjectId int `json:"projectId" gorm:"index"`
	Crc32Hash  uint32 `json:"crc32Hash" gorm:"index"`
	//Md5Hash string `json:"md5Hash" gorm:"size:40"`

	// KB
	//Size int `json:"size"`
	Type StoreType `json:"type"`
	StorePath string `json:"storePath" gorm:"size:150"`
	Creator user.User `json:"creator" gorm:"foreignKey:CreatorId;references:UserId"`
	CreatorId int `json:"creatorId" `
	UploadTime time.Time `json:"uploadTime"`
}


type ImageResp struct {
	Id int `json:"id"`
	Name string `json:"name"`
	ProjectId int `json:"projectId"`
	Type StoreType `json:"type"`
	Url string `json:"url"`
	CreatorId int `json:"creatorId"`
	CreatorName string `json:"creatorName"`
	UploadTime string `json:"uploadTime"`
}



func (image *Image) GetUrl() string {
	if image.Type==Backend {
		return fmt.Sprintf("/api/image/%d/%d/%s",image.ProjectId,image.Crc32Hash,image.Name)
	} else if image.Type==OSS{
		return image.StorePath
	} else {
		return "https://baidu.com"
	}
}