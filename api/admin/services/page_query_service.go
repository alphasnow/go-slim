package services

import "gorm.io/gorm"

// PageQueryService
// ser := NewPageQueryService{db,1,10,"id DESC"}
type pageQueryService struct {
	db   *gorm.DB
	sort string
	page int
	size int

	resultSize int
}

func NewPageQueryService(db *gorm.DB, page int, size int, sort string) *pageQueryService {
	return &pageQueryService{
		db: db, sort: sort, page: page, size: size,
	}
}

func (s *pageQueryService) Find(result interface{}) (err error) {
	offset := (s.page - 1) * s.size
	if err = s.db.Offset(offset).Limit(s.size).Order(s.sort).Find(result).Error; err != nil {
		return
	}
	return
}
func (s *pageQueryService) ResultSize(size int) *pageQueryService {
	s.resultSize = size
	return s
}
func (s *pageQueryService) Count(count *int64) (err error) {
	if 0 < s.size && 0 < s.resultSize && s.resultSize < s.size {
		offset := (s.page - 1) * s.size
		total := int64(s.resultSize + offset)
		count = &total
		return
	}

	if err = s.db.Offset(-1).Limit(-1).Count(count).Error; err != nil {
		return
	}
	return
}
