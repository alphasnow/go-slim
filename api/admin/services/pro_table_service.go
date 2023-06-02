package services

import (
	"go-slim/api/admin/schema"
	"gorm.io/gorm"
)

// ProTableService
// Deprecated: Use NewPageQueryService
type ProTableService struct {
	DB *gorm.DB
}

func (s *ProTableService) PageQuery(page int, limit int) (*schema.ProTableRes, error) {
	var result []map[string]interface{}
	offset := (page - 1) * limit
	if err := s.DB.Offset(offset).Limit(limit).Order("id desc").Find(&result).Error; err != nil {
		return nil, err
	}

	var count int64
	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return &schema.ProTableRes{Data: result, Total: count}, nil
	}

	if err := s.DB.Offset(-1).Limit(-1).Count(&count).Error; err != nil {
		return nil, err
	}

	return &schema.ProTableRes{Data: result, Total: count}, nil
}
