package models

import (
	"crypto/md5"
	"encoding/hex"
	"gorm.io/gorm"
	"time"
)

// AdminUser
type AdminUser struct {
	ID        uint           `json:"id" form:"-" gorm:"column:id;primarykey"`
	CreatedAt time.Time      `json:"created_at" form:"-" gorm:"column:created_at;type:datetime(0);"`
	UpdatedAt time.Time      `json:"updated_at" form:"-" gorm:"column:updated_at;type:datetime(0);"`
	DeletedAt gorm.DeletedAt `json:"-" form:"-" gorm:"column:deleted_at;type:datetime(0);index"`

	Username string `json:"username" form:"username" gorm:"column:username;type:varchar(32);default:null;uniqueIndex:uni_user_username_app,priority:1"` // 账号登录
	Password string `json:"password" form:"password" gorm:"column:password;type:varchar(32);default:null"`
	Phone    string `json:"phone" form:"phone" gorm:"column:phone;type:varchar(11);default:null;uniqueIndex:uni_user_phone_app,priority:1"` // 手机号登录
	Email    string `json:"email" form:"email" gorm:"column:email;type:varchar(32);default:null;uniqueIndex:uni_user_email_app,priority:1"` // 邮箱登录
	Name     string `json:"name" form:"name" gorm:"column:name;type:varchar(32)"`                                                           // 真实姓名
	Avatar   string `json:"avatar" form:"avatar" gorm:"column:avatar;type:varchar(64)"`

	DepartmentID uint `json:"department_id" form:"department_id" gorm:"column:department_id;type:bigint(20)"`
}

func (u *AdminUser) BeforeSave(tx *gorm.DB) error {
	// https://gorm.io/zh_CN/docs/settings.html
	if (u.Password != "" && u.ID == 0) || tx.Statement.Changed("Password") {
		ctx := md5.New()
		ctx.Write([]byte(u.Password))
		u.Password = hex.EncodeToString(ctx.Sum(nil))
	}
	return nil
}

type AdminDepartment struct {
	ID        uint           `json:"id" form:"-" gorm:"column:id;primarykey"`
	CreatedAt time.Time      `json:"created_at" form:"-" gorm:"column:created_at;type:datetime(0);"`
	UpdatedAt time.Time      `json:"updated_at" form:"-" gorm:"column:updated_at;type:datetime(0);"`
	DeletedAt gorm.DeletedAt `json:"-" form:"-" gorm:"column:deleted_at;type:datetime(0);index"`

	Name   string `json:"name" form:"name" gorm:"column:name;type:varchar(32)"`
	Remark string `json:"remark" form:"remark" gorm:"column:remark;type:varchar(64)"`
	PID    uint   `json:"pid" form:"pid" gorm:"column:pid"`
}

type AdminUpload struct {
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

	User *AdminUser `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
