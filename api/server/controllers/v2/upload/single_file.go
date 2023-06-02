package upload

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-slim/api/server/controllers/v2/common"
	"go-slim/api/server/middlewares"
	"go-slim/api/server/schema"
	"go-slim/api/server/services/file_upload"
	"gorm.io/gorm"
	"mime/multipart"
)

type SingleFile struct {
	Mysql    *gorm.DB
	FormFile *file_upload.FormFile
}
type SingleFileReq struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}
type SingleFileRes struct {
	IsComplete uint `json:"is_complete"`
	UploadID   uint `json:"upload_id"`
}

func (s *SingleFile) Upload(c *gin.Context) {
	req := new(SingleFileReq)
	if err := c.ShouldBindWith(req, binding.FormMultipart); err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}
	fh, err := c.FormFile("file")
	if err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	up := s.FormFile.NewModel(fh, "tmp")

	uploadToken, err := middlewares.GetUploadToken(c)
	if err = uploadToken.ValidationByModal(up); err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	localPath, err := s.FormFile.LocalUploadDir(up.Path)
	if err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	if err = c.SaveUploadedFile(fh, localPath); err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	up.IsComplete = 1
	s.Mysql.Create(up)

	common.Success(c, SingleFileRes{
		UploadID:   up.ID,
		IsComplete: 1,
	})
}
