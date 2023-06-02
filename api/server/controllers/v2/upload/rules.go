package upload

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-slim/api/server/controllers/v2/common"
	"go-slim/api/server/schema"
	"go-slim/api/server/services/file_upload"
)

type Rules struct {
}
type rulesReq struct {
	CheckRule string `json:"check_rule" binding:"required"` // 文件验证规则 HumanSegment
}
type rulesRes struct {
	UploadRule *file_upload.UploadRule `json:"upload_rule"`
}

func (ctl *Rules) GetRule(c *gin.Context) {
	req := new(rulesReq)
	if err := c.ShouldBindBodyWith(req, binding.JSON); err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	rule, err := file_upload.GetUploadRule(req.CheckRule)
	if err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	common.Success(c, rulesRes{rule})
}
