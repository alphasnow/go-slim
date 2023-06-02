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

type AdminUserController struct {
	DB *gorm.DB
}

func (s *AdminUserController) Index(c *gin.Context) {
	var req schema.ProTableReq
	if err := c.BindQuery(&req); err != nil {
		fail(c, err)
		return
	}

	tx := s.DB.
		Model(&models.AdminUser{}).
		Select("id", "username", "name", "avatar", "phone", "email", "updated_at").
		WithContext(c)
	tx = s.search(c, tx)

	var data []models.AdminUser
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

func (s *AdminUserController) search(c *gin.Context, tx *gorm.DB) *gorm.DB {
	search := c.QueryMap("search")
	if val, ok := search["name"]; ok == true {
		tx = tx.Where("name like ?", val+"%")
	}
	if val, ok := search["username"]; ok == true {
		tx = tx.Where("username like ?", val+"%")
	}
	if val, ok := search["email"]; ok == true {
		tx = tx.Where("email like ?", val+"%")
	}
	if val, ok := search["phone"]; ok == true {
		tx = tx.Where("phone like ?", val+"%")
	}
	return tx
}

func (s *AdminUserController) Create(c *gin.Context) {
	success(c, gin.H{})
}
func (s *AdminUserController) Store(c *gin.Context) {
	// 获取全部参数
	data := models.AdminUser{}
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

func (s *AdminUserController) Edit(c *gin.Context) {
	// 获取id
	id := utils.QueryID(c)
	if id == 0 {
		fail(c, IDErr)
		return
	}

	data := models.AdminUser{}
	if res := s.DB.WithContext(c).Where("id=?", id).First(&data); res.Error != nil {
		fail(c, res.Error)
		return
	}

	data.Password = ""
	success(c, gin.H{
		"data": data,
	})
}
func (s *AdminUserController) Update(c *gin.Context) {
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
	data := models.AdminUser{}
	if err := c.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		fail(c, err)
		return
	}

	if res := s.DB.WithContext(c).Where("id=?", id).Updates(&data); res.Error != nil {
		fail(c, res.Error)
		return
	}

	data.Password = ""
	success(c, gin.H{
		"data": data,
	})
}

func (s *AdminUserController) Show(c *gin.Context) {
	// 获取id
	id := utils.QueryID(c)
	if id == 0 {
		fail(c, IDErr)
		return
	}

	data := models.AdminUser{}
	if res := s.DB.WithContext(c).Where("id=?", id).First(&data); res.Error != nil {
		fail(c, res.Error)
		return
	}

	data.Password = ""
	success(c, gin.H{
		"data": data,
	})
}

func (s *AdminUserController) Destroy(c *gin.Context) {
	// 获取id
	id := utils.QueryID(c)
	if id == 0 {
		fail(c, IDErr)
		return
	}

	result := models.AdminUser{}
	if res := s.DB.WithContext(c).Where("id=?", id).Delete(&result); res.Error != nil {
		fail(c, res.Error)
		return
	}

	success(c, nil)
}

func (s *AdminUserController) Form(c *gin.Context) {
	var adminRoles = []schema.ProFormOption[string]{
		{Label: "管理", Value: "admin"},
		{Label: "编辑", Value: "editor"},
	}

	var dp []schema.ProFormOption[int]
	s.DB.Model(&models.AdminDepartment{}).Select("id as value", "name as label").Find(&dp)
	success(c, gin.H{
		"roles":       adminRoles,
		"departments": dp,
	})
}
