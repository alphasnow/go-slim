package models

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	ID        uint           `json:"id" form:"-" gorm:"column:id;primarykey"`
	CreatedAt time.Time      `json:"created_at" form:"-" gorm:"column:created_at;type:datetime(0);"`
	UpdatedAt time.Time      `json:"updated_at" form:"-" gorm:"column:updated_at;type:datetime(0);"`
	DeletedAt gorm.DeletedAt `json:"-" form:"-" gorm:"column:deleted_at;type:datetime(0);index"`

	CategoryID  uint   `gorm:"column:category_id;type:bigint(20) unsigned" json:"category_id"`
	Title       string `gorm:"column:title;type:varchar(64)" json:"title"`
	Keywords    string `gorm:"column:keywords;type:varchar(128)" json:"keywords"`
	Description string `gorm:"column:description;type:varchar(128)" json:"description"`
	Thumb       string `gorm:"column:thumb;type:varchar(128)" json:"thumb"`

	ArticleContent ArticleContent `gorm:"foreignkey:ArticleID;references:ID"`
	Tags           []*Tag         `gorm:"many2many:article_tag"`
}

type ArticleCategory struct {
	ID        uint           `json:"id" form:"-" gorm:"column:id;primarykey"`
	CreatedAt time.Time      `json:"created_at" form:"-" gorm:"column:created_at;type:datetime(0);"`
	UpdatedAt time.Time      `json:"updated_at" form:"-" gorm:"column:updated_at;type:datetime(0);"`
	DeletedAt gorm.DeletedAt `json:"-" form:"-" gorm:"column:deleted_at;type:datetime(0);index"`

	Title       string `gorm:"column:title;type:varchar(64)" json:"title"`
	Keywords    string `gorm:"column:keywords;type:varchar(128)" json:"keywords"`
	Description string `gorm:"column:description;type:varchar(128)" json:"description"`

	articles []Article `gorm:"foreignKey:CategoryID;references:ID"`
}

type ArticleContent struct {
	ArticleID uint   `gorm:"column:article_id;type:bigint(20) unsigned;index:fk_article_article_content,priority:1" json:"article_id"`
	Content   string `gorm:"column:content;type:longtext" json:"content"`
}

type ArticleTag struct {
	ArticleID uint `gorm:"column:article_id;type:bigint(20) unsigned" json:"article_id"`
	TagID     uint `gorm:"column:tag_id;type:bigint(20) unsigned" json:"tag_id"`
}
