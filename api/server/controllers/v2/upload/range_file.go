package upload

import (
	"github.com/gin-gonic/gin"
	"go-slim/api/server/controllers/v2/common"
	"go-slim/api/server/middlewares"
	"go-slim/api/server/schema"
	"go-slim/api/server/services/file_upload"
	"go-slim/internal/models"
	"go-slim/pkg/xcrypt"
	"gorm.io/gorm"
	"io"
	"mime/multipart"
	"os"
)

type RangeFile struct {
	FormFile *file_upload.FormFile
	FormData *file_upload.FormData
	Mysql    *gorm.DB
	AesCbc   *xcrypt.AesCbc
}
type rangeFileReq struct {
	File               *multipart.FileHeader `form:"file" binding:"required"`
	UploadID           uint                  `form:"upload_id"`
	ContentRange       string                `header:"Content-Range"`
	ContentRangeParsed [3]int                `form:"-" header:"-"`
}
type rangeFileRes struct {
	IsComplete uint `json:"is_complete"`
	UploadID   uint `json:"upload_id"`
}

func (ctl *RangeFile) Upload(c *gin.Context) {
	req := new(rangeFileReq)
	if err := c.ShouldBind(req); err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}
	if err := c.ShouldBindHeader(req); err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	var err error
	req.ContentRangeParsed, err = ctl.FormData.ParseContentRange(req.ContentRange)
	if err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	if req.UploadID == 0 {
		ctl.uploadFirst(c, req)
		return
	}

	ctl.uploadAppend(c, req)
}

func (ctl *RangeFile) uploadFirst(c *gin.Context, req *rangeFileReq) {

	up := ctl.FormFile.NewModel(req.File, "tmp")
	up.Size = uint(req.ContentRangeParsed[2])

	localPath, err := ctl.FormFile.LocalUploadDir(up.Path)
	if err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	if err = c.SaveUploadedFile(req.File, localPath); err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	ctl.Mysql.Create(up)

	common.Success(c, rangeFileRes{
		IsComplete: uint(up.IsComplete),
		UploadID:   up.ID,
	})
}
func (ctl *RangeFile) uploadAppend(c *gin.Context, req *rangeFileReq) {
	up := &models.UserUpload{}
	ctl.Mysql.Where("id=?", req.UploadID).First(up)

	localPath, err := ctl.FormFile.LocalUploadDir(up.Path)
	if err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	f, _ := os.OpenFile(localPath, os.O_RDWR|os.O_APPEND, 0666)
	defer f.Close()
	fo, err := req.File.Open()
	if err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}
	defer fo.Close()

	io.Copy(f, fo)

	if uint(req.ContentRangeParsed[1]) == up.Size-1 {
		up.IsComplete = 1
		ctl.Mysql.Where("id=?", up.ID).Updates(models.UserUpload{IsComplete: 1})

		uploadToken, err := middlewares.GetUploadToken(c)
		if err = uploadToken.ValidationByModal(up); err != nil {
			common.Error(c, schema.BadRequest, err)
			return
		}
	}

	common.Success(c, rangeFileRes{
		IsComplete: up.IsComplete,
		UploadID:   up.ID,
	})
}
