package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-slim/api/server/controllers/v2/common"
	"go-slim/api/server/schema"
	"go-slim/api/server/services/file_upload"
	"go-slim/pkg/xcrypt"
	"time"
)

const UploadToken = "_internal/middlewares/upload_token"

func NewUploadTokenCheck(aesCbc *xcrypt.AesCbc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ut := c.PostForm("upload_token")
		if ut == "" {
			common.Error(c, schema.BadRequest, errors.New("upload token empty"))
			return
		}
		cReq := &file_upload.UploadToken{}
		if err := aesCbc.DecryptBase64(ut, &cReq); err != nil {
			common.Error(c, schema.BadRequest, err)
			return
		}
		if cReq.ExpireTime < time.Now().Unix() {
			common.Error(c, schema.BadRequest, errors.New("upload token expire"))
			return
		}

		c.Set(UploadToken, cReq)
		c.Next()
	}
}

func GetUploadToken(c *gin.Context) (*file_upload.UploadToken, error) {
	ut, exist := c.Get(UploadToken)
	if exist == false {
		return nil, errors.New("upload token not exist")
	}
	token, ok := ut.(*file_upload.UploadToken)
	if ok == false {
		return nil, errors.New("upload token type error")
	}
	return token, nil
}
