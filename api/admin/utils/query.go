package utils

import (
	"gorm.io/gorm"
)

// PageQuery
// Deprecated
func PageQuery[M any](db *gorm.DB, result []interface{}, page int, limit int) (total int, err error) {
	offset := (page - 1) * limit
	if err = db.Offset(offset).Limit(limit).Order("id desc").Find(&result).Error; err != nil {
		return
	}

	var count int64
	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	if err = db.Offset(-1).Limit(-1).Count(&count).Error; err != nil {
		return
	}

	return
}
