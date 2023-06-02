package models

import (
	"gorm.io/gorm"
	"time"
)

type App struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"column:name;type:varchar(128);not null" json:"name"`
	Appid     string         `gorm:"column:appid;type:varchar(64);not null" json:"appid"`
	Appkey    string         `gorm:"column:appkey;type:varchar(64);not null" json:"appkey"`
}
