package controllers

import (
	"github.com/gin-gonic/gin"
	"go-slim/api/admin/schema"
	"go-slim/api/admin/services"
	"go-slim/api/admin/utils"
	"go-slim/internal/models"
	"gorm.io/gorm"
)

type AdminDepartmentController struct {
	DB *gorm.DB
}

func (s *AdminDepartmentController) Index(c *gin.Context) {
	var req schema.ProTableReq
	if err := c.BindQuery(&req); err != nil {
		fail(c, err)
		return
	}

	tx := s.DB.Model(&models.AdminDepartment{}).WithContext(c)
	tx = s.search(c, tx)

	var data []models.AdminDepartment
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

func (s *AdminDepartmentController) search(c *gin.Context, tx *gorm.DB) *gorm.DB {
	search := c.QueryMap("search")
	if val, ok := search["name"]; ok == true {
		tx = tx.Where("name like ?", val+"%")
	}
	return tx
}

func (s *AdminDepartmentController) Create(c *gin.Context) {
	success(c, gin.H{})
}
func (s *AdminDepartmentController) Store(c *gin.Context) {
	// 获取全部参数
	data := models.AdminDepartment{}
	if err := c.Bind(&data); err != nil {
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

func (s *AdminDepartmentController) Edit(c *gin.Context) {
	// 获取id
	id := utils.QueryID(c)
	if id == 0 {
		fail(c, IDErr)
		return
	}

	data := models.AdminDepartment{}
	if res := s.DB.WithContext(c).Where("id=?", id).First(&data); res.Error != nil {
		fail(c, res.Error)
		return
	}
	success(c, gin.H{
		"data": data,
	})
}
func (s *AdminDepartmentController) Update(c *gin.Context) {
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
	data := models.AdminDepartment{}
	if err := c.Bind(&data); err != nil {
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

func (s *AdminDepartmentController) Show(c *gin.Context) {
	// 获取id
	id := utils.QueryID(c)
	if id == 0 {
		fail(c, IDErr)
		return
	}

	data := models.AdminDepartment{}
	if res := s.DB.WithContext(c).Where("id=?", id).First(&data); res.Error != nil {
		fail(c, res.Error)
		return
	}

	success(c, gin.H{
		"data": data,
	})
}

func (s *AdminDepartmentController) Destroy(c *gin.Context) {
	// 获取id
	id := utils.QueryID(c)
	if id == 0 {
		fail(c, IDErr)
		return
	}

	result := models.AdminDepartment{}
	if res := s.DB.WithContext(c).Where("id=?", id).Delete(&result); res.Error != nil {
		fail(c, res.Error)
		return
	}

	success(c, nil)
}

func (s *AdminDepartmentController) Form(c *gin.Context) {

	success(c, gin.H{})
}
