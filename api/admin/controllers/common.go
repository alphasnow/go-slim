package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-slim/api/admin/schema"
	"go-slim/api/admin/services"
	"gorm.io/gorm"
	"net/http"
)

var (
	IDErr    = errors.New("id error")
	LoginErr = errors.New("login error")
)

func success(c *gin.Context, data any) {
	if data == nil {
		c.JSON(http.StatusOK, schema.Response{
			Success: true,
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

func fail(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusOK, schema.Response{
		Success:      false,
		ErrorMessage: err.Error(),
		ShowType:     schema.ERROR_MESSAGE,
	})
}

// modelPageQuery
// Deprecated
func modelPageQuery[T any](db *gorm.DB, req *schema.ProTableReq, res *schema.ProTableRes) (err error) {
	var data []T
	ser := services.NewPageQueryService(db, req.GetPage(), req.GetSize(), "id DESC")
	if err = ser.Find(&data); err != nil {
		return
	}
	res.Data = data
	if err = ser.ResultSize(len(data)).Count(&res.Total); err != nil {
		return
	}
	return
}
