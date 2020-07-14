package model

import (
	"media-matrix/lib/helper"
	"media-matrix/lib/mysql"

	"github.com/jinzhu/gorm"
)

type Article struct {
	Id        int              `json:"id"`
	Name      string           `gorm:"not null;default:'';type:varchar(100) " json:"name" `
	Content   string           `gorm:"default:null;type:text" json:"content" `
	CateId    int              `gorm:"not null;default:0; " json:"cate_id" `
	Sort      int8             `gorm:"not null;default:0; " json:"sort" `
	Recommend int8             `gorm:"not null;default:0; " json:"recommend" `
	ImgUrl    int8             `gorm:"not null;default:'';type:varchar(100) " json:"img_url" `
	CreatedAt helper.JSONTime  `gorm:"not null;type:datetime;" json:"created_at"`
	UpdatedAt helper.JSONTime  `gorm:"not null;type:datetime;" json:"updated_at"`
	DeletedAt *helper.JSONTime `gorm:"default:null;type:datetime;" json:"deleted_at"`
}

func (m Article) Model() *gorm.DB {
	return mysql.Mysql((&m).Connection()).Model(&m)
}

func (*Article) Connection() string {
	return "default"
}

func (*Article) TableName() string {
	return "article"
}
