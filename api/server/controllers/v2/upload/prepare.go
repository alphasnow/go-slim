package upload

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-slim/api/server/controllers/v2/common"
	"go-slim/api/server/schema"
	"go-slim/api/server/services/file_upload"
	"go-slim/pkg/xcrypt"
	"time"
)

type Check struct {
	AesCbc *xcrypt.AesCbc
}

type checkReq struct {
	*file_upload.UploadInfo `json:",inline,omitempty"`

	CheckRule string `json:"check_rule" binding:"required"` // 文件验证规则 humanSegment
}
type checkResp struct {
	UploadToken string `json:"upload_token"`
}

// Prepare
func (ctl *Check) Prepare(c *gin.Context) {
	req := new(checkReq)
	if err := c.ShouldBindBodyWith(req, binding.JSON); err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	rule, err := file_upload.GetUploadRule(req.CheckRule)
	if err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	checker := new(file_upload.UploadChecker)
	err = checker.Validation(rule, req.UploadInfo)
	if err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	token := new(file_upload.UploadToken)
	token.UploadInfo = req.UploadInfo
	token.UploadInfo.FileHead = ""
	token.ExpireTime = time.Now().Add(1 * time.Hour).Unix()
	tokenStr, _ := ctl.AesCbc.EncryptBase64(token)

	common.Success(c, checkResp{UploadToken: tokenStr})
}
