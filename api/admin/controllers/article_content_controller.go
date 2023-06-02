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

type ArticleContentController struct {
	DB *gorm.DB
}

func (s *ArticleContentController) Index(c *gin.Context) {
	var req schema.ProTableReq
	if err := c.BindQuery(&req); err != nil {
		fail(c, err)
		return
	}

	tx := s.DB.
		Model(&models.Article{}).
		WithContext(c)
	tx = s.search(c, tx)

	var data []models.Article
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

func (s *ArticleContentController) search(c *gin.Context, tx *gorm.DB) *gorm.DB {
	search := c.QueryMap("search")
	if val, ok := search["title"]; ok == true {
		tx = tx.Where("title like ?", val+"%")
	}
	return tx
}

func (s *ArticleContentController) Create(c *gin.Context) {
	success(c, gin.H{})
}
func (s *ArticleContentController) Store(c *gin.Context) {
	// 获取全部参数
	data := models.Article{}
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

func (s *ArticleContentController) Edit(c *gin.Context) {
	// 获取id
	id := utils.QueryID(c)
	if id == 0 {
		fail(c, IDErr)
		return
	}

	data := models.Article{}
	if res := s.DB.WithContext(c).Where("id=?", id).First(&data); res.Error != nil {
		fail(c, res.Error)
		return
	}

	success(c, gin.H{
		"data": data,
	})
}
func (s *ArticleContentController) Update(c *gin.Context) {
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
	data := models.Article{}
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

func (s *ArticleContentController) Show(c *gin.Context) {
	// 获取id
	id := utils.QueryID(c)
	if id == 0 {
		fail(c, IDErr)
		return
	}

	data := models.Article{}
	if res := s.DB.WithContext(c).Where("id=?", id).First(&data); res.Error != nil {
		fail(c, res.Error)
		return
	}

	success(c, gin.H{
		"data": data,
	})
}

func (s *ArticleContentController) Destroy(c *gin.Context) {
	// 获取id
	id := utils.QueryID(c)
	if id == 0 {
		fail(c, IDErr)
		return
	}

	result := models.Article{}
	if res := s.DB.WithContext(c).Where("id=?", id).Delete(&result); res.Error != nil {
		fail(c, res.Error)
		return
	}

	success(c, nil)
}

func (s *ArticleContentController) Form(c *gin.Context) {
	var dp []schema.ProFormOption[int]
	s.DB.Model(&models.ArticleCategory{}).Select("id as value", "name as label").Find(&dp)
	success(c, gin.H{
		"categories": dp,
	})
}
