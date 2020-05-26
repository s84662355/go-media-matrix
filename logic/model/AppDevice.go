package model

import (
    "github.com/jinzhu/gorm"
    "media-matrix/lib/mysql"
    "time"
)

 
type AppDevice struct {
    Id    uint64           `json:"id"`
    AndroidId string       `gorm:"not null;default:'';type:varchar(32) " json:"android_id" `
    DeviceId  string       `gorm:"not null;default:'';type:varchar(100) " json:"device_id" `
    Imei  string           `gorm:"not null;default:'';type:varchar(32) " json:"imei" ` 
    IsClone bool           `gorm:"not null;default:0;type:tinyint(4) " json:"is_clone" ` 
    MacAddress string      `gorm:"not null;default:'';type:varchar(32) " json:"mac_address" ` 
    TModel  string          `gorm:"column:model;not null;default:'';type:varchar(32) " json:"model" `   
    NetInfo string         `gorm:"not null;default:'';type:varchar(100) " json:"net_info" ` 
    NetType string         `gorm:"not null;default:'';type:varchar(100) " json:"net_type" ` 
    PluginVersion string   `gorm:"not null;default:'';type:varchar(20) " json:"plugin_version" ` 
    StaffId int64          `gorm:"not null;default:0;type:int(11);unsigned" json:"staff_id" ` 
    RegionId int64         `gorm:"not null;default:0;type:int(11)" json:"region_id" `  
    RegionName string      `gorm:"not null;default:'';type:varchar(100)" json:"region_name" ` 
    ScreenHeight int16     `gorm:"not null;default:0;type:int(11);unsigned" json:"screen_height" `  
    ScreenWidth int16      `gorm:"not null;default:0;type:int(11);unsigned" json:"screen_width" `  
    SerialId string        `gorm:"not null;default:'';type:varchar(50) " json:"serial_id" `  
    SystemVersion uint8    `gorm:"not null;default:0;type:tinyint(3) unsigned" json:"system_version" `   
    IsBusy uint8           `gorm:"not null;default:0;type:tinyint(3) " json:"is_busy" `
    OnLine uint8           `gorm:"column:online;not null;default:0;type:tinyint(3) " json:"online" `  
    CreatedAt time.Time    `gorm:"not null;type:datetime;" json:"created_at"`
    UpdatedAt time.Time    `gorm:"not null;type:datetime;" json:"updated_at"`
    DeletedAt *time.Time   `gorm:"default:null;type:datetime;" json:"deleted_at"`
}

func (m AppDevice) Model() *gorm.DB {
    return mysql.Mysql((&m).Connection()).Model(&m)
}

func (*AppDevice) Connection() string {
    return "default"
}

func (*AppDevice) TableName() string {
    return "media_app_device"
}
