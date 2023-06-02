package models

import (
	"crypto/md5"
	"encoding/hex"
	"gorm.io/gorm"
	"time"
)

// User
type User struct {
	ID        uint           `gorm:"column:id;primarykey"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Username string `gorm:"column:username;type:varchar(32);default:null;uniqueIndex:uni_user_username_app,priority:1" json:"username"` // 账号登录
	Password string `gorm:"column:password;type:varchar(32);default:null"`
	Mobile   string `gorm:"column:mobile;type:varchar(11);default:null;uniqueIndex:uni_user_mobile_app,priority:1"`
	Email    string `gorm:"column:email;type:varchar(32);default:null;uniqueIndex:uni_user_email_app,priority:1"`
	UID      uint   `gorm:"column:uid;type:bigint(20) unsigned;default:null;uniqueIndex:uni_uid"`
	AppID    uint   `gorm:"column:app_id;type:bigint(20) unsigned;index:fk_user_app;uniqueIndex:uni_user_username_app,priority:2;uniqueIndex:uni_user_mobile_app,priority:2;uniqueIndex:uni_user_email_app,priority:2;" json:"app_id"` // 应用

	Clients []*UserClient `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Info    *UserInfo     `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	App     *App          `gorm:"foreignKey:AppID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	// https://gorm.io/zh_CN/docs/settings.html
	// u.UID = generateUID()
	if u.Password != "" {
		ctx := md5.New()
		ctx.Write([]byte(u.Password))
		u.Password = hex.EncodeToString(ctx.Sum(nil))
	}
	return nil
}

// UserInfo
type UserInfo struct {
	ID        uint      `gorm:"column:id;primarykey"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	Avatar   string    `gorm:"column:avatar;type:longtext" json:"avatar"`
	Nickname string    `gorm:"column:nickname;type:longtext" json:"nickname"`
	Gender   uint      `gorm:"column:gender;type:tinyint(3) unsigned" json:"gender"`
	Birthday time.Time `gorm:"column:birthday;type:date" json:"birthday"`
	UserID   uint      `gorm:"column:user_id;type:bigint(20) unsigned;index:fk_user_info" json:"user_id"` // 会员

	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// UserClient
type UserClient struct {
	ID        uint      `gorm:"column:id;primarykey"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	UserAgent      string `gorm:"column:user_agent;type:varchar(128)" json:"user_agent"`
	AcceptLanguage string `gorm:"column:accept_language;type:varchar(64)" json:"accept_language"`
	ClientIP       string `gorm:"column:client_ip;type:varchar(16)" json:"client_ip"`
	DeviceUUID     string `gorm:"column:device_uuid;type:varchar(48);uniqueIndex:uni_user_client_device_uuid,priority:1" json:"device_uuid"` // 服务UUID登录
	DeviceID       string `gorm:"column:device_id;type:varchar(64);uniqueIndex:uni_user_client_device_id,priority:1" json:"device_id"`       // 设备ID登录
	AppVersion     string `gorm:"column:app_version;type:varchar(16);" json:"app_version"`                                                   // 应用版本

	AppID  uint  `gorm:"column:app_id;type:bigint(20) unsigned;index:fk_user_client_app;uniqueIndex:uni_user_client_device_uuid,priority:2;uniqueIndex:uni_user_client_device_id,priority:2" json:"app_id"` // 应用
	UserID uint  `gorm:"column:user_id;type:bigint(20) unsigned;index:fk_user_clients" json:"user_id"`                                                                                                      // 会员
	App    *App  `gorm:"foreignKey:AppID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	User   *User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type UserUpload struct {
	ID        uint      `gorm:"column:id;primarykey"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	UserID     uint           `gorm:"column:user_id;type:bigint(20) unsigned" json:"user_id"`
	Name       string         `gorm:"column:name;type:varchar(128)" json:"name"`
	Type       string         `gorm:"column:type;type:varchar(32)" json:"type"`
	Size       uint           `gorm:"column:size;type:int(10) unsigned" json:"size"`
	Path       string         `gorm:"column:path;type:varchar(128)" json:"path"`
	Hash       string         `gorm:"column:hash;type:varchar(64)" json:"hash"`
	Disk       string         `gorm:"column:disk;type:varchar(32)" json:"disk"`
	Rule       string         `gorm:"column:rule;type:varchar(100)" json:"rule"`
	OwnerID    uint           `gorm:"column:owner_id;type:bigint(20) unsigned" json:"owner_id"`
	OwnerType  string         `gorm:"column:owner_type;type:varchar(100)" json:"owner_type"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;index:user_upload_deleted_at_IDX,priority:1" json:"deleted_at"`
	IsComplete uint           `gorm:"column:is_complete;type:tinyint(1)" json:"is_complete"`

	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
