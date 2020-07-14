package model

import (
	"media-matrix/lib/helper"
	"media-matrix/lib/mysql"

	"github.com/jinzhu/gorm"
)

type Cate struct {
	Id        int              `json:"id"`
	Name      string           `gorm:"not null;default:'';type:varchar(100) " json:"name" `
	Detailed  string           `gorm:"not null;default:'';type:varchar(100) " json:"detailed" `
	Icon      string           `gorm:"not null;default:'';type:varchar(1000) " json:"icon" `
	Pid       int              `gorm:"not null;default:0; " json:"pid" `
	Sort      int8             `gorm:"not null;default:0; " json:"sort" `
	Recommend int8             `gorm:"not null;default:0; " json:"recommend" `
	CreatedAt helper.JSONTime  `gorm:"not null;type:datetime;" json:"created_at"`
	UpdatedAt helper.JSONTime  `gorm:"not null;type:datetime;" json:"updated_at"`
	DeletedAt *helper.JSONTime `gorm:"default:null;type:datetime;" json:"deleted_at"`
}

func (m Cate) Model() *gorm.DB {
	return mysql.Mysql((&m).Connection()).Model(&m)
}

func (*Cate) Connection() string {
	return "default"
}

func (*Cate) TableName() string {
	return "cate"
}
