package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-slim/api/admin/schema"
	"go-slim/api/admin/services"
	"go-slim/api/admin/utils"
	"go-slim/internal/models"
	"gorm.io/gorm"
)

type ArticleCategoryController struct {
	DB *gorm.DB
}

func (s *ArticleCategoryController) Index(c *gin.Context) {
	var req schema.ProTableReq
	if err := c.BindQuery(&req); err != nil {
		fail(c, err)
		return
	}

	tx := s.DB.
		Model(&models.ArticleCategory{}).
		WithContext(c)
	tx = s.search(c, tx)

	var data []models.ArticleCategory
	ser := services.NewPageQueryService(tx, req.GetPage(), req.GetSize(), "id DESC")
	if err := ser.Find(&data); err != nil {
		fail(c, err)
		return
	}
	var total int64
	if err := ser.ResultSize(len(data)).Count(&total); err != nil {
		fail(c, err)
		return
	}

	success(c, schema.ProTableRes{
		Data:  data,
		Total: total,
	})
}

func (s *ArticleCategoryController) search(c *gin.Context, tx *gorm.DB) *gorm.DB {
	search := c.QueryMap("search")
	if val, ok := search["title"]; ok == true {
		tx = tx.Where("title like ?", val+"%")
	}
	return tx
}

func (s *ArticleCategoryController) Create(c *gin.Context) {
	success(c, gin.H{})
}
func (s *ArticleCategoryController) Store(c *gin.Context) {
	// 获取全部参数
	data := models.ArticleCategory{}
	if err := c.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		fail(c, err)
		return
	}

	if res := s.DB.WithContext(c).Create(&data); res.Error != nil {
		fail(c, res.Error)
		return
	}

	success(c, gin.H{
		"data": data,
	})
}

func (s *ArticleCategoryController) Edit(c *gin.Context) {
	// 获取id
	id := utils.QueryID(c)
	if id == 0 {
		fail(c, IDErr)
		return
	}

	data := models.ArticleCategory{}
	if res := s.DB.WithContext(c).Where("id=?", id).First(&data); res.Error != nil {
		fail(c, res.Error)
		return
	}

	success(c, gin.H{
		"data": data,
	})
}
func (s *ArticleCategoryController) Update(c *gin.Context) {
	// 获取id
	id := utils.QueryID(c)
	if id == 0 {
		fail(c, IDErr)
		return
	}

	// 获取全部参数
	if err := c.Request.ParseForm(); err != nil {
		fail(c, err)
		return
	}
	data := models.ArticleCategory{}
	if err := c.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		fail(c, err)
		return
	}

	if res := s.DB.WithContext(c).Where("id=?", id).Updates(&data); res.Error != nil {
		fail(c, res.Error)
		return
	}

	success(c, gin.H{
		"data": data,
	})
}

func (s *ArticleCategoryController) Show(c *gin.Context) {
	// 获取id
	id := utils.QueryID(c)
	if id == 0 {
		fail(c, IDErr)
		return
	}

	data := models.ArticleCategory{}
	if res := s.DB.WithContext(c).Where("id=?", id).First(&data); res.Error != nil {
		fail(c, res.Error)
		return
	}

	success(c, gin.H{
		"data": data,
	})
}

func (s *ArticleCategoryController) Destroy(c *gin.Context) {
	// 获取id
	id := utils.QueryID(c)
	if id == 0 {
		fail(c, IDErr)
		return
	}

	result := models.ArticleCategory{}
	if res := s.DB.WithContext(c).Where("id=?", id).Delete(&result); res.Error != nil {
		fail(c, res.Error)
		return
	}

	success(c, nil)
}

func (s *ArticleCategoryController) Form(c *gin.Context) {

	success(c, gin.H{})
}
