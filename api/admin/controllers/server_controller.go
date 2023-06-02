package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-slim/api/admin/middlewares"
	"go-slim/api/admin/schema"
	services2 "go-slim/api/admin/services"
	"go-slim/internal/services"
	"gorm.io/gorm"
	"mime/multipart"
)

type ServerController struct {
	UrlGen   *services.UrlGen
	FormFile *services2.FormFile
	DB       *gorm.DB
}
type singleFileReq struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
	Path string                `form:"path" binding:"required"`
}

func (s *ServerController) UploadImage(c *gin.Context) {
	req := new(singleFileReq)
	if err := c.ShouldBindWith(req, binding.FormMultipart); err != nil {
		fail(c, err)
		return
	}
	up := s.FormFile.NewModel(req.File, req.Path)
	localPath, err := s.FormFile.LocalUploadDir(up.Path)
	if err != nil {
		fail(c, err)
		return
	}
	if err = c.SaveUploadedFile(req.File, localPath); err != nil {
		fail(c, err)
		return
	}

	up.IsComplete = 1
	up.UserID = uint(middlewares.GetTokenUserID(c))
	s.DB.Create(up)

	upUrl := s.UrlGen.Uploads(up.Path)

	success(c, schema.ProUploadRes{
		Url:  upUrl,
		Path: up.Path,
		ID:   up.ID,
	})
}
