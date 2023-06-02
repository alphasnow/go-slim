package models

import (
	"gorm.io/gorm"
	"time"
)

type model struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&Article{},
		&ArticleCategory{},
		&ArticleContent{},
		&ArticleTag{},

		&Tag{},

		&App{},

		&User{},
		&UserClient{},
		&UserInfo{},
		&UserUpload{},

		&AdminUser{},
		&AdminDepartment{},
	)
}

const DefaultAppid = "*****************"
const DefaultAppkey = "*************************************"

const DefaultAdminUsername = "****"
const DefaultAdminPassword = "****"

func AutoSeeder(db *gorm.DB) error {
	var appCount int64
	db.Model(&App{}).Count(&appCount)
	if appCount == 0 {
		db.Create(&App{Name: "default", Appid: DefaultAppid, Appkey: DefaultAppkey})
	}

	var AdminUserCount int64
	db.Model(&AdminUser{}).Count(&AdminUserCount)
	if AdminUserCount == 0 {
		db.Create(&AdminUser{Username: DefaultAdminUsername, Password: DefaultAdminPassword, Name: "****", Email: "****@admin.dev", Phone: "136********"})
	}
	return nil
}
