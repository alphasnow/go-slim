package utils

import (
	"go-slim/api/admin/schema"
	"gorm.io/gorm"
)

// ProTablePageList
// Deprecated
func ProTablePageList[M any](db *gorm.DB, result schema.ProTableList[M], page int, limit int) error {

	offset := (page - 1) * limit
	if err := db.Offset(offset).Limit(limit).Order("id desc").Find(&result).Error; err != nil {
		return err
	}

	var count int64
	if size := len(result.Data); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return nil
	}

	if err := db.Offset(-1).Limit(-1).Count(&count).Error; err != nil {
		return err
	}

	return nil
}
