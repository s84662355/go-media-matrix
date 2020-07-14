package model

import (
	"media-matrix/lib/helper"
	"media-matrix/lib/mysql"

	"github.com/jinzhu/gorm"
)

type Banner struct {
	Id        int              `json:"id"`
	ImgUrl    string           `gorm:"not null;default:'';type:varchar(100) " json:"img_url" `
	CateId    int              `gorm:"not null;default:0; " json:"cate_id" `
	Sort      int8             `gorm:"not null;default:0; " json:"sort" `
	CreatedAt helper.JSONTime  `gorm:"not null;type:datetime;" json:"created_at"`
	UpdatedAt helper.JSONTime  `gorm:"not null;type:datetime;" json:"updated_at"`
	DeletedAt *helper.JSONTime `gorm:"default:null;type:datetime;" json:"deleted_at"`
}

func (m Banner) Model() *gorm.DB {
	return mysql.Mysql((&m).Connection()).Model(&m)
}

func (*Banner) Connection() string {
	return "default"
}

func (*Banner) TableName() string {
	return "banner"
}
