package verify

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-slim/api/server/controllers/v2/common"
	"go-slim/api/server/schema"
	"go-slim/pkg/xcaptcha"
	"go-slim/pkg/xuuid"
)

type Captcha struct {
	StringCaptcha *xcaptcha.StringCaptcha
}
type captchaImageReq struct {
	VerifyPath string `json:"verify_path" binding:"required"`
	DeviceUUID string `json:"device_uuid"`
}
type captchaImageRes struct {
	captchaImageReq
	VerifyImage string `json:"verify_image"`
}

func (s *Captcha) Image(c *gin.Context) {
	req := captchaImageReq{}
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}

	if req.DeviceUUID == "" {
		req.DeviceUUID = xuuid.Generate()
	}

	key := req.DeviceUUID + ":" + req.VerifyPath
	_, img, err := s.StringCaptcha.Generate(c, key)
	if err != nil {
		common.Error(c, schema.BadRequest, err)
		return
	}
	imgBas64 := base64.StdEncoding.EncodeToString(img)

	common.Success(c, captchaImageRes{req, imgBas64})
}
